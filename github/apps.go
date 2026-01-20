package github

import (
	"bytes"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/go-jose/go-jose/v3"
	"github.com/go-jose/go-jose/v3/jwt"
)

// GenerateOAuthTokenFromApp generates a GitHub OAuth access token from a set of valid GitHub App credentials.
// The returned token can be used to interact with both GitHub's REST and GraphQL APIs.
func GenerateOAuthTokenFromApp(apiURL *url.URL, appID, appInstallationID, pemData string) (string, error) {
	return GenerateOAuthTokenFromAppWithRepositories(apiURL, appID, appInstallationID, pemData, nil)
}

// GenerateOAuthTokenFromAppWithRepositories generates a GitHub OAuth access token from a set of valid GitHub App credentials,
// optionally scoped to specific repositories. If repositories is nil or empty, the token will have access to all
// repositories the installation has access to.
func GenerateOAuthTokenFromAppWithRepositories(apiURL *url.URL, appID, appInstallationID, pemData string, repositories []string) (string, error) {
	appJWT, err := generateAppJWT(appID, time.Now(), []byte(pemData))
	if err != nil {
		return "", err
	}

	token, err := getInstallationAccessToken(apiURL, appJWT, appInstallationID, repositories)
	if err != nil {
		return "", err
	}

	return token, nil
}

func getInstallationAccessToken(apiURL *url.URL, jwt, installationID string, repositories []string) (string, error) {
	var reqBody io.Reader
	if len(repositories) > 0 {
		body := struct {
			Repositories []string `json:"repositories"`
		}{
			Repositories: repositories,
		}
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return "", err
		}
		reqBody = bytes.NewReader(bodyBytes)
	}

	req, err := http.NewRequest(http.MethodPost, apiURL.JoinPath("app/installations", installationID, "access_tokens").String(), reqBody)
	if err != nil {
		return "", err
	}

	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", jwt))
	if len(repositories) > 0 {
		req.Header.Add("Content-Type", "application/json")
	}

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
