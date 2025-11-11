package github

import (
	"context"
	"fmt"
	"strconv"

	"github.com/google/go-github/v66/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubRepositoryPreReceiveHook() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubRepositoryPreReceiveHookCreate,
		Read:   resourceGithubRepositoryPreReceiveHookRead,
		Update: resourceGithubRepositoryPreReceiveHookUpdate,
		Delete: resourceGithubRepositoryPreReceiveHookDelete,
		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The repository of the pre-receive hook.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the pre-receive hook.",
			},
			"enforcement": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"enabled", "disabled", "testing"}, false),
				Description:  "The state of enforcement for the hook on the repository. Possible values for enforcement are 'enabled', 'disabled' and 'testing'. 'disabled' indicates the pre-receive hook will not run. 'enabled' indicates it will run and reject any pushes that result in a non-zero status. 'testing' means the script will run but will not cause any pushes to be rejected.",
			},
			"configuration_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for the endpoint where enforcement is set.",
			},
		},
	}
}

func resourceGithubRepositoryPreReceiveHookCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client

	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	hookName := d.Get("name").(string)

	hook, err := fetchGitHubRepositoryPreReceiveHookByName(meta, repoName, hookName)
	if err != nil {
		return err
	}

	enforcement := d.Get("enforcement").(string)
	hook.Enforcement = &enforcement

	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	_, _, err = client.Repositories.UpdatePreReceiveHook(ctx, owner, repoName, hook.GetID(), hook)
	if err != nil {
		return err
	}
	d.SetId(strconv.FormatInt(hook.GetID(), 10))

	return resourceGithubRepositoryPreReceiveHookRead(d, meta)
}

func resourceGithubRepositoryPreReceiveHookRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)

	hookID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}

	hook, _, err := client.Repositories.GetPreReceiveHook(ctx, owner, repoName, hookID)
	if err != nil {
		return err
	}
	if err = d.Set("enforcement", hook.Enforcement); err != nil {
		return err
	}
	if err = d.Set("configuration_url", hook.ConfigURL); err != nil {
		return err
	}

	return nil
}

func resourceGithubRepositoryPreReceiveHookUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)
	enforcement := d.Get("enforcement").(string)

	hookID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}

	hook := &github.PreReceiveHook{
		Enforcement: &enforcement,
	}
	_, _, err = client.Repositories.UpdatePreReceiveHook(ctx, owner, repoName, hookID, hook)
	if err != nil {
		return err
	}

	return resourceGithubRepositoryPreReceiveHookRead(d, meta)
}

func resourceGithubRepositoryPreReceiveHookDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	ctx := context.WithValue(context.Background(), ctxId, d.Id())

	owner := meta.(*Owner).name
	repoName := d.Get("repository").(string)

	hookID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return unconvertibleIdErr(d.Id(), err)
	}
	_, err = client.Repositories.DeletePreReceiveHook(ctx, owner, repoName, hookID)
	if err != nil {
		return err
	}

	return nil
}

func fetchGitHubRepositoryPreReceiveHookByName(meta interface{}, repoName, hookName string) (*github.PreReceiveHook, error) {
	ctx := context.Background()
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name

	opt := &github.ListOptions{
		PerPage: 100,
	}

	var hook *github.PreReceiveHook

	for {
		hooks, resp, err := client.Repositories.ListPreReceiveHooks(ctx, owner, repoName, opt)
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
		return nil, fmt.Errorf("no pre-receive hook with name %s found on %s/%s", hookName, owner, repoName)
	}

	return hook, nil
}
