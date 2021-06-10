package crypto

import (
	"crypto/x509"
	"encoding/pem"
	"testing"
)

func TestGenerateRSAKeyPair(t *testing.T) {
	publicKey, privateKey, err := GenerateRSAKeyPair()
	if err != nil {
		t.Error(err)
	}

	t.Logf(
		"PublicKey:\n %s, PrivateKey:\n %s",
		pem.EncodeToMemory(&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(publicKey),
		}),
		pem.EncodeToMemory(&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		}),
	)
}
