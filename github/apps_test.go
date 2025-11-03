package github

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/go-jose/go-jose/v3"
	"github.com/go-jose/go-jose/v3/jwt"
)

const (
	testGitHubAppID             string = "123456789"
	testGitHubAppInstallationID string = "987654321"
	testGitHubAppPublicKeyFile  string = "test-fixtures/github-app-key.pub"
	testGitHubAppPrivateKeyFile string = "test-fixtures/github-app-key.pem"
)

var (
	testEpochTime = time.Unix(0, 0)

	testGitHubAppPrivateKeyPemData, _ = os.ReadFile(testGitHubAppPrivateKeyFile)
)

func TestGenerateAppJWT(t *testing.T) {
	appJWT, err := generateAppJWT(testGitHubAppID, testEpochTime, testGitHubAppPrivateKeyPemData)
	t.Log(appJWT)
	if err != nil {
		t.Logf("Failed to generate GitHub app JWT: %s", err)
		t.FailNow()
	}

	t.Run("produces a properly shaped jwt", func(t *testing.T) {
		parts := strings.Split(appJWT, ".")

		if len(parts) != 3 {
			t.Logf("Failed to produce a properly shaped jwt token: '%s'", appJWT)
			t.Fail()
		}
	})

	t.Run("produces a jwt with expected algorithm and type", func(t *testing.T) {
		tok, err := jwt.ParseSigned(appJWT)
		if err != nil {
			t.Logf("Failed to decode JWT '%s': %s", appJWT, err)
			t.Fail()
		}

		if len(tok.Headers) != 1 {
			t.Logf("Failed to decode JWT '%s': multiple header entries were found", appJWT)
			t.FailNow()
		}

		headers := tok.Headers[0]

		expectedAlgorithm := string(jose.RS256)
		if headers.Algorithm != expectedAlgorithm {
			t.Logf("The generated JWT '%s' does not use the expected algorithm - Expected: %s - Found: %s", appJWT, expectedAlgorithm, headers.Algorithm)
			t.Fail()
		}

		if value, ok := headers.ExtraHeaders[jose.HeaderType]; !ok || value != "JWT" {
			t.Logf("The generated JWT '%s' does not contain the expected 'typ' header or its value isn't set to 'JWT'", appJWT)
			t.Fail()
		}
	})

	t.Run("produces a jwt with expected claims", func(t *testing.T) {
		tok, err := jwt.ParseSigned(appJWT)
		if err != nil {
			t.Logf("Failed to decode JWT '%s': %s", appJWT, err)
			t.Fail()
		}

		claims := &jwt.Claims{}
		err = tok.UnsafeClaimsWithoutVerification(claims)
		if err != nil {
			t.Logf("Failed to extract claims from JWT '%s': %s", appJWT, err)
			t.Fail()
		}

		if claims.Issuer != testGitHubAppID {
			t.Logf("Unexpected 'iss' claim - Expected: %s - Found: %s", testGitHubAppID, claims.Issuer)
			t.Fail()
		}

		expectedIssuedAt := testEpochTime.Add(time.Duration(-60) * time.Second)
		if claims.IssuedAt.Time() != expectedIssuedAt {
			t.Logf("Unexpected 'iss' claim - Expected: %d - Found: %d", expectedIssuedAt.Unix(), claims.IssuedAt)
			t.Fail()
		}

		expectedExpiration := testEpochTime.Add(time.Duration(5) * time.Minute)
		if claims.Expiry.Time() != expectedExpiration {
			t.Logf("Unexpected 'exp' claim - Expected: %d - Found: %d", expectedExpiration.Unix(), claims.Expiry)
			t.Fail()
		}

		if claims.Subject != "" || claims.Audience != nil || claims.ID != "" || claims.NotBefore != nil {
			t.Logf("Extra claims found in JWT: %+v", claims)
			t.Fail()
		}
	})

	t.Run("produces a verifiable jwt", func(t *testing.T) {
		publicKeyData, err := os.ReadFile(testGitHubAppPublicKeyFile)
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

		tok, err := jwt.ParseSigned(appJWT)
		if err != nil {
			t.Logf("Failed to decode JWT '%s': %s", appJWT, err)
			t.Fail()
		}

		claims := &jwt.Claims{}
		err = tok.Claims(publicKey.(*rsa.PublicKey), claims)
		if err != nil {
			t.Logf("Failed to decode JWT '%s': %s", appJWT, err)
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
			ExpectedUri: fmt.Sprintf("/api/v3/app/installations/%s/access_tokens", testGitHubAppInstallationID),
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
