package github

import (
	"context"
	"errors"
	"fmt"
	"maps"
	"net/http"
	"strconv"

	"github.com/google/go-github/v93/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubActionsRepositoryPolicy() *schema.Resource {
	s := map[string]*schema.Schema{
		"repository": {
			Type:             schema.TypeString,
			Required:         true,
			ForceNew:         true,
			Description:      "The name of the repository.",
			ValidateDiagFunc: validation.ToDiagFunc(validation.StringLenBetween(1, 100)),
		},
	}
	maps.Copy(s, actionsPolicyCommonSchema())

	return &schema.Resource{
		Description:   "Manages a GitHub Actions policy (workflow execution protection) for a repository. Actions policies restrict which events and actors can trigger workflows.\n\nYou must have admin access to the repository to use this resource.",
		CreateContext: resourceGithubActionsRepositoryPolicyCreate,
		ReadContext:   resourceGithubActionsRepositoryPolicyRead,
		UpdateContext: resourceGithubActionsRepositoryPolicyUpdate,
		DeleteContext: resourceGithubActionsRepositoryPolicyDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubActionsRepositoryPolicyImport,
		},
		Schema: s,
	}
}

func resourceGithubActionsRepositoryPolicyCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	repoName := d.Get("repository").(string)
	ctx = tflog.SetField(ctx, "repository", repoName)

	req := github.CreateActionsPolicyRequest{
		Name:        d.Get("name").(string),
		Enforcement: github.RulesetEnforcement(d.Get("enforcement").(string)),
		Rules:       expandActionsPolicyRules(d),
	}
	if wp := expandActionsPolicyWorkflowPath(d, false); wp != nil {
		req.Conditions = &github.ActionsPolicyConditions{WorkflowPath: wp}
	}

	policy, _, err := meta.v3client.Actions.CreateRepositoryPolicy(ctx, meta.name, repoName, req)
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := buildID(escapeIDPart(repoName), strconv.FormatInt(policy.ID, 10))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	return resourceGithubActionsRepositoryPolicyRead(ctx, d, m)
}

func resourceGithubActionsRepositoryPolicyRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	repoName, policyID, err := parseActionsRepositoryPolicyID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	ctx = tflog.SetField(ctx, "repository", repoName)

	policy, _, err := meta.v3client.Actions.GetRepositoryPolicy(ctx, meta.name, repoName, policyID)
	if err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			tflog.Info(ctx, "Removing Actions policy from state because it no longer exists in GitHub", map[string]any{"policy_id": policyID})
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	if err := d.Set("repository", repoName); err != nil {
		return diag.FromErr(err)
	}
	if err := flattenActionsPolicyCommon(d, policy); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func resourceGithubActionsRepositoryPolicyUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	repoName, policyID, err := parseActionsRepositoryPolicyID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	ctx = tflog.SetField(ctx, "repository", repoName)

	req := github.UpdateActionsPolicyRequest{
		Name:        new(d.Get("name").(string)),
		Enforcement: new(github.RulesetEnforcement(d.Get("enforcement").(string))),
		Conditions:  &github.ActionsPolicyConditions{WorkflowPath: expandActionsPolicyWorkflowPath(d, true)},
		Rules:       expandActionsPolicyRules(d),
	}

	if _, _, err := meta.v3client.Actions.UpdateRepositoryPolicy(ctx, meta.name, repoName, policyID, req); err != nil {
		return diag.FromErr(err)
	}

	return resourceGithubActionsRepositoryPolicyRead(ctx, d, m)
}

func resourceGithubActionsRepositoryPolicyDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta := m.(*Owner)
	repoName, policyID, err := parseActionsRepositoryPolicyID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	ctx = tflog.SetField(ctx, "repository", repoName)

	if _, err := meta.v3client.Actions.DeleteRepositoryPolicy(ctx, meta.name, repoName, policyID); err != nil {
		if ghErr, ok := errors.AsType[*github.ErrorResponse](err); ok && ghErr.Response.StatusCode == http.StatusNotFound {
			return nil
		}
		return diag.FromErr(err)
	}
	return nil
}

func resourceGithubActionsRepositoryPolicyImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	if _, _, err := parseActionsRepositoryPolicyID(d.Id()); err != nil {
		return nil, fmt.Errorf("invalid import ID %q, expected <repository>:<policy_id>: %w", d.Id(), err)
	}
	return []*schema.ResourceData{d}, nil
}

func parseActionsRepositoryPolicyID(id string) (string, int64, error) {
	repoPart, policyPart, err := parseID2(id)
	if err != nil {
		return "", 0, err
	}
	policyID, err := strconv.ParseInt(policyPart, 10, 64)
	if err != nil {
		return "", 0, unconvertibleIdErr(policyPart, err)
	}
	return unescapeIDPart(repoPart), policyID, nil
}
