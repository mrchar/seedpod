package jwt

import (
	"crypto/x509"
	"encoding/pem"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)

var defaultIssueVerifier *IssueVerifier

type IssueVerifier struct {
	config        Config
	signingMethod jwt.SigningMethod
	secret        interface{}
	publicKey     interface{}
	privateKey    interface{}
}

func Default() *IssueVerifier {
	if defaultIssueVerifier == nil {
		err := viper.UnmarshalKey("jwt", &defaultConfig)
		if err != nil {
			logrus.Panic(err)
		}

		defaultIssueVerifier, err = New(defaultConfig)
		if err != nil {
			logrus.Panic(err)
		}
	}
	return defaultIssueVerifier
}

func New(config Config) (*IssueVerifier, error) {
	issueVerifier := &IssueVerifier{
		config: config,
	}

	if err := issueVerifier.init(); err != nil {
		return nil, err
	}
	return issueVerifier, nil
}

func (i *IssueVerifier) Issue(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(i.signingMethod, claims)
	return token.SignedString(i.signatureKey())
}

func (i *IssueVerifier) Verify(tokenString string, claims jwt.Claims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, claims, i.verificationKey)
}

func (i *IssueVerifier) init() error {
	signingMethod := jwt.GetSigningMethod(strings.ToUpper(i.config.Algorithm))
	if signingMethod == nil {
		logrus.Panic(errors.Errorf("不支持的签名算法: %s", i.config.Algorithm))
	}
	i.signingMethod = signingMethod

	switch i.config.Algorithm {
	case "HS256", "HS384", "HS512":
		i.secret = []byte(i.config.Secret)
	case "RS256", "RS384", "RS512":
		pubBlock, _ := pem.Decode([]byte(i.config.PublicKey))
		publicKey, err := x509.ParsePKCS1PublicKey(pubBlock.Bytes)
		if err != nil {
			return err
		}
		i.publicKey = publicKey

		priBlock, _ := pem.Decode([]byte(i.config.PrivateKey))
		privateKey, err := x509.ParsePKCS1PrivateKey(priBlock.Bytes)
		if err != nil {
			return err
		}
		i.privateKey = privateKey

	case "ES256", "ES384", "ES512":
		pubBlock, _ := pem.Decode([]byte(i.config.PublicKey))
		publicKey, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
		if err != nil {
			return err
		}
		i.publicKey = publicKey

		priBlock, _ := pem.Decode([]byte(i.config.PrivateKey))
		privateKey, err := x509.ParseECPrivateKey(priBlock.Bytes)
		if err != nil {
			return err
		}
		i.privateKey = privateKey
	default:
		errors.Errorf("不支持的签名算法: %s", i.signingMethod.Alg())
	}
	return nil
}

func (i *IssueVerifier) signatureKey() interface{} {
	var signatureKey interface{}
	switch i.signingMethod.Alg() {
	case "HS256", "HS384", "HS512":
		signatureKey = i.secret
	case "RS256", "RS384", "RS512", "ES256", "ES384", "ES512":
		signatureKey = i.privateKey
	default:
		logrus.Panic(errors.Errorf("不支持的签名算法: %s", i.signingMethod.Alg()))
	}

	return signatureKey
}

func (i *IssueVerifier) verificationKey(token *jwt.Token) (interface{}, error) {
	var verificationKey interface{}
	switch i.signingMethod.Alg() {
	case "HS256", "HS384", "HS512":
		verificationKey = i.secret
	case "RS256", "RS384", "RS512", "ES256", "ES384", "ES512":
		verificationKey = i.publicKey
	default:
		return nil, errors.Errorf("不支持的签名算法: %s", i.signingMethod.Alg())
	}

	return verificationKey, nil
}
