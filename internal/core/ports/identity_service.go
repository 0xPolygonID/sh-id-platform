package ports

import (
	"context"

	core "github.com/iden3/go-iden3-core"
	"github.com/iden3/go-schema-processor/verifiable"

	"github.com/polygonid/sh-id-platform/internal/core/domain"
	"github.com/polygonid/sh-id-platform/internal/kms"
)

// IdentityService is the interface implemented by the identity service
type IdentityService interface {
	Create(ctx context.Context, hostURL string) (*domain.Identity, error)
	SignClaimEntry(ctx context.Context, authClaim *domain.Claim, claimEntry *core.Claim) (*verifiable.BJJSignatureProof2021, error)
	Get(ctx context.Context) (identities []string, err error)
	UpdateState(ctx context.Context, did *core.DID) (*domain.IdentityState, error)
	Exists(ctx context.Context, identifier *core.DID) (bool, error)
	GetLatestStateByID(ctx context.Context, identifier *core.DID) (*domain.IdentityState, error)
	GetKeyIDFromAuthClaim(ctx context.Context, authClaim *domain.Claim) (kms.KeyID, error)
	GetUnprocessedIssuersIDs(ctx context.Context) ([]*core.DID, error)
	GetNonTransactedStates(ctx context.Context) ([]domain.IdentityState, error)
	UpdateIdentityState(ctx context.Context, state *domain.IdentityState) error
	GetTransactedStates(ctx context.Context) ([]domain.IdentityState, error)
}
