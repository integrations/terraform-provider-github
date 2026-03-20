package github

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/go-github/v84/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type runnerGroupNetworking struct {
	NetworkConfigurationID *string `json:"network_configuration_id,omitempty"`
}

func getRunnerGroupNetworking(client *github.Client, ctx context.Context, path string) (*runnerGroupNetworking, *github.Response, error) {
	req, err := client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var runnerGroup runnerGroupNetworking
	resp, err := client.Do(ctx, req, &runnerGroup)
	if err != nil {
		var ghErr *github.ErrorResponse
		if errors.As(err, &ghErr) && ghErr.Response != nil && ghErr.Response.StatusCode == http.StatusNotModified {
			return nil, resp, nil
		}

		return nil, resp, err
	}

	return &runnerGroup, resp, nil
}

func updateRunnerGroupNetworking(client *github.Client, ctx context.Context, path string, networkConfigurationID *string) (*github.Response, error) {
	payload := map[string]any{
		"network_configuration_id": networkConfigurationID,
	}

	req, err := client.NewRequest("PATCH", path, payload)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func setRunnerGroupNetworkingState(d *schema.ResourceData, runnerGroup *runnerGroupNetworking) error {
	var networkConfigurationID any
	if runnerGroup != nil {
		networkConfigurationID = runnerGroup.NetworkConfigurationID
	}

	return d.Set("network_configuration_id", networkConfigurationID)
}
