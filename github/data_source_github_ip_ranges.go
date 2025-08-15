package github

import (
	"fmt"
	"net"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGithubIpRanges() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGithubIpRangesRead,

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
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"dependabot": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"dependabot_ipv4": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"dependabot_ipv6": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceGithubIpRangesRead(d *schema.ResourceData, meta any) error {
	owner := meta.(*Owner)

	api, _, err := owner.v3client.Meta.Get(owner.StopContext)
	if err != nil {
		return err
	}

	cidrHooksIpv4, cidrHooksIpv6, err := splitIpv4Ipv6Cidrs(&api.Hooks)
	if err != nil {
		return err
	}

	cidrGitIpv4, cidrGitIpv6, err := splitIpv4Ipv6Cidrs(&api.Git)
	if err != nil {
		return err
	}

	cidrPackagesIpv4, cidrPackagesIpv6, err := splitIpv4Ipv6Cidrs(&api.Packages)
	if err != nil {
		return err
	}

	cidrPagesIpv4, cidrPagesIpv6, err := splitIpv4Ipv6Cidrs(&api.Pages)
	if err != nil {
		return err
	}

	cidrImporterIpv4, cidrImporterIpv6, err := splitIpv4Ipv6Cidrs(&api.Importer)
	if err != nil {
		return err
	}

	cidrActionsIpv4, cidrActionsIpv6, err := splitIpv4Ipv6Cidrs(&api.Actions)
	if err != nil {
		return err
	}

	cidrDependabotIpv4, cidrDependabotIpv6, err := splitIpv4Ipv6Cidrs(&api.Dependabot)
	if err != nil {
		return err
	}

	cidrWebIpv4, cidrWebIpv6, err := splitIpv4Ipv6Cidrs(&api.Web)
	if err != nil {
		return err
	}

	cidrApiIpv4, cidrApiIpv6, err := splitIpv4Ipv6Cidrs(&api.API)
	if err != nil {
		return err
	}

	if len(api.Hooks)+len(api.Git)+len(api.Pages)+len(api.Importer)+len(api.Actions)+len(api.Dependabot) > 0 {
		d.SetId("github-ip-ranges")
	}
	if len(api.Hooks) > 0 {
		err = d.Set("hooks", api.Hooks)
		if err != nil {
			return err
		}
		err = d.Set("hooks_ipv4", cidrHooksIpv4)
		if err != nil {
			return err
		}
		err = d.Set("hooks_ipv6", cidrHooksIpv6)
		if err != nil {
			return err
		}
	}
	if len(api.Git) > 0 {
		err = d.Set("git", api.Git)
		if err != nil {
			return err
		}
		err = d.Set("git_ipv4", cidrGitIpv4)
		if err != nil {
			return err
		}
		err = d.Set("git_ipv6", cidrGitIpv6)
		if err != nil {
			return err
		}
	}
	if len(api.Packages) > 0 {
		_ = d.Set("packages", api.Packages)
		_ = d.Set("packages_ipv4", cidrPackagesIpv4)
		_ = d.Set("packages_ipv6", cidrPackagesIpv6)
	}
	if len(api.Pages) > 0 {
		err = d.Set("pages", api.Pages)
		if err != nil {
			return err
		}
		err = d.Set("pages_ipv4", cidrPagesIpv4)
		if err != nil {
			return err
		}
		err = d.Set("pages_ipv6", cidrPagesIpv6)
		if err != nil {
			return err
		}
	}
	if len(api.Importer) > 0 {
		err = d.Set("importer", api.Importer)
		if err != nil {
			return err
		}
		err = d.Set("importer_ipv4", cidrImporterIpv4)
		if err != nil {
			return err
		}
		err = d.Set("importer_ipv6", cidrImporterIpv6)
		if err != nil {
			return err
		}
	}
	if len(api.Actions) > 0 {
		err = d.Set("actions", api.Actions)
		if err != nil {
			return err
		}
		err = d.Set("actions_ipv4", cidrActionsIpv4)
		if err != nil {
			return err
		}
		err = d.Set("actions_ipv6", cidrActionsIpv6)
		if err != nil {
			return err
		}
	}
	if len(api.Dependabot) > 0 {
		err = d.Set("dependabot", api.Dependabot)
		if err != nil {
			return err
		}
		err = d.Set("dependabot_ipv4", cidrDependabotIpv4)
		if err != nil {
			return err
		}
		err = d.Set("dependabot_ipv6", cidrDependabotIpv6)
		if err != nil {
			return err
		}
	}
	if len(api.Web) > 0 {
		err = d.Set("web", api.Web)
		if err != nil {
			return err
		}
		err = d.Set("web_ipv4", cidrWebIpv4)
		if err != nil {
			return err
		}
		err = d.Set("web_ipv6", cidrWebIpv6)
		if err != nil {
			return err
		}
	}
	if len(api.API) > 0 {
		err = d.Set("api", api.API)
		if err != nil {
			return err
		}
		err = d.Set("api_ipv4", cidrApiIpv4)
		if err != nil {
			return err
		}
		err = d.Set("api_ipv6", cidrApiIpv6)
		if err != nil {
			return err
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
			return nil, nil, fmt.Errorf("failed parsing cidr %s (%v)", cidr, err)
		}
		if cidrHost.To4() != nil {
			cidrIpv4 = append(cidrIpv4, cidr)
		} else {
			cidrIpv6 = append(cidrIpv6, cidr)
		}
	}

	return &cidrIpv4, &cidrIpv6, nil
}
