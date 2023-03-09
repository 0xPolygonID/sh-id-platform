package api_admin

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	core "github.com/iden3/go-iden3-core"
	"github.com/iden3/iden3comm"

	"github.com/polygonid/sh-id-platform/internal/config"
	"github.com/polygonid/sh-id-platform/internal/core/ports"
	"github.com/polygonid/sh-id-platform/internal/health"
	"github.com/polygonid/sh-id-platform/internal/log"
)

// Server implements StrictServerInterface and holds the implementation of all API controllers
// This is the glue to the API autogenerated code
type Server struct {
	cfg              *config.Configuration
	identityService  ports.IdentityService
	claimService     ports.ClaimsService
	schemaService    ports.SchemaAdminService
	publisherGateway ports.Publisher
	packageManager   *iden3comm.PackageManager
	health           *health.Status
}

// NewServer is a Server constructor
func NewServer(cfg *config.Configuration, identityService ports.IdentityService, claimsService ports.ClaimsService, schemaService ports.SchemaAdminService, publisherGateway ports.Publisher, packageManager *iden3comm.PackageManager, health *health.Status) *Server {
	return &Server{
		cfg:              cfg,
		identityService:  identityService,
		claimService:     claimsService,
		schemaService:    schemaService,
		publisherGateway: publisherGateway,
		packageManager:   packageManager,
		health:           health,
	}
}

// Health is a method
func (s *Server) Health(_ context.Context, _ HealthRequestObject) (HealthResponseObject, error) {
	var resp Health200JSONResponse = s.health.Status()

	return resp, nil
}

// ImportSchema is the UI endpoint to import schema metadata
func (s *Server) ImportSchema(ctx context.Context, request ImportSchemaRequestObject) (ImportSchemaResponseObject, error) {
	req := request.Body
	if _, err := guardImportSchemaReq(req); err != nil {
		return ImportSchema400JSONResponse{N400JSONResponse{Message: fmt.Sprint("bad request: %w", err.Error())}}, nil
	}
	isuerDID, err := core.ParseDID(s.cfg.Admin.IssuerDID)
	if err != nil {
		return ImportSchema500JSONResponse{N500JSONResponse{Message: err.Error()}}, nil
	}
	schema, err := s.schemaService.ImportSchema(ctx, *isuerDID, req.Url, req.SchemaType)
	if err != nil {
		return ImportSchema500JSONResponse{N500JSONResponse{Message: err.Error()}}, nil
	}
	return ImportSchema201JSONResponse{Id: schema.ID.String()}, nil
}

func guardImportSchemaReq(req *ImportSchemaJSONRequestBody) (core.SchemaHash, error) {
	if req != nil {
		return core.SchemaHash{}, errors.New("empty body")
	}
	if strings.TrimSpace(req.Url) == "" {
		return core.SchemaHash{}, errors.New("empty url")
	}
	if strings.TrimSpace(req.SchemaType) == "" {
		return core.SchemaHash{}, errors.New("empty type")
	}
	if strings.TrimSpace(req.Hash) == "" {
		return core.SchemaHash{}, errors.New("empty hash")
	}
	if strings.TrimSpace(req.BigInt) == "" {
		return core.SchemaHash{}, errors.New("empty bigInt")
	}
	if _, err := url.ParseRequestURI(req.Url); err != nil {
		return core.SchemaHash{}, fmt.Errorf("parsing url: %w", err)
	}
	hash, err := core.NewSchemaHashFromHex(req.Hash)
	if err != nil {
		return core.SchemaHash{}, errors.New("hash wrong format")
	}
	const base10 = 10
	n, ok := new(big.Int).SetString(req.BigInt, base10)
	if !ok {
		return core.SchemaHash{}, errors.New("bigInt wrong format")
	}
	if n.Cmp(hash.BigInt()) != 0 {
		return core.SchemaHash{}, errors.New("hash and bigInt does not match")
	}
	return hash, nil
}

// SayHi - Say Hi
func (s *Server) SayHi(ctx context.Context, request SayHiRequestObject) (SayHiResponseObject, error) {
	return SayHi200JSONResponse{Message: "Hi!"}, nil
}

// GetDocumentation this method will be overridden in the main function
func (s *Server) GetDocumentation(_ context.Context, _ GetDocumentationRequestObject) (GetDocumentationResponseObject, error) {
	return nil, nil
}

// AuthCallback receives the authentication information of a holder
func (s *Server) AuthCallback(ctx context.Context, request AuthCallbackRequestObject) (AuthCallbackResponseObject, error) {
	if request.Params.SessionID == nil || *request.Params.SessionID == "" {
		log.Debug(ctx, "empty sessionID auth-callback request")
		return AuthCallback400JSONResponse{N400JSONResponse{"Cannot proceed with empty sessionID"}}, nil
	}

	if request.Body == nil || *request.Body == "" {
		log.Debug(ctx, "empty request body auth-callback request")
		return AuthCallback400JSONResponse{N400JSONResponse{"Cannot proceed with empty body"}}, nil
	}

	err := s.identityService.Authenticate(ctx, *request.Body, *request.Params.SessionID, s.cfg.APIUI.ServerURL, s.cfg.APIUI.IssuerDID)
	if err != nil {
		log.Debug(ctx, "error authenticating", err.Error())
		return AuthCallback500JSONResponse{}, nil
	}

	return AuthCallback200Response{}, nil
}

// AuthQRCode returns the qr code for authenticating a user
func (s *Server) AuthQRCode(ctx context.Context, _ AuthQRCodeRequestObject) (AuthQRCodeResponseObject, error) {
	qrCode, err := s.identityService.CreateAuthenticationQRCode(ctx, s.cfg.APIUI.ServerURL, s.cfg.APIUI.IssuerDID)
	if err != nil {
		return AuthQRCode500JSONResponse{N500JSONResponse{"Unexpected error while creating qr code"}}, nil
	}

	return AuthQRCode200JSONResponse{
		Body: struct {
			CallbackUrl string        `json:"callbackUrl"`
			Reason      string        `json:"reason"`
			Scope       []interface{} `json:"scope"`
		}{
			qrCode.Body.CallbackURL,
			qrCode.Body.Reason,
			[]interface{}{},
		},
		From: qrCode.From,
		Id:   qrCode.ID,
		Thid: qrCode.ThreadID,
		Typ:  string(qrCode.Typ),
		Type: string(qrCode.Type),
	}, nil
}

// GetYaml this method will be overridden in the main function
func (s *Server) GetYaml(_ context.Context, _ GetYamlRequestObject) (GetYamlResponseObject, error) {
	return nil, nil
}

// RegisterStatic add method to the mux that are not documented in the API.
func RegisterStatic(mux *chi.Mux) {
	mux.Get("/", documentation)
	mux.Get("/static/docs/api_ui/api.yaml", swagger)
}

func documentation(w http.ResponseWriter, _ *http.Request) {
	writeFile("api_ui/spec.html", w)
}

func swagger(w http.ResponseWriter, _ *http.Request) {
	writeFile("api_ui/api.yaml", w)
}

func writeFile(path string, w http.ResponseWriter) {
	f, err := os.ReadFile(path)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("not found"))
	}
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(f)
}
