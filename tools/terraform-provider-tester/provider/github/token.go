package github

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	gogithub "github.com/google/go-github/v88/github"
)

type authMethod int

const (
	authNone  authMethod = iota
	authToken            //nolint:deadcode
	authApp
)

type credInfo struct {
	Method      authMethod
	Login       string
	Scopes      []string
	FineGrained bool
	Token       string
}

// detectAuth determines the authentication method from env, mirroring acc_test.go lines 167-180.
// It performs no network calls.
func detectAuth(env func(string) string) (authMethod, error) {
	token := env("GITHUB_TOKEN")
	appID := env("GITHUB_APP_ID")
	installID := env("GITHUB_APP_INSTALLATION_ID")
	pem := env("GITHUB_APP_PEM_FILE")

	if token == "" && appID == "" {
		return authNone, errors.New("authentication not configured")
	}
	if token != "" && appID != "" {
		return authNone, errors.New("Both token and app auth configured")
	}
	if appID != "" && (installID == "" || pem == "") {
		return authNone, errors.New("App auth configured without all required parameters")
	}
	if token != "" {
		return authToken, nil
	}
	return authApp, nil
}

// validateToken validates a token by performing a GET /user request.
// A 401/403 response returns a typed "token invalid or expired" error.
// Classic PATs have X-OAuth-Scopes; fine-grained PATs do not.
// The token value is never placed in any returned error message.
func validateToken(ctx context.Context, baseURL, token string) (credInfo, error) {
	opts := []gogithub.ClientOptionsFunc{gogithub.WithAuthToken(token)}
	if baseURL != "" {
		opts = append(opts, gogithub.WithURLs(&baseURL, nil))
	}
	c, err := gogithub.NewClient(opts...)
	if err != nil {
		return credInfo{}, fmt.Errorf("creating github client: %w", err)
	}

	user, resp, err := c.Users.Get(ctx, "")
	if err != nil {
		if resp != nil && (resp.StatusCode == 401 || resp.StatusCode == 403) {
			return credInfo{}, errors.New("token invalid or expired")
		}
		return credInfo{}, fmt.Errorf("validating token: %w", err)
	}

	info := credInfo{Method: authToken, Login: user.GetLogin()}
	scopes := resp.Header.Get("X-OAuth-Scopes")
	if scopes != "" {
		// Classic PAT - parse comma-separated scopes.
		for _, s := range strings.Split(scopes, ",") {
			s = strings.TrimSpace(s)
			if s != "" {
				info.Scopes = append(info.Scopes, s)
			}
		}
	} else {
		// Fine-grained PAT - capabilities are unverified.
		info.FineGrained = true
	}
	return info, nil
}

// validateApp validates app auth by parsing provider-compatible PEM content
// and minting an installation token. PEM contents are never placed in returned
// errors.
func validateApp(ctx context.Context, baseURL, appID, installID, pemData string) (credInfo, error) {
	if appID == "" || installID == "" || pemData == "" {
		return credInfo{}, errors.New("App auth configured without all required parameters")
	}
	key, err := parseAppPrivateKey(pemData)
	if err != nil {
		return credInfo{}, errors.New("GitHub App PEM content is invalid")
	}
	token, err := mintInstallationToken(ctx, baseURL, appID, installID, key)
	if err != nil {
		return credInfo{}, fmt.Errorf("validating GitHub App auth: %w", err)
	}
	return credInfo{Method: authApp, FineGrained: true, Token: token}, nil
}

func parseAppPrivateKey(pemData string) (*rsa.PrivateKey, error) {
	normalized := strings.ReplaceAll(pemData, `\n`, "\n")
	block, _ := pem.Decode([]byte(normalized))
	if block == nil {
		return nil, errors.New("no PEM block")
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func mintInstallationToken(ctx context.Context, baseURL, appID, installID string, key *rsa.PrivateKey) (string, error) {
	jwt, err := signAppJWT(appID, key)
	if err != nil {
		return "", fmt.Errorf("signing app jwt: %w", err)
	}
	apiBase := "https://api.github.com/"
	if baseURL != "" {
		apiBase = baseURL
	}
	u, err := url.Parse(apiBase)
	if err != nil {
		return "", fmt.Errorf("parsing api base url: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u.JoinPath("app", "installations", installID, "access_tokens").String(), nil)
	if err != nil {
		return "", fmt.Errorf("creating app token request: %w", err)
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+jwt)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("requesting installation token: %w", err)
	}
	defer resp.Body.Close() //nolint:errcheck
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("reading installation token response: %w", err)
	}
	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("requesting installation token: status %d", resp.StatusCode)
	}
	var payload struct {
		Token string `json:"token"`
	}
	if err := json.Unmarshal(body, &payload); err != nil {
		return "", fmt.Errorf("decoding installation token response: %w", err)
	}
	if payload.Token == "" {
		return "", errors.New("installation token response missing token")
	}
	return payload.Token, nil
}

func signAppJWT(appID string, key *rsa.PrivateKey) (string, error) {
	now := time.Now()
	header := map[string]string{"alg": "RS256", "typ": "JWT"}
	claims := map[string]any{
		"iat": now.Add(-60 * time.Second).Unix(),
		"exp": now.Add(5 * time.Minute).Unix(),
		"iss": appID,
	}
	head, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	body, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}
	enc := base64.RawURLEncoding
	signingInput := enc.EncodeToString(head) + "." + enc.EncodeToString(body)
	digest := sha256.Sum256([]byte(signingInput))
	sig, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, digest[:])
	if err != nil {
		return "", err
	}
	return signingInput + "." + enc.EncodeToString(sig), nil
}
