package github

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/go-github/v32/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGithubRepositoryEnvironment() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubRepositoryEnvironmentCreateOrUpdate,
		Read:   resourceGithubRepositoryEnvironmentRead,
		Update: resourceGithubRepositoryEnvironmentCreateOrUpdate,
		Delete: resourceGithubRepositoryEnvironmentDelete,

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"wait_timer": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"reviewers": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 5,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"id": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"protected_branches": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"custom_branch_policies": {
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceGithubRepositoryEnvironmentCreateOrUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	ctx := context.Background()

	repo := d.Get("repository").(string)
	environmentName := d.Get("name").(string)
	waitTimer := d.Get("wait_timer").(int)
	reviewers := d.Get("reviewers").([]interface{})

	environment := &createEnvironmentRequest{
		Name:      environmentName,
		WaitTimer: waitTimer,
	}
	for _, v := range reviewers {
		r, ok := v.(map[string]interface{})
		if !ok {
			return fmt.Errorf("failed to unpack environment reviewer block")
		}
		reviewerType := r["type"].(string)
		reviewerId := r["id"].(int)
		protectedBranches := r["protected_branches"].(bool)
		customBranchPolicies := r["custom_branch_policies"].(bool)
		p := requestReviewer{
			Type: reviewerType,
			Id:   reviewerId,
			BranchPolicy: DeploymentBranchPolicy{
				ProtectedBranches:    protectedBranches,
				CustomBranchPolicies: customBranchPolicies,
			},
		}
		environment.Reviewers = append(environment.Reviewers, p)
	}
	log.Printf("[DEBUG] Creating repo environment: %s/%s", repo, environmentName)
	_, err := createOrUpdateRepoEnvironment(ctx, owner, repo, environment, client)
	if err != nil {
		return err
	}
	d.SetId(buildTwoPartID(repo, environmentName))

	return resourceGithubRepositoryEnvironmentRead(d, meta)
}

type createEnvironmentRequest struct {
	Name      string            `json:"-"`
	WaitTimer int               `json:"wait_timer"`
	Reviewers []requestReviewer `json:"reviewers"`
}
type requestReviewer struct {
	Type         string                 `json:"type"`
	Id           int                    `json:"id"`
	BranchPolicy DeploymentBranchPolicy `json:"deployment_branch_policy"`
}

// createOrUpdateRepoEnvironment creates or updates a repository environment
// GitHub API docs: https://docs.github.com/en/rest/reference/repos#environments
func createOrUpdateRepoEnvironment(ctx context.Context, owner, repo string, environment *createEnvironmentRequest, c *github.Client) (*github.Response, error) {
	u := fmt.Sprintf("repos/%v/%v/environments/%v", owner, repo, environment.Name)
	req, err := c.NewRequest(http.MethodPut, u, environment)
	if err != nil {
		return nil, err
	}
	return c.Do(ctx, req, nil)
}

func resourceGithubRepositoryEnvironmentRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	ctx := context.Background()
	repo := d.Get("repository").(string)
	environmentName := d.Get("name").(string)
	log.Printf("[DEBUG] Reading repo environment: %s/%s", repo, environmentName)
	environment, _, err := readRepoEnvironment(ctx, owner, repo, environmentName, client)
	if err != nil {
		if ghErr, ok := err.(*github.ErrorResponse); ok {
			if ghErr.Response.StatusCode == http.StatusNotModified {
				return nil
			}
			if ghErr.Response.StatusCode == http.StatusNotFound {
				log.Printf("[WARN] Removing repository environment %s from state because it no longer exists in GitHub",
					d.Id())
				d.SetId("")
				return nil
			}
		}
		return err
	}
	d.Set("name", environment.Name)

	for _, protectionRule := range environment.ProtectionRules {
		switch protectionRule.Type {
		case "wait_tier":
			d.Set("wait_timer", protectionRule.WaitTimer)
		case "required_reviewers":
			d.Set("reviewers", []interface{}{protectionRule.Reviewers})
		}
	}
	return nil
}

// readRepoEnvironment read a repository environment
// GitHub API docs: https://docs.github.com/en/rest/reference/repos#environments
func readRepoEnvironment(ctx context.Context, owner, repo, environmentName string, c *github.Client) (*Environment, *github.Response, error) {
	u := fmt.Sprintf("repos/%v/%v/environments/%v", owner, repo, environmentName)
	req, err := c.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}
	environment := new(Environment)
	resp, err := c.Do(ctx, req, environment)
	if err != nil {
		return nil, nil, err
	}
	return environment, resp, err
}

// Environment represents an environment in a repository.
type Environment struct {
	ID                     int                    `json:"id"`
	NodeID                 string                 `json:"node_id"`
	Name                   string                 `json:"name"`
	URL                    string                 `json:"url"`
	HTMLURL                string                 `json:"html_url"`
	CreatedAt              time.Time              `json:"created_at"`
	UpdatedAt              time.Time              `json:"updated_at"`
	ProtectionRules        []ProtectionRule       `json:"protection_rules"`
	DeploymentBranchPolicy DeploymentBranchPolicy `json:"deployment_branch_policy"`
}

// ProtectionRule represents a list of protection rules for an environment.
type ProtectionRule struct {
	ID        int    `json:"id"`
	NodeID    string `json:"node_id"`
	Type      string `json:"type"`
	WaitTimer int    `json:"wait_timer,omitempty"`
	Reviewers []struct {
		Type     string `json:"type"`
		Reviewer struct {
			Id int `json:"id"`
		} `json:"reviewer"`
	} `json:"reviewers,omitempty"`
}

// DeploymentBranchPolicy represents deployment branch policy for an environment.
type DeploymentBranchPolicy struct {
	ProtectedBranches    bool `json:"protected_branches"`
	CustomBranchPolicies bool `json:"custom_branch_policies"`
}

func resourceGithubRepositoryEnvironmentDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*Owner).v3client
	owner := meta.(*Owner).name
	ctx := context.Background()

	repo := d.Get("repository").(string)
	environmentName := d.Get("name").(string)
	log.Printf("[DEBUG] Deleting repo environment: %s/%s", repo, environmentName)
	time.Sleep(1 * time.Minute)
	_, err := deleteRepoEnvironment(ctx, owner, repo, environmentName, client)
	if err != nil {
		return err
	}
	return nil
}

// deleteRepoEnvironment delete a repository environment
// GitHub API docs: https://docs.github.com/en/rest/reference/repos#environments
func deleteRepoEnvironment(ctx context.Context, owner, repo, environmentName string, c *github.Client) (*github.Response, error) {
	u := fmt.Sprintf("repos/%v/%v/environments/%v", owner, repo, environmentName)
	req, err := c.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(ctx, req, nil)
}
