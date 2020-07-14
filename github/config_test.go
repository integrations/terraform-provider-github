package github

import (
	"testing"
)

func TestConfigClients(t *testing.T) {

	t.Run("returns a client for the v3 REST API", func(t *testing.T) {
		config := Config{
			Token:      "token",
			Individual: true,
		}

		meta, err := config.Clients()
		if err != nil {
			t.Fatalf("failed to return clients without error: %s", err.Error())
		}

		if client := meta.(*Organization).v3client; client == nil {
			t.Fatalf("failed to return a v3 client")
		}
	})

	t.Run("returns a client for the v4 GraphQL API", func(t *testing.T) {
		config := Config{
			Token:      "token",
			Individual: true,
		}

		meta, err := config.Clients()
		if err != nil {
			t.Fatalf("failed to return clients without error: %s", err.Error())
		}

		if client := meta.(*Organization).v4client; client == nil {
			t.Fatalf("failed to return a v4 client")
		}
	})

	t.Run("returns clients configured as anonymous", func(t *testing.T) {
		config := Config{
			Token:      "",
			Anonymous:  true,
			Individual: true,
		}

		meta, err := config.Clients()
		if err != nil {
			t.Fatalf("failed to return clients without error: %s", err.Error())
		}

		if client := meta.(*Organization).v4client; client == nil {
			t.Fatalf("failed to return a v4 client")
		}

		if client := meta.(*Organization).v3client; client == nil {
			t.Fatalf("failed to return a v3 client")
		}
	})

	t.Run("returns clients configured as individual", func(t *testing.T) {
		config := Config{
			Organization: "",
			Anonymous:    true,
			Individual:   true,
		}

		meta, err := config.Clients()
		if err != nil {
			t.Fatalf("failed to return clients without error: %s", err.Error())
		}

		if client := meta.(*Organization).v4client; client == nil {
			t.Fatalf("failed to return a v4 client")
		}

		if client := meta.(*Organization).v3client; client == nil {
			t.Fatalf("failed to return a v3 client")
		}
	})
}
