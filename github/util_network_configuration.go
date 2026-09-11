package github

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/google/go-github/v89/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var networkConfigurationNamePattern = regexp.MustCompile(`^[a-zA-Z0-9._-]+$`)

func setNetworkConfigurationState(d *schema.ResourceData, configuration *github.NetworkConfiguration) error {
	if err := d.Set("name", configuration.GetName()); err != nil {
		return err
	}
	if configuration.ComputeService != nil {
		if err := d.Set("compute_service", string(*configuration.ComputeService)); err != nil {
			return err
		}
	}
	if err := d.Set("network_settings_ids", configuration.NetworkSettingsIDs); err != nil {
		return err
	}
	createdOn := ""
	if configuration.CreatedOn != nil {
		createdOn = configuration.CreatedOn.Format(time.RFC3339)
	}
	if err := d.Set("created_on", createdOn); err != nil {
		return err
	}

	return nil
}

func networkSettingsScopeError(err error, scope string) error {
	if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusUnprocessableEntity {
		return fmt.Errorf("%w. verify the network settings ID belongs to the same %s: Azure GitHub.Network/networkSettings resources are registered against a single organization or enterprise and cannot be shared across scopes", err, scope)
	}

	return err
}
