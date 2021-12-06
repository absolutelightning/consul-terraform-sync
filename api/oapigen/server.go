// Package oapigen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.3 DO NOT EDIT.
package oapigen

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Triggers a dryrun for a provided task
	// (POST /v1/dryrun/task)
	ExecuteTaskDryrun(w http.ResponseWriter, r *http.Request)
	// Creates a new task
	// (POST /v1/tasks)
	CreateTask(w http.ResponseWriter, r *http.Request, params CreateTaskParams)
	// Deletes a task by name
	// (DELETE /v1/tasks/{name})
	DeleteTaskByName(w http.ResponseWriter, r *http.Request, name string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// ExecuteTaskDryrun operation middleware
func (siw *ServerInterfaceWrapper) ExecuteTaskDryrun(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ExecuteTaskDryrun(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// CreateTask operation middleware
func (siw *ServerInterfaceWrapper) CreateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params CreateTaskParams

	// ------------- Optional query parameter "run" -------------
	if paramValue := r.URL.Query().Get("run"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "run", r.URL.Query(), &params.Run)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter run: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateTask(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// DeleteTaskByName operation middleware
func (siw *ServerInterfaceWrapper) DeleteTaskByName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "name" -------------
	var name string

	err = runtime.BindStyledParameter("simple", false, "name", chi.URLParam(r, "name"), &name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter name: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteTaskByName(w, r, name)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL     string
	BaseRouter  chi.Router
	Middlewares []MiddlewareFunc
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
	}

	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/v1/dryrun/task", wrapper.ExecuteTaskDryrun)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/v1/tasks", wrapper.CreateTask)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/v1/tasks/{name}", wrapper.DeleteTaskByName)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}