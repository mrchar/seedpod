package provider

// Provider 用于提供身份认证
type Provider interface {
	Register(credential IdentityCredential) error
	Login(credential IdentityCredential) (IdentityCredential, error)
}

type IdentityCredential interface {
}
