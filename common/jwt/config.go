package jwt

var defaultConfig = Config{
	Algorithm: "ES256",
	PublicKey: `-----BEGIN ECC PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE5C3O98FgRVb5WGUQx2W1gJ1WbL3B
OXlgBmKroTjE9BNekOlS6LbwkM5OM7bOoGROuRZ+W8Ur6dvpyzXar1aOzg==
-----END ECC PUBLIC KEY-----`,
	PrivateKey: `-----BEGIN ECC PRIVATE KEY-----
MHcCAQEEILfNf1hbu8Adh+ErwTa7RoD4fPReJBUqYI9V8DEHQbRzoAoGCCqGSM49
AwEHoUQDQgAE5C3O98FgRVb5WGUQx2W1gJ1WbL3BOXlgBmKroTjE9BNekOlS6Lbw
kM5OM7bOoGROuRZ+W8Ur6dvpyzXar1aOzg==
-----END ECC PRIVATE KEY-----`,
}

type Config struct {
	Algorithm string `json:"algorithm" yaml:"algorithm" mapstructure:"algorithm"`
	// TODO: 使用Issuer
	Issuer     string `json:"issuer" yaml:"issuer" mapstructure:"issuer"`
	Secret     string `json:"secret" yaml:"secret" mapstructure:"secret"`
	PublicKey  string `json:"public_key" yaml:"public_key" mapstructure:"public_key"`
	PrivateKey string `json:"private_key" yaml:"private_key" mapstructure:"private_key"`
}
