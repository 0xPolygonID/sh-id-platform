// Package api_admin provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package api_admin

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
)

const (
	BasicAuthScopes = "basicAuth.Scopes"
)

// AuthenticationQrCodeResponse defines model for AuthenticationQrCodeResponse.
type AuthenticationQrCodeResponse struct {
	Body struct {
		CallbackUrl string        `json:"callbackUrl"`
		Reason      string        `json:"reason"`
		Scope       []interface{} `json:"scope"`
	} `json:"body"`
	From string `json:"from"`
	Id   string `json:"id"`
	Thid string `json:"thid"`
	Typ  string `json:"typ"`
	Type string `json:"type"`
}

// GenericErrorMessage defines model for GenericErrorMessage.
type GenericErrorMessage struct {
	Message string `json:"message"`
}

// Health defines model for Health.
type Health map[string]bool

// ImportSchemaRequest defines model for ImportSchemaRequest.
type ImportSchemaRequest struct {
	SchemaType string `json:"schemaType"`
	Url        string `json:"url"`
}

// SayHi defines model for SayHi.
type SayHi struct {
	Message string `json:"message"`
}

// UUIDResponse defines model for UUIDResponse.
type UUIDResponse struct {
	Id string `json:"id"`
}

// SessionID defines model for sessionID.
type SessionID = string

// N400 defines model for 400.
type N400 = GenericErrorMessage

// N500 defines model for 500.
type N500 = GenericErrorMessage

// AuthCallbackTextBody defines parameters for AuthCallback.
type AuthCallbackTextBody = string

// AuthCallbackParams defines parameters for AuthCallback.
type AuthCallbackParams struct {
	// SessionID Session ID
	SessionID *SessionID `form:"sessionID,omitempty" json:"sessionID,omitempty"`
}

// AuthCallbackTextRequestBody defines body for AuthCallback for text/plain ContentType.
type AuthCallbackTextRequestBody = AuthCallbackTextBody

// ImportSchemaJSONRequestBody defines body for ImportSchema for application/json ContentType.
type ImportSchemaJSONRequestBody = ImportSchemaRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get the documentation
	// (GET /)
	GetDocumentation(w http.ResponseWriter, r *http.Request)
	// Say Hi endpoint
	// (GET /say-hi)
	SayHi(w http.ResponseWriter, r *http.Request)
	// Get the documentation yaml file
	// (GET /static/docs/api_admin/api.yaml)
	GetYaml(w http.ResponseWriter, r *http.Request)
	// Healthcheck
	// (GET /status)
	Health(w http.ResponseWriter, r *http.Request)
	// authentication callback
	// (POST /v1/authentication/callback)
	AuthCallback(w http.ResponseWriter, r *http.Request, params AuthCallbackParams)
	// get authentication qrcode
	// (GET /v1/authentication/qrcode)
	AuthQRCode(w http.ResponseWriter, r *http.Request)
	// Import JSON schema
	// (POST /v1/schemas)
	ImportSchema(w http.ResponseWriter, r *http.Request)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetDocumentation operation middleware
