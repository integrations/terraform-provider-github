package github

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"golang.org/x/oauth2/jws"
	"io/ioutil"
	"strings"
	"testing"
	"time"
)

const (
	testGitHubAppID             string = "123456789"
	testGitHubAppInstallationID string = "987654321"
	testGitHubAppPrivateKeyFile string = "test-fixtures/github-app-key.pem"
	testGitHubAppPublicKeyFile  string = "test-fixtures/github-app-key.pub"
)

var (
	testEpochTime time.Time = time.Unix(0, 0)
)

func TestGenerateAppJWT(t *testing.T) {
	pemData, err := ioutil.ReadFile(testGitHubAppPrivateKeyFile)
	if err != nil {
		t.Logf("Failed to read private key file '%s': %s", testGitHubAppPrivateKeyFile, err)
		t.FailNow()
	}

	jwt, err := generateAppJWT(testGitHubAppID, testEpochTime, pemData)
	t.Log(jwt)
	if err != nil {
		t.Logf("Failed to generate GitHub app JWT: %s", err)
		t.FailNow()
	}

	t.Run("produces a properly shaped jwt", func(t *testing.T) {
		parts := strings.Split(jwt, ".")

		if len(parts) != 3 {
			t.Logf("The produced JWT an invalid JWT token: %s", jwt)
			t.Fail()
		}
	})

	t.Run("produces a jwt with expected claims", func(t *testing.T) {
		claims, err := jws.Decode(jwt)
		if err != nil {
			t.Logf("Failed to decode generated JWT: %s", err)
			t.Fail()
		}

		if claims.Iss != testGitHubAppID {
			t.Logf("Unexpected 'iss' claim - Expected: %s - Found: %s", testGitHubAppID, claims.Iss)
			t.Fail()
		}

		expectedIssuedAt := testEpochTime.Add(time.Duration(-60) * time.Second).Unix()
		if claims.Iat != expectedIssuedAt {
			t.Logf("Unexpected 'iss' claim - Expected: %d - Found: %s", expectedIssuedAt, claims.Iss)
			t.Fail()
		}

		expectedExpiration := testEpochTime.Add(time.Duration(5) * time.Minute).Unix()
		if claims.Exp != expectedExpiration {
			t.Logf("Unexpected 'exp' claim - Expected: %d - Found: %d", expectedExpiration, claims.Exp)
			t.Fail()
		}

		if claims.Sub != "" || claims.Aud != "" || claims.Typ != "" || claims.Scope != "" || claims.Prn != "" {
			t.Logf("Extra claims found in JWT: %+v", claims)
			t.Fail()
		}

		if !t.Failed() && len(claims.PrivateClaims) != 0 {
			t.Logf("Extra claims found in JWT: %+v", claims)
			t.Fail()
		}
	})

	t.Run("produces a verifiable jwt", func(t *testing.T) {
		publicKeyData, err := ioutil.ReadFile(testGitHubAppPublicKeyFile)
		if err != nil {
			t.Logf("Failed to read public key file '%s': %s", testGitHubAppPublicKeyFile, err)
			t.FailNow()
		}

		block, _ := pem.Decode(publicKeyData)
		publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			t.Logf("Failed to decode public key file '%s': %s", testGitHubAppPublicKeyFile, err)
			t.FailNow()
		}

		err = jws.Verify(jwt, publicKey.(*rsa.PublicKey))
		if err != nil {
			t.Logf("Failed to verify JWT's signature: %s", jwt)
			t.Fail()
		}
	})
}

func TestGetInstallationAccessToken(t *testing.T) {
	fakeJWT := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9" +
		".eyJpc3MiOiIxMjM0NTY3ODkiLCJhdWQiOiIiLCJleHAiOjMwMCwiaWF0IjotNjB9" +
		".jpx6AFGoZzHzre79JveY_nyKop11v-bLxLEMvEDrn2wDF9S1FeX-zfTiA6Xi00Akn0Wklj7OYx0wHCvi37aiD4zjp0qPz5i5V7aMrRsWsO6eCzNfY0VLuV6pX8jlAHFfo71SvpdAMWH4in8ty5bNVUMv0NmwWdlHAQ0LLIPSxE4"

	expectedAccessToken := "W+2e/zjiMTweDAr2b35toCF+h29l7NW92rJIPvFrCJQK"

	ts := githubApiMock([]*mockResponse{
		{
			ExpectedUri: fmt.Sprintf("/app/installations/%s/access_tokens", testGitHubAppInstallationID),
			ExpectedHeaders: map[string]string{
				"Accept":        "application/vnd.github.v3+json",
				"Authorization": fmt.Sprintf("Bearer %s", fakeJWT),
			},

			ResponseBody: fmt.Sprintf(`{"token": "%s"}`, expectedAccessToken),
			StatusCode:   201,
		},
	})
	defer ts.Close()

	accessToken, err := getInstallationAccessToken(ts.URL+"/", fakeJWT, testGitHubAppInstallationID)

	if err != nil {
		t.Logf("Unexpected error: %s", err)
		t.Fail()
	}

	if accessToken != expectedAccessToken {
		t.Logf("Unexpected access token - Found: %s - Expected: %s", accessToken, expectedAccessToken)
		t.Fail()
	}
}
