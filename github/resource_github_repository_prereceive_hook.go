package github

import (
	"context"
	"fmt"
	"strconv"

	"github.com/google/go-github/v25/github"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceGithubRepositoryPreReceiveHook() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubRepositoryPreReceiveHookCreateUpdate,
		Read:   resourceGithubRepositoryPreReceiveHookRead,
		Update: resourceGithubRepositoryPreReceiveHookCreateUpdate,
		Delete: resourceGithubRepositoryPreReceiveHookDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"repository": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"enforcement": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"enabled", "disabled", "testing"}, false),
			},
			"config_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func fetchGitHubRepositoryPreReceiveHookByName(meta interface{}, repoName, hookName string) (*github.PreReceiveHook, error) {
	ctx := context.Background()
	client := meta.(*Organization).client
	orgName := meta.(*Organization).name

	opt := &github.ListOptions{
		PerPage: 100,
	}

	var hook *github.PreReceiveHook

	for {
		hooks, resp, err := client.Repositories.ListPreReceiveHooks(ctx, orgName, repoName, opt)
		if err != nil {
			return nil, err
		}

		for _, h := range hooks {
			n := *h.Name
			if n == hookName {
				hook = h
				break
			}
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	if *hook.ID <= 0 {
		return nil, fmt.Errorf("No pre-receive hook with name %s found on %s/%s", hookName, orgName, repoName)
	}

	return hook, nil
}

func resourceGithubRepositoryPreReceiveHookCreateUpdate(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	repoName := d.Get("repository").(string)
	hookName := d.Get("name").(string)

	hook, err := fetchGitHubRepositoryPreReceiveHookByName(meta, repoName, hookName)
	if err != nil {
		return err
	}

	enforcement := d.Get("enforcement").(string)
	hook.Enforcement = &enforcement

	if v, ok := d.GetOk("config_url"); ok {
		configURL := v.(string)
		hook.ConfigURL = &configURL
	}

	ctx := context.Background()
	client := meta.(*Organization).client
	orgName := meta.(*Organization).name
	_, _, err = client.Repositories.UpdatePreReceiveHook(ctx, orgName, repoName, *hook.ID, hook)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", orgName, repoName, strconv.FormatInt(*hook.ID, 10)))

	return resourceGithubRepositoryPreReceiveHookRead(d, meta)
}

func resourceGithubRepositoryPreReceiveHookRead(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	repoName := d.Get("repository").(string)
	hookName := d.Get("name").(string)

	hook, err := fetchGitHubRepositoryPreReceiveHookByName(meta, repoName, hookName)
	if err != nil {
		return err
	}

	d.Set("enforcement", hook.Enforcement)

	if _, ok := d.GetOk("config_url"); ok {
		d.Set("config_url", hook.ConfigURL)
	}

	return nil
}

func resourceGithubRepositoryPreReceiveHookDelete(d *schema.ResourceData, meta interface{}) error {
	err := checkOrganization(meta)
	if err != nil {
		return err
	}

	repoName := d.Get("repository").(string)
	hookName := d.Get("name").(string)

	hook, err := fetchGitHubRepositoryPreReceiveHookByName(meta, repoName, hookName)
	if err != nil {
		return err
	}

	disabled := "disabled"
	hook.Enforcement = &disabled

	ctx := context.Background()
	client := meta.(*Organization).client
	orgName := meta.(*Organization).name
	_, _, err = client.Repositories.UpdatePreReceiveHook(ctx, orgName, repoName, *hook.ID, hook)
	if _, _, err = client.Repositories.UpdatePreReceiveHook(ctx, orgName, repoName, *hook.ID, hook); err != nil {
		return err
	}

	return nil
}