func (siw *ServerInterfaceWrapper) GetDocumentation(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetDocumentation(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// SayHi operation middleware
func (siw *ServerInterfaceWrapper) SayHi(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.SayHi(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetYaml operation middleware
func (siw *ServerInterfaceWrapper) GetYaml(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetYaml(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// Health operation middleware
func (siw *ServerInterfaceWrapper) Health(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.Health(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// AuthCallback operation middleware
func (siw *ServerInterfaceWrapper) AuthCallback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params AuthCallbackParams

	// ------------- Optional query parameter "sessionID" -------------

	err = runtime.BindQueryParameter("form", true, false, "sessionID", r.URL.Query(), &params.SessionID)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "sessionID", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AuthCallback(w, r, params)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// AuthQRCode operation middleware
func (siw *ServerInterfaceWrapper) AuthQRCode(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AuthQRCode(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// ImportSchema operation middleware
func (siw *ServerInterfaceWrapper) ImportSchema(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.ImportSchema(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshallingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshallingParamError) Error() string {
	return fmt.Sprintf("Error unmarshalling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshallingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
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
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/", wrapper.GetDocumentation)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/say-hi", wrapper.SayHi)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/static/docs/api_admin/api.yaml", wrapper.GetYaml)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/status", wrapper.Health)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/v1/authentication/callback", wrapper.AuthCallback)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/v1/authentication/qrcode", wrapper.AuthQRCode)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/v1/schemas", wrapper.ImportSchema)
	})

	return r
}

type N400JSONResponse GenericErrorMessage

type N500JSONResponse GenericErrorMessage

type GetDocumentationRequestObject struct {
}

type GetDocumentationResponseObject interface {
	VisitGetDocumentationResponse(w http.ResponseWriter) error
}

type GetDocumentation200Response struct {
}

func (response GetDocumentation200Response) VisitGetDocumentationResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type SayHiRequestObject struct {
}

type SayHiResponseObject interface {
	VisitSayHiResponse(w http.ResponseWriter) error
}

type SayHi200JSONResponse SayHi

func (response SayHi200JSONResponse) VisitSayHiResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type SayHi500JSONResponse struct{ N500JSONResponse }

func (response SayHi500JSONResponse) VisitSayHiResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type GetYamlRequestObject struct {
}

type GetYamlResponseObject interface {
	VisitGetYamlResponse(w http.ResponseWriter) error
}

type GetYaml200Response struct {
}

func (response GetYaml200Response) VisitGetYamlResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type HealthRequestObject struct {
}

type HealthResponseObject interface {
	VisitHealthResponse(w http.ResponseWriter) error
}

type Health200JSONResponse Health

func (response Health200JSONResponse) VisitHealthResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type Health500JSONResponse struct{ N500JSONResponse }

func (response Health500JSONResponse) VisitHealthResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type AuthCallbackRequestObject struct {
	Params AuthCallbackParams
	Body   *AuthCallbackTextRequestBody
}

type AuthCallbackResponseObject interface {
	VisitAuthCallbackResponse(w http.ResponseWriter) error
}

type AuthCallback200Response struct {
}

func (response AuthCallback200Response) VisitAuthCallbackResponse(w http.ResponseWriter) error {
	w.WriteHeader(200)
	return nil
}

type AuthCallback400JSONResponse struct{ N400JSONResponse }

func (response AuthCallback400JSONResponse) VisitAuthCallbackResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type AuthCallback500JSONResponse struct{ N500JSONResponse }

func (response AuthCallback500JSONResponse) VisitAuthCallbackResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type AuthQRCodeRequestObject struct {
}

type AuthQRCodeResponseObject interface {
	VisitAuthQRCodeResponse(w http.ResponseWriter) error
}

type AuthQRCode200JSONResponse AuthenticationQrCodeResponse

func (response AuthQRCode200JSONResponse) VisitAuthQRCodeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type AuthQRCode500JSONResponse struct{ N500JSONResponse }

func (response AuthQRCode500JSONResponse) VisitAuthQRCodeResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

type ImportSchemaRequestObject struct {
	Body *ImportSchemaJSONRequestBody
}

type ImportSchemaResponseObject interface {
	VisitImportSchemaResponse(w http.ResponseWriter) error
}

type ImportSchema201JSONResponse UUIDResponse

func (response ImportSchema201JSONResponse) VisitImportSchemaResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type ImportSchema400JSONResponse struct{ N400JSONResponse }

func (response ImportSchema400JSONResponse) VisitImportSchemaResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type ImportSchema500JSONResponse struct{ N500JSONResponse }

func (response ImportSchema500JSONResponse) VisitImportSchemaResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Get the documentation
	// (GET /)
	GetDocumentation(ctx context.Context, request GetDocumentationRequestObject) (GetDocumentationResponseObject, error)
	// Say Hi endpoint
	// (GET /say-hi)
	SayHi(ctx context.Context, request SayHiRequestObject) (SayHiResponseObject, error)
	// Get the documentation yaml file
	// (GET /static/docs/api_admin/api.yaml)
	GetYaml(ctx context.Context, request GetYamlRequestObject) (GetYamlResponseObject, error)
	// Healthcheck
	// (GET /status)
	Health(ctx context.Context, request HealthRequestObject) (HealthResponseObject, error)
	// authentication callback
	// (POST /v1/authentication/callback)
	AuthCallback(ctx context.Context, request AuthCallbackRequestObject) (AuthCallbackResponseObject, error)
	// get authentication qrcode
	// (GET /v1/authentication/qrcode)
	AuthQRCode(ctx context.Context, request AuthQRCodeRequestObject) (AuthQRCodeResponseObject, error)
	// Import JSON schema
	// (POST /v1/schemas)
	ImportSchema(ctx context.Context, request ImportSchemaRequestObject) (ImportSchemaResponseObject, error)
}

type StrictHandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request, args interface{}) (interface{}, error)

type StrictMiddlewareFunc func(f StrictHandlerFunc, operationID string) StrictHandlerFunc

type StrictHTTPServerOptions struct {
	RequestErrorHandlerFunc  func(w http.ResponseWriter, r *http.Request, err error)
	ResponseErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}}
}

func NewStrictHandlerWithOptions(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc, options StrictHTTPServerOptions) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: options}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
	options     StrictHTTPServerOptions
}

// GetDocumentation operation middleware
func (sh *strictHandler) GetDocumentation(w http.ResponseWriter, r *http.Request) {
	var request GetDocumentationRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetDocumentation(ctx, request.(GetDocumentationRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetDocumentation")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetDocumentationResponseObject); ok {
		if err := validResponse.VisitGetDocumentationResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("Unexpected response type: %T", response))
	}
}

// SayHi operation middleware
func (sh *strictHandler) SayHi(w http.ResponseWriter, r *http.Request) {
	var request SayHiRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.SayHi(ctx, request.(SayHiRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "SayHi")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(SayHiResponseObject); ok {
		if err := validResponse.VisitSayHiResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("Unexpected response type: %T", response))
	}
}

// GetYaml operation middleware
func (sh *strictHandler) GetYaml(w http.ResponseWriter, r *http.Request) {
	var request GetYamlRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetYaml(ctx, request.(GetYamlRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetYaml")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetYamlResponseObject); ok {
		if err := validResponse.VisitGetYamlResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("Unexpected response type: %T", response))
	}
}

// Health operation middleware
func (sh *strictHandler) Health(w http.ResponseWriter, r *http.Request) {
	var request HealthRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.Health(ctx, request.(HealthRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "Health")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(HealthResponseObject); ok {
		if err := validResponse.VisitHealthResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("Unexpected response type: %T", response))
	}
}

// AuthCallback operation middleware
func (sh *strictHandler) AuthCallback(w http.ResponseWriter, r *http.Request, params AuthCallbackParams) {
	var request AuthCallbackRequestObject

	request.Params = params

	data, err := io.ReadAll(r.Body)
	if err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't read body: %w", err))
		return
	}
	body := AuthCallbackTextRequestBody(data)
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.AuthCallback(ctx, request.(AuthCallbackRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "AuthCallback")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(AuthCallbackResponseObject); ok {
		if err := validResponse.VisitAuthCallbackResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("Unexpected response type: %T", response))
	}
}

