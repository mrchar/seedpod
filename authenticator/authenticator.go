package authenticator

import (
	"github.com/mrchar/seedpod/account"
	"github.com/mrchar/seedpod/application"
	"github.com/mrchar/seedpod/common/jwt"
	"github.com/pkg/errors"
	"time"
)

var defaultAuthenticator *Authenticator

func DefaultAuthenticator() *Authenticator {
	if defaultAuthenticator == nil {
		defaultAuthenticator = NewAuthenticator(
			jwt.Default(),
			account.DefaultManager(),
			application.DefaultManager(),
		)
	}

	return defaultAuthenticator
}

func NewAuthenticator(
	iver *jwt.IssueVerifier,
	accountManager *account.Manager,
	applicationManager *application.Manager,
) *Authenticator {
	return &Authenticator{
		iver:               iver,
		accountManager:     accountManager,
		applicationManager: applicationManager,
	}
}

type Authenticator struct {
	iver               *jwt.IssueVerifier
	accountManager     *account.Manager
	applicationManager *application.Manager
}

type AuthClaims struct {
	Issuer    string `json:"iss"`
	Audience  string `json:"aud"`
	OpenId    string `json:"sub"`
	ExpiresAt int64  `json:"exp"`
}

func (a AuthClaims) Valid() error {
	panic("implement me")
}

// Auth 验证身份
func (a *Authenticator) Auth(accountName, appId string) (string, error) {
	appAct, err := a.applicationManager.GetAccountByAppIdAndAccountName(appId, accountName)
	if err != nil {
		return "", errors.Wrap(err, "获取ApplicationAccount失败")
	}

	token, err := a.iver.Issue(AuthClaims{
		Issuer:    "seedpod",
		Audience:  appId,
		OpenId:    appAct.OpenId,
		ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
	})
	if err != nil {
		return "", err
	}

	return token, nil
}
