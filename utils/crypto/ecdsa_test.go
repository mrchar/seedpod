package crypto

import (
	"crypto/x509"
	"encoding/pem"
	"testing"
)

func TestGenerateECDSAKeyPair(t *testing.T) {
	publicKey, privateKey, err := GenerateECDSAKeyPair()
	if err != nil {
		t.Error(err)
	}

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		t.Error(err)
	}

	privateKeyBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return
	}

	t.Logf(
		`PublicKey:\n "%s", PrivateKey:\n "%s"`,
		pem.EncodeToMemory(&pem.Block{
			Type:  "ECC PUBLIC KEY",
			Bytes: publicKeyBytes,
		}),
		pem.EncodeToMemory(&pem.Block{
			Type:  "ECC PRIVATE KEY",
			Bytes: privateKeyBytes,
		}),
	)
}
