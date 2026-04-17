package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/go-jose/go-jose/v3/jwt"
)

// Signer is an interface for signing JWTs.
// It allow for different implementations (e.g., using a local PEM file or delegating to AWS KMS).
type Signer interface {
	SignJWT(ctx context.Context, claims jwt.Claims) (string, error)
}

// GenerateOAuthTokenFromApp generates a GitHub OAuth access token from a set of valid GitHub App credentials.
// The returned token can be used to interact with both GitHub's REST and GraphQL APIs.
func GenerateOAuthTokenFromApp(apiURL *url.URL, appID, appInstallationID, pemData string) (string, error) {
	signer, err := NewPEMSigner([]byte(pemData))
	if err != nil {
		return "", err
	}

	appJWT, err := generateAppJWT(context.Background(), appID, time.Now(), signer)
	if err != nil {
		return "", err
	}

	token, err := getInstallationAccessToken(apiURL, appJWT, appInstallationID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func getInstallationAccessToken(apiURL *url.URL, jwt, installationID string) (string, error) {
	req, err := http.NewRequest(http.MethodPost, apiURL.JoinPath("app/installations", installationID, "access_tokens").String(), nil)
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

func generateAppJWT(ctx context.Context, appID string, now time.Time, signer Signer) (string, error) {
	claims := jwt.Claims{
		Issuer: appID,
		// Using now - 60s to accommodate any client/server clock drift.
		IssuedAt: jwt.NewNumericDate(now.Add(time.Duration(-60) * time.Second)),
		// The JWT's lifetime can be short as it is only used immediately
		// after to retrieve the installation's access token.
		Expiry: jwt.NewNumericDate(now.Add(time.Duration(5) * time.Minute)),
	}

	return signer.SignJWT(ctx, claims)
}