// AuthQRCode operation middleware
func (sh *strictHandler) AuthQRCode(w http.ResponseWriter, r *http.Request) {
	var request AuthQRCodeRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.AuthQRCode(ctx, request.(AuthQRCodeRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "AuthQRCode")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(AuthQRCodeResponseObject); ok {
		if err := validResponse.VisitAuthQRCodeResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("Unexpected response type: %T", response))
	}
}

// ImportSchema operation middleware
func (sh *strictHandler) ImportSchema(w http.ResponseWriter, r *http.Request) {
	var request ImportSchemaRequestObject

	var body ImportSchemaJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.ImportSchema(ctx, request.(ImportSchemaRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "ImportSchema")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(ImportSchemaResponseObject); ok {
		if err := validResponse.VisitImportSchemaResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("Unexpected response type: %T", response))
	}
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7xYT3PbthP9Kvjh1yMlKnF64c2RO7HauHWs+JBxPZ0VuBJhkwAMgIpZj757BwAlkRLk",
	"Ohq7l0QCFvvnvQfsyk+UyUpJgcIamj1RBRoqtKj9N4PGcCkmZ+5LjoZpriyXgmZ0GrbI5IwmlLuVhxp1",
	"QxMqoEKadc4m1LACK3BObKP8ptVcLOhqtUqoRqOkMOgjfhiN3H9MCovCuo+gVMkZuLDpnXGxnzr+ftI4",
	"pxn9f7otIw27Jv2EAjVnv2gt9QUaAwsMEfuVfIScXOFDjcbSVUJ//u8zmAiLWkBJpqiXqAk6e+rsWkcu",
	"zmltCxS2TeSLHsscr1roPHNaKtSWBxxnMm/2VxmU5QzY/bUuI2Q4KqCtb2/LMKmwswNaQ0MDfw8115jT",
	"7KYXYONuffg2WR+WsztkHu25llU0Hs+jy7Y4tNGoQ+sYE14/b57T4KI90AZKAo5tmrH8YwzvwV5tN/AR",
	"KlU6H1NZoS24WJAClEJBk3/Jce0llsY5QmkLL9Y8504gUF72cmiPzKQsEcQWmbWThD4OZMUtVso2NJtD",
	"aXCV0EmlpLZTL8P1JdmrL6j0a4v0tsQlMMaF1+vYWc+deHG/0ITWQZDbo4W1ymRpquH7cMFtUc9qg7q9",
	"lEMmq5TnKE5SVgKvBiGBwVIymKUVcLG5ge6+pr99G58ucKwxdxcIysHyZHgXlPk85LWXcae8hBZgCqcL",
	"vpgIG+ViCs05f6EIzvn/juf9AGXX15Ozww9D7PpEHe3fkP1i3buArNbcNl4i7dsDhjP3XG204ZXnVrfF",
	"OoLDW8jFXO73lzPJ6gqF9eohc6mJLZBMjKlRkwG5npDTy8mfnkJuPZaXsmwWviORwa4hTegStQmuR8N3",
	"w5FjSioUoDjN6IlfSqgCW/gaUvfPAr3YHXw+jUlOM/oJbS83utPB3of+0a/G1IyhMQRETjTaWgvj68l7",
	"VXJBzr9efHbVVtDCW1cV6CbE3T/iqeNt+6CZ1bUnLjXQDAp+sIYg0Xjir9L4QoBIq5tCQwpOUORKctFt",
	"uDF3m/xSZ9SVG81uekK7uV3ddtFycc67cTwqDjSW5pKZFBT/C/KKC/dp2EBVPsf4N7f/qkQ7jz9ANGm8",
	"PS/xIOUWbG0OFtF2iDfkvI0QIf20LIlBveQMDQGNRNdCtC//j7G/ASkEYwWy+8Dt8l0KvQkpXc8i/g2U",
	"oW/t5NU7QDYHkh3snN14u9mdkW/iqW9N0u0c7BSqQw/92I5nHeAtPtpUlcB3IN82i7vvfw+svI+OCr3X",
	"2mviJWqV946CDy+hwBkdTxccRNrCwsHoLejtASofNJM5drT9LI+tdYzFL1fjsPVmt+DZMT1yNyCa+/FI",
	"L9ASOIBHHOvOb4z1NekD150B6XMSPh612Jj5IlG/e7UUemNTrHN5O8J9ppi/7cV5YZcLsJFfp3/8Ttoq",
	"V+G8+x0Znqd+FZ8lA9fJ/MDth7AsTUu3WEhjs5PR6L1/p1qp7B4fy7JE5lUl55vuaojGEizmxEoy8UO2",
	"7fwlYLOySo7wN3YTvtl689+Pc3UhZ6GBtq5OF041q9vVPwEAAP//PrUnbAYRAAA=",
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
