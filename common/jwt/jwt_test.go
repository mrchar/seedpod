package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestHmac(t *testing.T) {
	testIssueVerifier(t, Config{
		Algorithm: "HS256",
		Secret:    "MySecret",
	})
}

func TestRSA(t *testing.T) {
	testIssueVerifier(t, Config{
		Algorithm: "RS256",
		PublicKey: `-----BEGIN RSA PUBLIC KEY-----
MIIBCgKCAQEA0SmvkqNSdKosKYr9jpt5KlddtCRqUQSauHBXCjvt/BtPzLLDWqAa
CW+8n+dQugpeazo6atIAao7EGlhJVgsihez2qvrf8LqTBNPCOicNgciDWSXRepvy
SuwTRjCKaCN2Wjxgn4yABnUQ3TRZr8HZJcWeakwWoAWcHeAG6w19jOIderxTkebq
Rq1BkPdBVlALGsFoeIa19mD4JeydtZHO+V4k7KN4KLwPQr1HpVgqTxSPiP2SeDmB
q0lklOuEjYNVJ1WtlSQ2X0hP7LMe4+Y8vz8NAzd3+kbfriA0kvHX0KyA4bMsA7r8
hFWA6BE1CjuaYvsbFOBQ5uxIxbgrm0+TxQIDAQAB
-----END RSA PUBLIC KEY-----`,
		PrivateKey: `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA0SmvkqNSdKosKYr9jpt5KlddtCRqUQSauHBXCjvt/BtPzLLD
WqAaCW+8n+dQugpeazo6atIAao7EGlhJVgsihez2qvrf8LqTBNPCOicNgciDWSXR
epvySuwTRjCKaCN2Wjxgn4yABnUQ3TRZr8HZJcWeakwWoAWcHeAG6w19jOIderxT
kebqRq1BkPdBVlALGsFoeIa19mD4JeydtZHO+V4k7KN4KLwPQr1HpVgqTxSPiP2S
eDmBq0lklOuEjYNVJ1WtlSQ2X0hP7LMe4+Y8vz8NAzd3+kbfriA0kvHX0KyA4bMs
A7r8hFWA6BE1CjuaYvsbFOBQ5uxIxbgrm0+TxQIDAQABAoIBAC0A8CL3+yTY/JmP
q1DEyQRAOgwpjaqS+AtZOJEeJe2JntjnWbslGZUQMqChL7BnzLr7k1gqiSZkQ3N7
rYPS74GrY8OUKRztt/Zg7bh/cJuNHh3PUkw0Q6S1OMxwY0dW4f82YH/TxjQdehxu
SCHV7rf+1j/+0RtrOZT//90RUQM/u/WeusTBbM96EgarI9Mt3WPoUPh4FZVxHsBg
E2gj+4wTu+hMdhKMLHlPJq4W29M600k5kvszSt93MTRm0chDHdAMGQvr1VbtclYx
UyEy5S/A4VXr6nDF4czLTXFrIHdWvKyo1BY/kVoYDFgD9eCt4WUiikQRPRdeOzw7
38n3IsECgYEA8cbb/+jP0tZCwgeg7P+dNtqyux9aw7N1cFDekEYXDOlO6H6X7IcM
9b1Fg6scodhlsdPVOSq3Ze3wSuCO0n+ANj93EgU0M/xCdAVDK0L/LPlFEqbo43nx
V+n6llan4HuJPNSbKiigg4RB9iJi0ITUSRmlgN17ngLcoCygXjtH5DECgYEA3Xep
oz35RW1HA8Me6YdN7tCTC6qFpd/2w853xYCg8XINv2XHKv5lUYUcjmqns1fvbt8+
Ksa489cEhgFD3+mBPYeiBgoaXA6yr/U9ti7x3/XLe2h3UBBG9beZpAgsqp7OVnPu
ZrQ9iWZxk2BapZN6V6Ptfu/awZXO8bLe56XfZ9UCgYAM8BpHuHqeiq6p2WSoKgmM
rOlRkBz4Sfsn1nwCdm23WCjL0jJpCtULtWQp5pcypfTTLkXDuGB2COSJ7ThXVVFU
FdNWWIbxnTclJD7y6rPjATfMBriBq73ZeYDaWKrFHXc7lRj0iZYFU6d/91kYVXNS
shekLLX3v6l0vM6cHEn9kQKBgQCTscvl29yzWk8zyRqCbwOgMmT+MLh0iMoOh1EE
2+V5X7CfZgbPO1ziYr1KlQJF7mz1KdhRurl5lHmlzI4xc44HNL7u/CncHsk343tG
VkRkMY2EPYTkVhaco9bIt9Lh7op4yVPCFo27ZiB0QpvxNEswy1gFgXwIAhpCwiE5
pzs4CQKBgQDBvN/2+kcNk5jgaTJax1/sF4NWUhRjclFu7vOl6iMdp9HSauTY0qUh
BpHkZYk+KPcU74Il8Q+nOtEqGpUEDcbsl98uUMcui/HDjVs9unbO7igNVJn/mlum
kZ3t+GhA3qTfbxGNUMi3zsMQD+/DIDsRhuR9c/fGHMBAt1UYYdhg4w==
-----END RSA PRIVATE KEY-----`,
	})
}

func TestECDSA(t *testing.T) {
	testIssueVerifier(t, Config{
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
	})
}

func testIssueVerifier(t *testing.T, config Config) {
	convey.Convey("创建IssueVerifier", t, func() {
		issueVerifier, err := New(config)
		convey.So(err, convey.ShouldBeNil)

		convey.Convey("创建JWT", func() {
			tokenString, err := issueVerifier.Issue(jwt.MapClaims{
				"foo": "bar",
				"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
				"exp": time.Now().Add(10 * time.Hour).Unix(),
			})

			convey.So(err, convey.ShouldBeNil)
			convey.Convey("验证签名", func() {
				claims := jwt.MapClaims{}
				token, err := issueVerifier.Verify(tokenString, &claims)
				convey.So(err, convey.ShouldBeNil)
				convey.Printf("Token: %+v\n Claims: %+v\n", token, claims)
			})
		})
	})
}
