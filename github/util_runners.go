package github

import (
	"github.com/google/go-github/v82/github"
)

func expandHostedRunnerImage(imageList []any) *github.HostedRunnerImage {
	if len(imageList) == 0 {
		return nil
	}

	imageMap := imageList[0].(map[string]any)
	image := &github.HostedRunnerImage{}

	if id, ok := imageMap["id"].(string); ok {
		image.ID = id
	}
	if source, ok := imageMap["source"].(string); ok {
		image.Source = source
	}
	if version, ok := imageMap["version"].(string); ok && version != "" {
		image.Version = version
	} else {
		// Default to 'latest' for GitHub-owned images as required by the API
		image.Version = "latest"
	}

	return image
}

func flattenHostedRunnerImage(image *github.HostedRunnerImageDetail) []any {
	if image == nil {
		return []any{}
	}

	result := make(map[string]any)

	if id := image.GetID(); id != "" {
		result["id"] = id
	}
	if source := image.GetSource(); source != "" {
		result["source"] = source
	}
	if version := image.GetVersion(); version != "" {
		result["version"] = version
	}
	if sizeGB := image.GetSizeGB(); sizeGB != 0 {
		result["size_gb"] = int(sizeGB)
	}
	if displayName := image.GetDisplayName(); displayName != "" {
		result["display_name"] = displayName
	}

	return []any{result}
}

func flattenHostedRunnerMachineSpec(spec *github.HostedRunnerMachineSpec) []any {
	if spec == nil {
		return []any{}
	}

	result := make(map[string]any)
	result["id"] = spec.ID
	result["cpu_cores"] = spec.CPUCores
	result["memory_gb"] = spec.MemoryGB
	result["storage_gb"] = spec.StorageGB

	return []any{result}
}

func flattenHostedRunnerPublicIPs(ips []*github.HostedRunnerPublicIP) []any {
	if ips == nil {
		return []any{}
	}

	result := make([]any, 0, len(ips))
	for _, ip := range ips {
		if ip == nil {
			continue
		}

		ipResult := make(map[string]any)
		ipResult["enabled"] = ip.Enabled
		ipResult["prefix"] = ip.Prefix
		ipResult["length"] = ip.Length
		result = append(result, ipResult)
	}

	return result
}
