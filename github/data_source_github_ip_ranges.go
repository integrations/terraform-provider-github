package github

import (
	"context"
	"fmt"
	"net"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubIpRanges() *schema.Resource {
	return &schema.Resource{
		Description: "Get the GitHub IP ranges used by various GitHub services.",
		ReadContext: dataSourceGithubIpRangesRead,
		Schema: map[string]*schema.Schema{
			"hooks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"git": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"web": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"api": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"packages": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"pages": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"importer": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"actions": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "An array of IP addresses in CIDR format specifying the addresses that GitHub Actions will originate from.",
			},
			"actions_macos": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "An array of IP addresses in CIDR format specifying the addresses that GitHub Actions macOS runners will originate from.",
			},
			"github_enterprise_importer": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "An array of IP addresses in CIDR format specifying the addresses that GitHub Enterprise Importer will originate from.",
			},
			"dependabot": {
				Deprecated: "This attribute is no longer returned form the API, Dependabot now uses the GitHub Actions IP addresses.",
				Type:       schema.TypeList,
				Computed:   true,
				Elem:       &schema.Schema{Type: schema.TypeString},
			},
			"hooks_ipv4": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"git_ipv4": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"web_ipv4": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"api_ipv4": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"packages_ipv4": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"pages_ipv4": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"importer_ipv4": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"actions_ipv4": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "An array of IPv4 addresses in CIDR format specifying the addresses that GitHub Actions will originate from.",
			},
			"actions_macos_ipv4": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "An array of IPv4 addresses in CIDR format specifying the addresses that GitHub Actions macOS runners will originate from.",
			},
			"github_enterprise_importer_ipv4": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "An array of IPv4 addresses in CIDR format specifying the addresses that GitHub Enterprise Importer will originate from.",
			},
			"dependabot_ipv4": {
				Deprecated: "This attribute is no longer returned form the API, Dependabot now uses the GitHub Actions IP addresses.",
				Type:       schema.TypeList,
				Computed:   true,
				Elem:       &schema.Schema{Type: schema.TypeString},
			},
			"hooks_ipv6": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"git_ipv6": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"web_ipv6": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"api_ipv6": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"packages_ipv6": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"pages_ipv6": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"importer_ipv6": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"actions_ipv6": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "An array of IPv6 addresses in CIDR format specifying the addresses that GitHub Actions will originate from.",
			},
			"actions_macos_ipv6": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "An array of IPv6 addresses in CIDR format specifying the addresses that GitHub Actions macOS runners will originate from.",
			},
			"github_enterprise_importer_ipv6": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "An array of IPv6 addresses in CIDR format specifying the addresses that GitHub Enterprise Importer will originate from.",
			},
			"dependabot_ipv6": {
				Deprecated: "This attribute is no longer returned form the API, Dependabot now uses the GitHub Actions IP addresses.",
				Type:       schema.TypeList,
				Computed:   true,
				Elem:       &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceGithubIpRangesRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	owner := meta.(*Owner)

	api, _, err := owner.v3client.Meta.Get(owner.StopContext)
	if err != nil {
		return diag.FromErr(err)
	}

	cidrHooksIpv4, cidrHooksIpv6, err := splitIpv4Ipv6Cidrs(&api.Hooks)
	if err != nil {
		return diag.FromErr(err)
	}

	cidrGitIpv4, cidrGitIpv6, err := splitIpv4Ipv6Cidrs(&api.Git)
	if err != nil {
		return diag.FromErr(err)
	}

	cidrPackagesIpv4, cidrPackagesIpv6, err := splitIpv4Ipv6Cidrs(&api.Packages)
	if err != nil {
		return diag.FromErr(err)
	}

	cidrPagesIpv4, cidrPagesIpv6, err := splitIpv4Ipv6Cidrs(&api.Pages)
	if err != nil {
		return diag.FromErr(err)
	}

	cidrImporterIpv4, cidrImporterIpv6, err := splitIpv4Ipv6Cidrs(&api.Importer)
	if err != nil {
		return diag.FromErr(err)
	}

	cidrActionsIpv4, cidrActionsIpv6, err := splitIpv4Ipv6Cidrs(&api.Actions)
	if err != nil {
		return diag.FromErr(err)
	}

	cidrActionsMacosIpv4, cidrActionsMacosIpv6, err := splitIpv4Ipv6Cidrs(&api.ActionsMacos)
	if err != nil {
		return diag.FromErr(err)
	}

	cidrGithubEnterpriseImporterIpv4, cidrGithubEnterpriseImporterIpv6, err := splitIpv4Ipv6Cidrs(&api.GithubEnterpriseImporter)
	if err != nil {
		return diag.FromErr(err)
	}

	cidrDependabotIpv4, cidrDependabotIpv6, err := splitIpv4Ipv6Cidrs(&api.Dependabot)
	if err != nil {
		return diag.FromErr(err)
	}

	cidrWebIpv4, cidrWebIpv6, err := splitIpv4Ipv6Cidrs(&api.Web)
	if err != nil {
		return diag.FromErr(err)
	}

	cidrApiIpv4, cidrApiIpv6, err := splitIpv4Ipv6Cidrs(&api.API)
	if err != nil {
		return diag.FromErr(err)
	}

	if len(api.Hooks)+len(api.Git)+len(api.Pages)+len(api.Importer)+len(api.Actions)+len(api.Dependabot) > 0 {
		d.SetId("github-ip-ranges")
	}
	if len(api.Hooks) > 0 {
		if err := d.Set("hooks", api.Hooks); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("hooks_ipv4", cidrHooksIpv4); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("hooks_ipv6", cidrHooksIpv6); err != nil {
			return diag.FromErr(err)
		}
	}
	if len(api.Git) > 0 {
		if err := d.Set("git", api.Git); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("git_ipv4", cidrGitIpv4); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("git_ipv6", cidrGitIpv6); err != nil {
			return diag.FromErr(err)
		}
	}
	if len(api.Packages) > 0 {
		if err := d.Set("packages", api.Packages); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("packages_ipv4", cidrPackagesIpv4); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("packages_ipv6", cidrPackagesIpv6); err != nil {
			return diag.FromErr(err)
		}
	}
	if len(api.Pages) > 0 {
		if err := d.Set("pages", api.Pages); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("pages_ipv4", cidrPagesIpv4); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("pages_ipv6", cidrPagesIpv6); err != nil {
			return diag.FromErr(err)
		}
	}
	if len(api.Importer) > 0 {
		if err := d.Set("importer", api.Importer); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("importer_ipv4", cidrImporterIpv4); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("importer_ipv6", cidrImporterIpv6); err != nil {
			return diag.FromErr(err)
		}
	}
	if len(api.Actions) > 0 {
		if err := d.Set("actions", api.Actions); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("actions_ipv4", cidrActionsIpv4); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("actions_ipv6", cidrActionsIpv6); err != nil {
			return diag.FromErr(err)
		}
	}
	if len(api.ActionsMacos) > 0 {
		if err := d.Set("actions_macos", api.ActionsMacos); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("actions_macos_ipv4", cidrActionsMacosIpv4); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("actions_macos_ipv6", cidrActionsMacosIpv6); err != nil {
			return diag.FromErr(err)
		}
	}
	if len(api.GithubEnterpriseImporter) > 0 {
		if err := d.Set("github_enterprise_importer", api.GithubEnterpriseImporter); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("github_enterprise_importer_ipv4", cidrGithubEnterpriseImporterIpv4); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("github_enterprise_importer_ipv6", cidrGithubEnterpriseImporterIpv6); err != nil {
			return diag.FromErr(err)
		}
	}
	if len(api.Dependabot) > 0 {
		if err := d.Set("dependabot", api.Dependabot); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("dependabot_ipv4", cidrDependabotIpv4); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("dependabot_ipv6", cidrDependabotIpv6); err != nil {
			return diag.FromErr(err)
		}
	}
	if len(api.Web) > 0 {
		if err := d.Set("web", api.Web); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("web_ipv4", cidrWebIpv4); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("web_ipv6", cidrWebIpv6); err != nil {
			return diag.FromErr(err)
		}
	}
	if len(api.API) > 0 {
		if err := d.Set("api", api.API); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("api_ipv4", cidrApiIpv4); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("api_ipv6", cidrApiIpv6); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func splitIpv4Ipv6Cidrs(cidrs *[]string) (*[]string, *[]string, error) {
	cidrIpv4 := []string{}
	cidrIpv6 := []string{}

	for _, cidr := range *cidrs {
		cidrHost, _, err := net.ParseCIDR(cidr)
		if err != nil {
			return nil, nil, fmt.Errorf("failed parsing cidr %s (%w)", cidr, err)
		}
		if cidrHost.To4() != nil {
			cidrIpv4 = append(cidrIpv4, cidr)
		} else {
			cidrIpv6 = append(cidrIpv6, cidr)
		}
	}

	return &cidrIpv4, &cidrIpv6, nil
}
