package github

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/google/go-github/v92/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/shurcooL/githubv4"

	"github.com/integrations/terraform-provider-github/v6/internal/ghclient"
)

type enterpriseAppConfig struct {
	slug, installerID, installerInstallationID, installerPEM string
	automationID, automationClientID, automationPEM          string
}

type enterpriseAppLifecycle struct {
	config                *enterpriseAppConfig
	enterpriseID          string
	creationJWT           *github.Client
	installerJWT          *github.Client
	installer             *github.Client
	automationJWT         *github.Client
	automationPermissions map[string]string
	opts                  ghclient.ClientOptions
	orgClients            func(int64) (*github.Client, *githubv4.Client, error)
}

func enterpriseAppAuthSchema() *schema.Schema {
	fields := map[string]*schema.Schema{}
	for name, description := range map[string]string{
		"enterprise_slug":           "Enterprise slug to manage. Opts into App-only enterprise organization lifecycle; deletion and updates are blocked.",
		"installer_app_id":          "Installer App numeric ID or client ID. Defaults to INSTALLER_APP_CLIENT_ID.",
		"installer_installation_id": "Installer enterprise installation ID. Defaults to INSTALLER_APP_INSTALLATION_ID.",
		"installer_pem_file":        "Installer private key PEM content, not a filename. Defaults to INSTALLER_APP_PRIVATE_KEY.",
		"automation_app_id":         "Automation App numeric ID. Defaults to ORG_APP_ID.",
		"automation_client_id":      "Automation App client ID. Defaults to ORG_APP_CLIENT_ID.",
		"automation_pem_file":       "Automation private key PEM content, not a filename. Defaults to ORG_APP_PRIVATE_KEY.",
	} {
		fields[name] = &schema.Schema{
			Type: schema.TypeString, Optional: true, Description: description,
			ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotWhiteSpace),
		}
	}
	fields["enterprise_slug"].Optional = false
	fields["enterprise_slug"].Required = true
	fields["enterprise_slug"].ValidateDiagFunc = validation.ToDiagFunc(validation.StringMatch(regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9-]*$`), "must be an enterprise slug"))
	for _, name := range []string{"automation_app_id", "installer_installation_id"} {
		fields[name].ValidateDiagFunc = validation.ToDiagFunc(validation.StringMatch(regexp.MustCompile(`^[1-9][0-9]*$`), "must be a positive numeric ID"))
	}
	fields["installer_pem_file"].Sensitive = true
	fields["automation_pem_file"].Sensitive = true
	return &schema.Schema{Type: schema.TypeList, Optional: true, MaxItems: 1, Description: "Opt-in enterprise App lifecycle credentials. All credentials must be supplied here or through their dedicated environment variables.", Elem: &schema.Resource{Schema: fields}}
}

func parseEnterpriseAppAuth(data *schema.ResourceData) (*enterpriseAppConfig, error) {
	value, ok := data.GetOk("enterprise_app_auth")
	if !ok {
		return nil, nil
	}
	blocks := value.([]any)
	if len(blocks) != 1 || blocks[0] == nil {
		return nil, fmt.Errorf("enterprise_app_auth requires a complete block")
	}
	fields := blocks[0].(map[string]any)
	get := func(name, env string) string {
		if value, _ := fields[name].(string); value != "" {
			return value
		}
		return os.Getenv(env)
	}
	config := &enterpriseAppConfig{
		slug: get("enterprise_slug", ""), installerID: get("installer_app_id", "INSTALLER_APP_CLIENT_ID"),
		installerInstallationID: get("installer_installation_id", "INSTALLER_APP_INSTALLATION_ID"),
		installerPEM:            get("installer_pem_file", "INSTALLER_APP_PRIVATE_KEY"),
		automationID:            get("automation_app_id", "ORG_APP_ID"), automationClientID: get("automation_client_id", "ORG_APP_CLIENT_ID"),
		automationPEM: get("automation_pem_file", "ORG_APP_PRIVATE_KEY"),
	}
	if config.slug == "" || config.installerID == "" || config.installerInstallationID == "" || config.installerPEM == "" || config.automationID == "" || config.automationClientID == "" || config.automationPEM == "" {
		return nil, fmt.Errorf("enterprise_app_auth requires enterprise_slug and all installer and automation credentials (attributes or dedicated environment variables)")
	}
	config.installerPEM = strings.ReplaceAll(config.installerPEM, `\n`, "\n")
	config.automationPEM = strings.ReplaceAll(config.automationPEM, `\n`, "\n")
	return config, nil
}
