package api

import (
	"context"
	"math/rand"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"

	"github.com/polygonid/sh-id-platform/internal/core/ports"
	"github.com/polygonid/sh-id-platform/internal/log"
)

// Server implements StrictServerInterface and holds the implementation of all API controllers
// This is the glue to the API autogenerated code
type Server struct {
	indentityService ports.IndentityService
}

// NewServer is a Server constructor
func NewServer(indentityService ports.IndentityService) *Server {
	return &Server{
		indentityService: indentityService,
	}
}

// Health is a method
func (s *Server) Health(_ context.Context, _ HealthRequestObject) (HealthResponseObject, error) {
	return Health200JSONResponse{
		Cache: true,
		Db:    false,
	}, nil
}

// Ping is a method
func (s *Server) Ping(ctx context.Context, _ PingRequestObject) (PingResponseObject, error) {
	log.Info(ctx, "ping")
	return Ping201JSONResponse{Response: ToPointer("pong")}, nil
}

// Random is a method
func (s *Server) Random(_ context.Context, _ RandomRequestObject) (RandomResponseObject, error) {
	randomMessages := []string{"might", "rays", "bicycle", "use", "certainly", "chicken", "tie", "rain", "tent"}
	i := rand.Intn(len(randomMessages))
	randomResponses := []RandomResponseObject{
		Random400JSONResponse{N400JSONResponse{Message: &randomMessages[i]}},
		Random401JSONResponse{N401JSONResponse{Message: &randomMessages[i]}},
		Random402JSONResponse{N402JSONResponse{Message: &randomMessages[i]}},
		Random407JSONResponse{N407JSONResponse{Message: &randomMessages[i]}},
		Random500JSONResponse{N500JSONResponse{Message: &randomMessages[i]}},
	}
	return randomResponses[rand.Intn(len(randomResponses))], nil
}

// RegisterStatic add method to the mux that are not documented in the API.
func RegisterStatic(mux *chi.Mux) {
	mux.Get("/", documentation)
	mux.Get("/static/docs/api/api.yaml", swagger)
}

func documentation(w http.ResponseWriter, _ *http.Request) {
	writeFile("api/spec.html", w)
}

func swagger(w http.ResponseWriter, _ *http.Request) {
	writeFile("api/api.yaml", w)
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

// CreateIdentity is create identity controller
func (s *Server) CreateIdentity(ctx context.Context, request CreateIdentityRequestObject) (CreateIdentityResponseObject, error) {
	return nil, nil
}

// CreateClaim is claim creation controller
func (s *Server) CreateClaim(ctx context.Context, request CreateClaimRequestObject) (CreateClaimResponseObject, error) {
	return nil, nil
}

// RevokeClaim is the revocation claim controller
func (s *Server) RevokeClaim(ctx context.Context, request RevokeClaimRequestObject) (RevokeClaimResponseObject, error) {
	return nil, nil
}

// GetRevocationStatus is the controller to get revocation status
func (s *Server) GetRevocationStatus(ctx context.Context, request GetRevocationStatusRequestObject) (GetRevocationStatusResponseObject, error) {
	return nil, nil
}
