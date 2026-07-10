package github

import (
	"fmt"
	"net/http"
	"regexp"
	"testing"
	"testing/synctest"
	"time"

	"github.com/google/go-github/v89/github"
)

func Test_retryUntilOK(t *testing.T) {
	t.Parallel()

	errTest := fmt.Errorf("test error")

	for _, tt := range []struct {
		name    string
		f       func() func() (int, bool, error)
		opts    *retryOptions
		want    int
		wantErr *string
	}{
		{
			name: "returns_value_immediately",
			f: func() func() (int, bool, error) {
				return func() (int, bool, error) {
					return 42, true, nil
				}
			},
			want: 42,
		},
		{
			name: "returns_error_immediately",
			f: func() func() (int, bool, error) {
				return func() (int, bool, error) {
					return 0, false, errTest
				}
			},
			wantErr: new(errTest.Error()),
		},
		{
			name: "retries_until_value_found",
			f: func() func() (int, bool, error) {
				staticCounter := 0
				return func() (int, bool, error) {
					staticCounter++
					if staticCounter < 3 {
						return 0, false, nil
					}
					return 42, true, nil
				}
			},
			want: 42,
		},
		{
			name: "retries_until_error",
			f: func() func() (int, bool, error) {
				staticCounter := 0
				return func() (int, bool, error) {
					staticCounter++
					if staticCounter < 3 {
						return 0, false, nil
					}
					return 0, false, errTest
				}
			},
			wantErr: new(errTest.Error()),
		},
		{
			name: "retries_until_timeout",
			f: func() func() (int, bool, error) {
				return func() (int, bool, error) {
					return 0, false, nil
				}
			},
			wantErr: new("timeout while waiting for state to become 'found'"),
		},
		{
			name: "retries_until_value_found_with_custom_options",
			f: func() func() (int, bool, error) {
				staticCounter := 0
				return func() (int, bool, error) {
					staticCounter++
					if staticCounter < 3 {
						return 0, false, nil
					}
					return 42, true, nil
				}
			},
			opts: &retryOptions{
				delay:      500 * time.Millisecond,
				maxRetries: 5,
				timeout:    5 * time.Second,
			},
			want: 42,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				got, err := retryUntilOK(t.Context(), tt.f(), tt.opts)

				if tt.wantErr != nil {
					if err == nil {
						t.Fatalf("expected error %q, got nil", *tt.wantErr)
					}
					if !regexp.MustCompile(regexp.QuoteMeta(*tt.wantErr)).MatchString(err.Error()) {
						t.Fatalf("expected error %q, got %q", *tt.wantErr, err.Error())
					}
					return
				}

				if err != nil {
					t.Fatalf("expected no error, got %v", err)
				}

				if got != tt.want {
					t.Fatalf("expected %v, got %v", tt.want, got)
				}
			})
		})
	}
}

func Test_retryUntilResourceFound(t *testing.T) {
	t.Parallel()

	errTest := fmt.Errorf("test error")

	for _, tt := range []struct {
		name    string
		f       func() func() (int, error)
		want    int
		wantErr *string
	}{
		{
			name: "returns_value_immediately",
			f: func() func() (int, error) {
				return func() (int, error) {
					return 42, nil
				}
			},
			want: 42,
		},
		{
			name: "returns_error_immediately",
			f: func() func() (int, error) {
				return func() (int, error) {
					return 0, errTest
				}
			},
			wantErr: new(errTest.Error()),
		},
		{
			name: "retries_until_value_found",
			f: func() func() (int, error) {
				staticCounter := 0
				return func() (int, error) {
					staticCounter++
					if staticCounter < 3 {
						return 0, &github.ErrorResponse{Response: &http.Response{StatusCode: 404}}
					}
					return 42, nil
				}
			},
			want: 42,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				got, err := retryUntilResourceFound(t.Context(), tt.f(), nil)

				if tt.wantErr != nil {
					if err == nil {
						t.Fatalf("expected error %q, got nil", *tt.wantErr)
					}
					if !regexp.MustCompile(regexp.QuoteMeta(*tt.wantErr)).MatchString(err.Error()) {
						t.Fatalf("expected error %q, got %q", *tt.wantErr, err.Error())
					}
					return
				}

				if err != nil {
					t.Fatalf("expected no error, got %v", err)
				}

				if got != tt.want {
					t.Fatalf("expected %v, got %v", tt.want, got)
				}
			})
		})
	}
}
