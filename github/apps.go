package github

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"golang.org/x/oauth2/jws"
	"io/ioutil"
	"net/http"
	"time"
)

// GenerateOAuthTokenFromApp generates a GitHub OAuth access token from a set of valid GitHub App credentials. The
// returned token can be used to interact with both GitHub's REST and GraphQL APIs.
func GenerateOAuthTokenFromApp(baseURL, appID, appInstallationID, appPemFile string) (string, error) {
	pemData, err := ioutil.ReadFile(appPemFile)
	if err != nil {
		return "", err
	}

	jwt, err := generateAppJWT(appID, time.Now(), pemData)
	if err != nil {
		return "", err
	}

	token, err := getInstallationAccessToken(baseURL, jwt, appInstallationID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func getInstallationAccessToken(baseURL string, jwt string, installationID string) (string, error) {
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

	resBytes, err := ioutil.ReadAll(res.Body)
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
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	header := &jws.Header{
		Algorithm: "RS256", // Dictated by GitHub's API.
		Typ:       "JWT",   // Dictated by JWT's spec.
	}

	claims := &jws.ClaimSet{
		Iss: appID,
		// Using now - 60s to accommodate any client/server clock drift.
		Iat: now.Add(time.Duration(-60) * time.Second).Unix(),
		// The JWT's lifetime can be short as it is only used immediately
		// after to retrieve the installation's access  token.
		Exp: now.Add(time.Duration(5) * time.Minute).Unix(),
	}

	token, err := jws.Encode(header, claims, privateKey)
	if err != nil {
		return "", err
	}

	return token, nil
}
