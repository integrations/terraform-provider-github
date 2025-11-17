package github

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/go-jose/go-jose/v3"
	"github.com/go-jose/go-jose/v3/jwt"
)

// GenerateOAuthTokenFromApp generates a GitHub OAuth access token from a set of valid GitHub App credentials.
// The returned token can be used to interact with both GitHub's REST and GraphQL APIs.
func GenerateOAuthTokenFromApp(baseURL, appID, appInstallationID, pemData string) (string, error) {
	appJWT, err := generateAppJWT(appID, time.Now(), []byte(pemData))
	if err != nil {
		return "", err
	}

	token, err := getInstallationAccessToken(baseURL, appJWT, appInstallationID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func getInstallationAccessToken(baseURL, jwt, installationID string) (string, error) {
	if baseURL != "https://api.github.com/" && !GHECDataResidencyMatch.MatchString(baseURL) {
		baseURL += "api/v3/"
	}

	url := fmt.Sprintf("%sapp/installations/%s/access_tokens", baseURL, installationID)

	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", jwt))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer func() { _ = res.Body.Close() }()

	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("failed to create OAuth token from GitHub App: %s", string(resBytes))
	}

	resData := struct {
		Token string `json:"token"`
	}{}

	err = json.Unmarshal(resBytes, &resData)
	if err != nil {
		return "", err
	}

	return resData.Token, nil
}

func generateAppJWT(appID string, now time.Time, pemData []byte) (string, error) {
	block, _ := pem.Decode(pemData)
	if block == nil {
		return "", errors.New("no decodeable PEM data found")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	signer, err := jose.NewSigner(
		jose.SigningKey{Algorithm: jose.RS256, Key: privateKey},
		(&jose.SignerOptions{}).WithType("JWT"),
	)
	if err != nil {
		return "", err
	}

	claims := &jwt.Claims{
		Issuer: appID,
		// Using now - 60s to accommodate any client/server clock drift.
		IssuedAt: jwt.NewNumericDate(now.Add(time.Duration(-60) * time.Second)),
		// The JWT's lifetime can be short as it is only used immediately
		// after to retrieve the installation's access  token.
		Expiry: jwt.NewNumericDate(now.Add(time.Duration(5) * time.Minute)),
	}

	token, err := jwt.Signed(signer).Claims(claims).CompactSerialize()
	if err != nil {
		return "", err
	}

	return token, nil
}
