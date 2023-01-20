package kms

import (
	"context"
	"os"
	"testing"

	"github.com/hashicorp/vault/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/polygonid/sh-id-platform/internal/config"
	"github.com/polygonid/sh-id-platform/internal/log"
	"github.com/polygonid/sh-id-platform/internal/providers"
)

var cfg config.KeyStore

type TestKMS struct {
	KMS       *KMS
	VaultCli  *api.Client
	writenIDs []KeyID
	t         testing.TB
}

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func testMain(m *testing.M) int {
	configForTesting, err := config.Load("")
	if err != nil {
		log.Error(context.Background(), "cannot load config", err)
		panic(err)
	}

	cfg = configForTesting.KeyStore
	return m.Run()
}

// TestKMSSetup checks if Vault is available and setup connection to it.
// Also, it registers cleanup function on test complete.
func testKMSSetup(t testing.TB) TestKMS {
	k := TestKMS{t: t}
	var err error

	k.VaultCli, err = providers.NewVaultClient(testVaultConfig(t))
	require.NoError(t, err)

	k.KMS = NewKMS()

	// err = k.KMS.RegisterKeyProvider(KeyTypeEthereum, NewVaultEthProvider(k.VaultCli, KeyTypeEthereum))
	// require.NoError(t, err)

	err = k.KMS.RegisterKeyProvider(KeyTypeBabyJubJub, NewVaultBJJKeyProvider(k.VaultCli, KeyTypeBabyJubJub))
	require.NoError(t, err)

	t.Cleanup(k.Close)
	return k
}

func testVaultConfig(t testing.TB) (vaultAddr, vaultToken string) {
	vaultAddr = cfg.Address
	vaultToken = cfg.Token
	if vaultAddr == "" {
		t.Skip(".vault address is not configured")
	}
	if vaultToken == "" {
		t.Skip(".vault token is not configured")
	}
	return
}

// Close cleans up Vault on test complete.
func (tk *TestKMS) Close() {
	for _, k := range tk.writenIDs {
		_, err := tk.VaultCli.Logical().Delete(absVaultSecretPath(k.ID))
		assert.NoError(tk.t, err)
	}
}
