package local

import (
	"encoding/base64"
	"github.com/mrchar/seedpod/account"
	"github.com/mrchar/seedpod/common/jwt"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var defaultProvider *Provider

type Provider struct {
	accountManager *account.Manager
	issueVerifier  *jwt.IssueVerifier
}

func Default() *Provider {
	if defaultProvider == nil {
		defaultProvider = New(
			account.DefaultManager(),
			jwt.Default(),
		)
	}

	return defaultProvider
}

func New(accountManager *account.Manager, issueVerifier *jwt.IssueVerifier) *Provider {
	return &Provider{
		accountManager: accountManager,
		issueVerifier:  issueVerifier,
	}
}

// Register 注册账户
func (p *Provider) Register(credential NameAndPasswordCredential) error {
	return p.accountManager.RegisterAccount(credential.Name, credential.Password)
}

type LoginClaims struct {
	Issuer    string   `json:"iss"`
	Name      string   `json:"sub"`
	Roles     []string `json:"roles"`
	ExpiresAt int64    `json:"exp"`
}

func (c *LoginClaims) Valid() error {
	if c.Issuer != "seedpod" {
		return errors.New("使用了错误的令牌")
	}
	if c.ExpiresAt < time.Now().Unix() {
		return errors.New("令牌已经过期")
	}
	return nil
}

func (p *Provider) Login(credential NameAndPasswordCredential) (*JWTCredential, error) {
	token, err := p.login(credential.Name, credential.Password)
	if err != nil {
		return nil, err
	}

	return &JWTCredential{Token: token}, nil
}

// Login 登录账户
func (p *Provider) login(name, password string) (string, error) {
	account, err := p.accountManager.GetAccount(name)
	if err != nil {
		return "", errors.Wrap(err, "用户名或密码错误")
	}

	hashedPassword, err := base64.RawStdEncoding.DecodeString(account.Password)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		return "", errors.Wrap(err, "用户名或密码错误")
	}

	// TODO: issuer
	token, err := p.issueVerifier.Issue(&LoginClaims{
		Issuer: "seedpod",
		Name:   account.Name,
		Roles: func() []string {
			var roles []string
			for _, role := range account.Roles {
				roles = append(roles, role.Name)
			}
			return roles
		}(),
		ExpiresAt: time.Now().Add(8 * time.Hour).Unix(),
	})
	if err != nil {
		return "", err
	}

	return token, nil
}
