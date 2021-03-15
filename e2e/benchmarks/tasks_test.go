// +build e2e

// $ go test ./e2e/benchmarks -bench=. -tags e2e
package benchmarks

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/hashicorp/consul-terraform-sync/config"
	"github.com/hashicorp/consul-terraform-sync/controller"
	"github.com/hashicorp/consul-terraform-sync/testutils"
)

func BenchmarkCtrl_t01_s01(b *testing.B) {
	benchmarkCtrl(b, 1, 1)
}

func BenchmarkCtrl_t01_s50(b *testing.B) {
	benchmarkCtrl(b, 1, 50)
}

func BenchmarkCtrl_t10_s01(b *testing.B) {
	benchmarkCtrl(b, 10, 1)
}

func BenchmarkCtrl_t10_s50(b *testing.B) {
	benchmarkCtrl(b, 10, 50)
}

func BenchmarkCtrl_t50_s01(b *testing.B) {
	benchmarkCtrl(b, 50, 1)
}

func BenchmarkCtrl_t50_s50(b *testing.B) {
	benchmarkCtrl(b, 50, 50)
}

func benchmarkCtrl(b *testing.B, numTasks int, numServices int) {
	// Benchmarks Init and Run for the ReadOnly controller
	//
	// ReadOnlyController.Init involves creating auto-generated Terraform files
	// and the hcat template file for each task.
	//
	// ReadOnlyController.Run involves rendering the template file and executing
	// Terraform init and Terraform plan serially across all tasks.

	basePath, err := filepath.Abs("../../testutils")
	if err != nil {
		b.Fatalf("unable to get current working directory for Consul test certs: %s", err)
	}

	srv := testutils.NewTestConsulServerHTTPS(b, basePath)
	defer srv.Stop()

	tempDir := b.Name()
	cleanup := testutils.MakeTempDir(b, tempDir)
	defer cleanup()

	ctx := context.Background()
	conf := generateConf(benchmarkConfig{
		consulAddr:  srv.HTTPSAddr,
		certPath:    filepath.Join(basePath, "cert.pem"),
		tempDir:     tempDir,
		numTasks:    numTasks,
		numServices: numServices,
	})

	b.Run("ReadOnlyCtrl", func(b *testing.B) {
		ctrl, err := controller.NewReadOnly(conf)
		if err != nil {
			b.Fatal(err)
		}

		b.Run("task setup", func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_, err = ctrl.Init(ctx)
				if err != nil {
					b.Fatal(err)
				}
			}
		})

		b.Run("task execution", func(b *testing.B) {
			err = ctrl.Run(ctx)
			if err != nil {
				b.Fatal(err)
			}
		})
	})
}

type benchmarkConfig struct {
	consulAddr  string
	certPath    string
	tempDir     string
	numTasks    int
	numServices int
}

func generateConf(bConf benchmarkConfig) *config.Config {
	serviceNames := make([]string, bConf.numServices)
	for i := 0; i < bConf.numServices; i++ {
		serviceNames[i] = fmt.Sprintf("service_%03d", i)
	}

	taskConfigs := make(config.TaskConfigs, bConf.numTasks)
	for i := 0; i < bConf.numTasks; i++ {
		taskConfigs[i] = &config.TaskConfig{
			Name:     config.String(fmt.Sprintf("task_%03d", i)),
			Source:   config.String("../../../test_modules/local_file"),
			Services: serviceNames,
		}
	}

	conf := config.DefaultConfig()
	conf.Tasks = &taskConfigs
	conf.Consul.Address = config.String(bConf.consulAddr)
	conf.Consul.TLS = &config.TLSConfig{
		Enabled: config.Bool(true),
		Verify:  config.Bool(false),

		// This is needed for Terraform Consul backend when CTS is
		// connecting over HTTP/2 using TLS.
		CACert: config.String(bConf.certPath),
	}
	conf.Finalize()
	conf.Driver.Terraform.WorkingDir = config.String(bConf.tempDir)
	conf.Driver.Terraform.Path = config.String("../../../")
	return conf
}
