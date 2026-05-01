package github

import (
	"testing"

	"github.com/google/go-github/v85/github"
)

func TestExpandHostedRunnerImage(t *testing.T) {
	t.Run("returns nil for empty list", func(t *testing.T) {
		result := expandHostedRunnerImage([]any{})
		if result != nil {
			t.Errorf("expected nil, got %v", result)
		}
	})

	t.Run("expands image with all fields", func(t *testing.T) {
		imageList := []any{
			map[string]any{
				"id":      "2306",
				"source":  "github",
				"version": "latest",
			},
		}
		result := expandHostedRunnerImage(imageList)
		if result == nil {
			t.Fatal("expected non-nil result")
		}
		if result.ID != "2306" {
			t.Errorf("expected ID '2306', got '%s'", result.ID)
		}
		if result.Source != "github" {
			t.Errorf("expected source 'github', got '%s'", result.Source)
		}
		if result.Version == nil || *result.Version != "latest" {
			t.Errorf("expected version 'latest', got %v", result.Version)
		}
	})

	t.Run("defaults version to latest when empty", func(t *testing.T) {
		imageList := []any{
			map[string]any{
				"id":      "2306",
				"source":  "github",
				"version": "",
			},
		}
		result := expandHostedRunnerImage(imageList)
		if result == nil {
			t.Fatal("expected non-nil result")
		}
		if result.Version == nil || *result.Version != "latest" {
			t.Errorf("expected version 'latest', got %v", result.Version)
		}
	})
}

func TestFlattenHostedRunnerImage(t *testing.T) {
	t.Run("returns empty slice for nil image", func(t *testing.T) {
		result := flattenHostedRunnerImage(nil)
		if len(result) != 0 {
			t.Errorf("expected empty slice, got %v", result)
		}
	})

	t.Run("flattens image details", func(t *testing.T) {
		id := "2306"
		source := "github"
		version := "latest"
		displayName := "Ubuntu 24.04"
		sizeGB := int64(86)
		image := &github.HostedRunnerImageDetail{
			ID:          &id,
			Source:      &source,
			Version:     &version,
			DisplayName: &displayName,
			SizeGB:      &sizeGB,
		}
		result := flattenHostedRunnerImage(image)
		if len(result) != 1 {
			t.Fatalf("expected 1 item, got %d", len(result))
		}
		m := result[0].(map[string]any)
		if m["id"] != "2306" {
			t.Errorf("expected id '2306', got %v", m["id"])
		}
		if m["source"] != "github" {
			t.Errorf("expected source 'github', got %v", m["source"])
		}
		if m["version"] != "latest" {
			t.Errorf("expected version 'latest', got %v", m["version"])
		}
		if m["display_name"] != "Ubuntu 24.04" {
			t.Errorf("expected display_name 'Ubuntu 24.04', got %v", m["display_name"])
		}
		if m["size_gb"] != 86 {
			t.Errorf("expected size_gb 86, got %v", m["size_gb"])
		}
	})
}

func TestFlattenHostedRunnerMachineSpec(t *testing.T) {
	t.Run("returns empty slice for nil spec", func(t *testing.T) {
		result := flattenHostedRunnerMachineSpec(nil)
		if len(result) != 0 {
			t.Errorf("expected empty slice, got %v", result)
		}
	})

	t.Run("flattens machine spec", func(t *testing.T) {
		spec := &github.HostedRunnerMachineSpec{
			ID:        "4-core",
			CPUCores:  4,
			MemoryGB:  16,
			StorageGB: 150,
		}
		result := flattenHostedRunnerMachineSpec(spec)
		if len(result) != 1 {
			t.Fatalf("expected 1 item, got %d", len(result))
		}
		m := result[0].(map[string]any)
		if m["id"] != "4-core" {
			t.Errorf("expected id '4-core', got %v", m["id"])
		}
		if m["cpu_cores"] != 4 {
			t.Errorf("expected cpu_cores 4, got %v", m["cpu_cores"])
		}
		if m["memory_gb"] != 16 {
			t.Errorf("expected memory_gb 16, got %v", m["memory_gb"])
		}
		if m["storage_gb"] != 150 {
			t.Errorf("expected storage_gb 150, got %v", m["storage_gb"])
		}
	})
}

func TestFlattenHostedRunnerPublicIPs(t *testing.T) {
	t.Run("returns empty slice for nil ips", func(t *testing.T) {
		result := flattenHostedRunnerPublicIPs(nil)
		if len(result) != 0 {
			t.Errorf("expected empty slice, got %v", result)
		}
	})

	t.Run("flattens public IPs", func(t *testing.T) {
		ips := []*github.HostedRunnerPublicIP{
			{Enabled: true, Prefix: "20.80.208.150", Length: 28},
			{Enabled: false, Prefix: "10.0.0.0", Length: 24},
		}
		result := flattenHostedRunnerPublicIPs(ips)
		if len(result) != 2 {
			t.Fatalf("expected 2 items, got %d", len(result))
		}
		first := result[0].(map[string]any)
		if first["enabled"] != true {
			t.Errorf("expected enabled true, got %v", first["enabled"])
		}
		if first["prefix"] != "20.80.208.150" {
			t.Errorf("expected prefix '20.80.208.150', got %v", first["prefix"])
		}
		if first["length"] != 28 {
			t.Errorf("expected length 28, got %v", first["length"])
		}
	})

	t.Run("skips nil entries", func(t *testing.T) {
		ips := []*github.HostedRunnerPublicIP{
			nil,
			{Enabled: true, Prefix: "20.80.208.150", Length: 28},
		}
		result := flattenHostedRunnerPublicIPs(ips)
		if len(result) != 1 {
			t.Fatalf("expected 1 item after skipping nil, got %d", len(result))
		}
	})
}
