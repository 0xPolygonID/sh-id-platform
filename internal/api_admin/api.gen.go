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
	uuid "github.com/google/uuid"
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

// GenericMessage defines model for GenericMessage.
type GenericMessage struct {
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

// Id defines model for id.
type Id = uuid.UUID

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
	// delete connection
	// (DELETE /v1/connections/{id})
	DeleteConnection(w http.ResponseWriter, r *http.Request, id Id)
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

// DeleteConnection operation middleware
func (siw *ServerInterfaceWrapper) DeleteConnection(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id Id

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{""})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteConnection(w, r, id)
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
		r.Delete(options.BaseURL+"/v1/connections/{id}", wrapper.DeleteConnection)
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

type DeleteConnectionRequestObject struct {
	Id Id `json:"id"`
}

type DeleteConnectionResponseObject interface {
	VisitDeleteConnectionResponse(w http.ResponseWriter) error
}

type DeleteConnection200JSONResponse GenericMessage

func (response DeleteConnection200JSONResponse) VisitDeleteConnectionResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type DeleteConnection400JSONResponse struct{ N400JSONResponse }

func (response DeleteConnection400JSONResponse) VisitDeleteConnectionResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type DeleteConnection500JSONResponse struct{ N500JSONResponse }

func (response DeleteConnection500JSONResponse) VisitDeleteConnectionResponse(w http.ResponseWriter) error {
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
	// delete connection
	// (DELETE /v1/connections/{id})
	DeleteConnection(ctx context.Context, request DeleteConnectionRequestObject) (DeleteConnectionResponseObject, error)
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

// DeleteConnection operation middleware
func (sh *strictHandler) DeleteConnection(w http.ResponseWriter, r *http.Request, id Id) {
	var request DeleteConnectionRequestObject

	request.Id = id

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteConnection(ctx, request.(DeleteConnectionRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteConnection")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(DeleteConnectionResponseObject); ok {
		if err := validResponse.VisitDeleteConnectionResponse(w); err != nil {
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

	"H4sIAAAAAAAC/8xYzXLbNhB+FRTtkRKdOL3o5sidWG3SJlF8yLiezgpckbBBgAZAJayH794BQImkSLmK",
	"Ju70YkvAYn++bxe70CNlKi+URGkNnT3SAjTkaFH7bzxxfxM0TPPCciXpzK1FlLtPBdiMRlRCjtt1jQ8l",
	"15jQmdUlRtSwDHNwSmxVOCljNZcpjejXSaomzWJZ8mR6fb247K5PeF4obd3ZxoITo1EwO6Mpt1m5mjKV",
	"x6lSqcDY79d1HVGDxnAlF5dD95dhi3hjPoyHEnXVxtGePey+N6LRFEoa9Ei9Ojtz/5iSFqV3GopCcAbO",
	"bHxnnO3Hjr6fNK7pjP4Yt/DHYdfEb1Ci5uwXrZV+h8ZAisFiP5LXkJCP+FCisbSO6M//vQcLaVFLEGSJ",
	"eoOaoJOnnoCgyNm5KG2G0jaOfNBzleDHBjqfcVoVqC0POK5UUg1XGQixAnZ/rcUIGY4KaOIbbBmmCuzs",
	"gNZQ0cDfNldvegZ26raHb6PtYbW6Q+bRXmuVj9oLJTNYttmhjao4tI5jidf32xeEU9EcaAxFAcfGzTH/",
	"xxgewJ63G/gV8kI4HUuVo824TEkGRYGSRv/i41bLE24c48HpRq4QhLsyHikkCXdZCOJ9z0xzZKWUQJAt",
	"/Fsl7lpSObeYF7aiszUIg3VEF/6GWvpc31biIIRQCp8aOlscN8AYl74o5k567SoEh2hGtAxZ3x7NrC3M",
	"LI41fJmGa7A0qJvK9zciT1Cex0wAzyfBgclGMVjFOXC5K3N3KcS/fZ5fpDjXmLgqBTHZnE/vQvo/DXnp",
	"a6UT3hj2S6iu+JGZdcV/OD2ZDlDkusrh22asJkcVDctuGKzvO6zU3FY+JZoLDQxn7g7c5YLPNLfaBusI",
	"DRcsl2s1bFqXipU5SuuzhayVJjZDsjCmRE0m5HpBLt4v/vSUceuxfK9Elfo2Ryb7gjSiG9QmqD6bvpie",
	"OaZUgRIKTmf03C+FRutjiN2fFH1yO/i8G4uEzugbtD3f6F5bfBmaUj8aUzKGxhCQCdFoSy2NjyfpRckl",
	"ufr07q2LNocG3jLPQVfB7vCIp443PSlMIO5UbKCaZPxgDCFFxx3/Lt00GBjpn0uoSMYJyqRQXHa7+Ji6",
	"nX+xE+qmG53d9BLt5ra+7aLl7Fx17XhUHGgsThQzMRT8L0hyLt2naQW5eIrxz27/uxLtNH4D0aTy8lzg",
	"Qcot2NIcDKLpCM/IeWNhhPQLIYhBveEMDQGNRJdSNjf9t7G/AykYYxmy+8Dt5kUMvbEr3g44/g5UoU/t",
	"+dU7QHYHoj3snNy83ew+GG7GXW9F4na4dhmqQ8983cx8HeAtfrVxIYDvQd42i7svf0+suh+dP+r9h0h9",
	"TLaqe0fBq2MocEKn0wUHkbaQOhi9BL09QOWDZirBTm4/yWMjPcbih4/zsPVsVfDk7D9SGzDq++lIp2gJ",
	"HMBjHGumpETmJE38yJM64CvQ4hDpsE7aIwOUL73EvCvwbfXiHrS3z0jQ3vw9QslzV8WRLWwM6y2DncUd",
	"j50H6Pa661PTnd3pU1fR6eCOPQ+OupxefDcXeuPv2ATi5Uj4qQWT/wXVATby6/KP30kTZfOjjt5sy6Yf",
	"xVvFwE0k/qHkh+lZHAu3mCljZ+dnZy99GTUJs398roQI+UPUejclGaJRgMWEWEUW/nFkOz8T7Vbq6AR9",
	"c/cyM602//00Ve/UKgxCjaqL1GVNfVv/EwAA//8yCv692xMAAA==",
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
