package github

import (
	"fmt"
	"net"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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

func dataSourceGithubIpRangesRead(d *schema.ResourceData, meta interface{}) error {
	owner := meta.(*Owner)

	api, _, err := owner.v3client.APIMeta(owner.StopContext)
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

	if len(api.Hooks)+len(api.Git)+len(api.Pages)+len(api.Importer)+len(api.Actions)+len(api.Dependabot) > 0 {
		d.SetId("github-ip-ranges")
	}
	if len(api.Hooks) > 0 {
		d.Set("hooks", api.Hooks)
		d.Set("hooks_ipv4", cidrHooksIpv4)
		d.Set("hooks_ipv6", cidrHooksIpv6)
	}
	if len(api.Git) > 0 {
		d.Set("git", api.Git)
		d.Set("git_ipv4", cidrGitIpv4)
		d.Set("git_ipv6", cidrGitIpv6)
	}
	if len(api.Pages) > 0 {
		d.Set("pages", api.Pages)
		d.Set("pages_ipv4", cidrPagesIpv4)
		d.Set("pages_ipv6", cidrPagesIpv6)
	}
	if len(api.Importer) > 0 {
		d.Set("importer", api.Importer)
		d.Set("importer_ipv4", cidrImporterIpv4)
		d.Set("importer_ipv6", cidrImporterIpv6)
	}
	if len(api.Actions) > 0 {
		d.Set("actions", api.Actions)
		d.Set("actions_ipv4", cidrActionsIpv4)
		d.Set("actions_ipv6", cidrActionsIpv6)
	}
	if len(api.Dependabot) > 0 {
		d.Set("dependabot", api.Dependabot)
		d.Set("dependabot_ipv4", cidrDependabotIpv4)
		d.Set("dependabot_ipv6", cidrDependabotIpv6)
	}

	return nil
}

func splitIpv4Ipv6Cidrs(cidrs *[]string) (*[]string, *[]string, error) {
	cidrIpv4 := []string{}
	cidrIpv6 := []string{}

	for _, cidr := range *cidrs {
		cidrHost, _, err := net.ParseCIDR(cidr)
		if err != nil {
			return nil, nil, fmt.Errorf("Failed parsing cidr %s (%v)", cidr, err)
		}
		if cidrHost.To4() != nil {
			cidrIpv4 = append(cidrIpv4, cidr)
		} else {
			cidrIpv6 = append(cidrIpv6, cidr)
		}
	}

	return &cidrIpv4, &cidrIpv6, nil
}
