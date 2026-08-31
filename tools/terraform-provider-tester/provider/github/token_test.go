package github

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"net/http"
	"net/http/httptest"
	"testing"
)

// mapEnv builds a env func from a key→value map; missing keys return "".
func mapEnv(m map[string]string) func(string) string {
	return func(key string) string { return m[key] }
}

// TestValidateTokenClassicScopes verifies that a classic PAT with X-OAuth-Scopes is parsed.
func TestValidateTokenClassicScopes(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-OAuth-Scopes", "repo, read:org, admin:org")
		json.NewEncoder(w).Encode(map[string]string{"login": "alice"})
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	info, err := validateToken(context.Background(), srv.URL+"/", "ghp_fake")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if info.Login != "alice" {
		t.Errorf("Login = %q, want %q", info.Login, "alice")
	}
	if info.FineGrained {
		t.Error("FineGrained should be false for classic PAT")
	}
	wantScopes := []string{"repo", "read:org", "admin:org"}
	if len(info.Scopes) != len(wantScopes) {
		t.Fatalf("Scopes = %v, want %v", info.Scopes, wantScopes)
	}
	for i, s := range wantScopes {
		if info.Scopes[i] != s {
			t.Errorf("Scopes[%d] = %q, want %q", i, info.Scopes[i], s)
		}
	}
}

// TestValidateTokenFineGrainedNoScopes verifies that absence of X-OAuth-Scopes marks FineGrained=true.
func TestValidateTokenFineGrainedNoScopes(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// No X-OAuth-Scopes header → fine-grained PAT.
		json.NewEncoder(w).Encode(map[string]string{"login": "bob"})
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	info, err := validateToken(context.Background(), srv.URL+"/", "ghp_fake")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !info.FineGrained {
		t.Error("FineGrained should be true when X-OAuth-Scopes is absent")
	}
	if len(info.Scopes) != 0 {
		t.Errorf("Scopes should be nil/empty for fine-grained PAT, got %v", info.Scopes)
	}
	if info.Login != "bob" {
		t.Errorf("Login = %q, want %q", info.Login, "bob")
	}
}

// TestValidateTokenInvalid verifies that a 401 response returns a typed error.
func TestValidateTokenInvalid(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	_, err := validateToken(context.Background(), srv.URL+"/", "ghp_fake")
	if err == nil {
		t.Fatal("expected error for 401 response, got nil")
	}
	if err.Error() != "token invalid or expired" {
		t.Errorf("error = %q, want %q", err.Error(), "token invalid or expired")
	}
}

// TestDetectAuthTokenXorApp tests the 4 branches from acc_test.go:167-180.
func TestDetectAuthTokenXorApp(t *testing.T) {
	appPEM := testAppPEM(t)
	tests := []struct {
		name       string
		env        map[string]string
		wantMeth   authMethod
		wantErrSub string // substring expected in error, "" means no error
	}{
		{
			name:       "both empty → not configured",
			env:        map[string]string{},
			wantMeth:   authNone,
			wantErrSub: "authentication not configured",
		},
		{
			name:       "both token and app set → ambiguous",
			env:        map[string]string{"GITHUB_TOKEN": "t", "GITHUB_APP_ID": "1"},
			wantMeth:   authNone,
			wantErrSub: "Both token and app auth configured",
		},
		{
			name:       "app without installation ID",
			env:        map[string]string{"GITHUB_APP_ID": "1", "GITHUB_APP_PEM_FILE": appPEM},
			wantMeth:   authNone,
			wantErrSub: "App auth configured without all required parameters",
		},
		{
			name:       "app without PEM file",
			env:        map[string]string{"GITHUB_APP_ID": "1", "GITHUB_APP_INSTALLATION_ID": "99"},
			wantMeth:   authNone,
			wantErrSub: "App auth configured without all required parameters",
		},
		{
			name:       "token only → authToken",
			env:        map[string]string{"GITHUB_TOKEN": "t"},
			wantMeth:   authToken,
			wantErrSub: "",
		},
		{
			name:       "app trio → authApp",
			env:        map[string]string{"GITHUB_APP_ID": "1", "GITHUB_APP_INSTALLATION_ID": "99", "GITHUB_APP_PEM_FILE": appPEM},
			wantMeth:   authApp,
			wantErrSub: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := detectAuth(mapEnv(tt.env))
			if tt.wantErrSub != "" {
				if err == nil {
					t.Fatalf("expected error containing %q, got nil", tt.wantErrSub)
				}
				if !contains(err.Error(), tt.wantErrSub) {
					t.Errorf("error = %q, want substring %q", err.Error(), tt.wantErrSub)
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
			}
			if m != tt.wantMeth {
				t.Errorf("authMethod = %d, want %d", m, tt.wantMeth)
			}
		})
	}
}

// TestValidateAppIncomplete verifies that missing installID or invalid PEM
// content returns an error.
func TestValidateAppIncomplete(t *testing.T) {
	appPEM := testAppPEM(t)
	t.Run("missing installID", func(t *testing.T) {
		_, err := validateApp(context.Background(), "", "42", "", appPEM)
		if err == nil {
			t.Fatal("expected error for missing installID")
		}
	})

	t.Run("missing pem content", func(t *testing.T) {
		_, err := validateApp(context.Background(), "", "42", "99", "")
		if err == nil {
			t.Fatal("expected error for missing PEM content")
		}
		want := "App auth configured without all required parameters"
		if err.Error() != want {
			t.Errorf("error = %q, want %q", err.Error(), want)
		}
	})

	t.Run("invalid PEM content", func(t *testing.T) {
		_, err := validateApp(context.Background(), "", "42", "99", "not PEM")
		if err == nil {
			t.Fatal("expected error for invalid PEM content")
		}
		if contains(err.Error(), "not PEM") {
			t.Fatalf("error leaked PEM content: %q", err.Error())
		}
	})

	t.Run("valid PEM content succeeds", func(t *testing.T) {
		mux := http.NewServeMux()
		mux.HandleFunc("POST /app/installations/99/access_tokens", func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("Authorization") == "" {
				t.Error("installation token request missing Authorization header")
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"token":"app_installation_token"}`)) //nolint:errcheck
		})
		srv := httptest.NewServer(mux)
		defer srv.Close()

		info, err := validateApp(context.Background(), srv.URL+"/", "42", "99", appPEM)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if info.Method != authApp {
			t.Errorf("Method = %v, want authApp", info.Method)
		}
		if !info.FineGrained {
			t.Error("FineGrained should be true for app auth")
		}
		if info.Token != "app_installation_token" {
			t.Errorf("Token = %q, want app_installation_token", info.Token)
		}
	})
}

// testAppPEM returns a throwaway PKCS1 RSA private key in PEM form. The key is
// generated fresh per call so the tests carry no committed secret and need no
// provider fixture; validateApp only needs a parseable key to sign the app JWT.
func testAppPEM(t testing.TB) string {
	t.Helper()
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("generating test RSA key: %v", err)
	}
	der := x509.MarshalPKCS1PrivateKey(key)
	return string(pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: der}))
}

func contains(s, sub string) bool {
	return len(sub) == 0 || (len(s) >= len(sub) && indexOf(s, sub) >= 0)
}

func indexOf(s, sub string) int {
	for i := range s {
		if i+len(sub) <= len(s) && s[i:i+len(sub)] == sub {
			return i
		}
	}
	return -1
}
