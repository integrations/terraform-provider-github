package github

import (
	"context"
	"encoding/json"
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

	for _, r := range environment.ProtectionRules {
		switch r.Type {
		case "wait_timer":
			d.Set("wait_timer", r.Rule.(*EnvironmentWaitTimer).Time)
		case "required_reviewers":
			d.Set("reviewers", []interface{}{r.Rule.(*EnvironmentRequiredReviewers).Reviewers})
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
	ID                     int                         `json:"id"`
	NodeID                 string                      `json:"node_id"`
	Name                   string                      `json:"name"`
	URL                    string                      `json:"url"`
	HTMLURL                string                      `json:"html_url"`
	CreatedAt              time.Time                   `json:"created_at"`
	UpdatedAt              time.Time                   `json:"updated_at"`
	ProtectionRules        []EnvironmentProtectionRule `json:"protection_rules"`
	DeploymentBranchPolicy DeploymentBranchPolicy      `json:"deployment_branch_policy"`
}

// EnvironmentProtectionRule represents an environment protection rule
type EnvironmentProtectionRule struct {
	EnvironmentProtectionRuleMeta
	Rule interface{}
}

// EnvironmentProtectionRuleMeta represents a protection rule metadata
type EnvironmentProtectionRuleMeta struct {
	ID     int    `json:"id"`
	Type   string `json:"type"`
	NodeID string `json:"node_id"`
}

func (p *EnvironmentProtectionRule) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &p.EnvironmentProtectionRuleMeta); err != nil {
		return err
	}
	switch p.Type {
	case "wait_timer":
		p.Rule = new(EnvironmentWaitTimer)
		return json.Unmarshal(data, p.Rule)
	case "required_reviewers":
		p.Rule = new(EnvironmentRequiredReviewers)
		return json.Unmarshal(data, p.Rule)
		//	case branch_policy not needed rule meta fulfill the required fields
	}
	return nil
}

//EnvironmentWaitTimer represents wait timer of an environment in minutes.
type EnvironmentWaitTimer struct {
	Time int `json:"wait_timer"`
}

//EnvironmentRequiredReviewers represents required reviewers.
type EnvironmentRequiredReviewers struct {
	Reviewers []EnvironmentRequiredReviewer `json:"reviewers,omitempty"`
}

//EnvironmentRequiredReviewer is a reviewer of environment can be either a user or a team.
type EnvironmentRequiredReviewer struct {
	EnvironmentRequiredReviewerMeta
	Reviewer struct {
		Entity interface{} `json:"reviewer"`
	}
}

//EnvironmentRequiredReviewerMeta represents meta data for a reviewer.
type EnvironmentRequiredReviewerMeta struct {
	Type string `json:"type"`
}

func (r *EnvironmentRequiredReviewer) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &r.EnvironmentRequiredReviewerMeta); err != nil {
		return err
	}
	switch r.Type {
	case "User":
		r.Reviewer.Entity = new(github.User)
		return json.Unmarshal(data, &r.Reviewer)
	case "Team":
		r.Reviewer.Entity = new(github.Team)
		return json.Unmarshal(data, &r.Reviewer)
	}
	return nil
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
