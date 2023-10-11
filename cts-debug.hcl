log_level   = "INFO"
working_dir = "sync-tasks" port        = 8558

syslog {}

buffer_period {
  enabled = true
  min     = "5s"
  max     = "20s"
}

consul {
  address = "https://13.53.36.60:8501"
  token = "8b3ddad2-7fbb-6fe8-f06b-58c7fe18a40a"
	tls {
	    ca_cert = "/Users/asheshvidyut/consul-terraform-sync/consul-certs/ca-cert.pem"
	}
}

log_level = "debug"

driver "terraform" {
  # version = "0.14.0"
  # path = ""
  log         = false
  persist_log = false

  backend "consul" {
    gzip = true
  }
}

task {
 name        = "learn-cts-example"
 description = "Example task with two services"
 module      = "findkim/print/cts"
 version     = "0.1.0"
 condition "services" {
  datacenter = "dc1"
  names = ["backend"]
 }
}
