package local

type NameAndPasswordCredential struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type JWTCredential struct {
	Token string `json:"token"`
}
