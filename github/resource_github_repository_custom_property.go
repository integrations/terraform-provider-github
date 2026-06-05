package github

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/go-github/v88/github"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubRepositoryCustomProperty() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubRepositoryCustomPropertyCreate,
		ReadContext:   resourceGithubRepositoryCustomPropertyRead,
		UpdateContext: resourceGithubRepositoryCustomPropertyUpdate,
		DeleteContext: resourceGithubRepositoryCustomPropertyDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGithubRepositoryCustomPropertyImport,
		},

		CustomizeDiff: customdiff.All(diffRepository, resourceGithubRepositoryCustomPropertyDiff),

		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceGithubRepositoryCustomPropertyV0().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceGithubRepositoryCustomPropertyStateUpgradeV0,
				Version: 0,
			},
		},

		Description: "Resource to manage GitHub repository custom properties.",

		Schema: map[string]*schema.Schema{
			"repository": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the repository.",
			},
			"repository_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "ID of the repository.",
			},
			"property_type": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "Type of the custom property. Valid values are `string`, `single_select`, `multi_select`, `true_false`, and `url`.",
				ForceNew:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{string(github.PropertyValueTypeString), string(github.PropertyValueTypeSingleSelect), string(github.PropertyValueTypeMultiSelect), string(github.PropertyValueTypeTrueFalse), string(github.PropertyValueTypeURL)}, false)),
			},
			"property_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the custom property.",
				ForceNew:    true,
			},
			"property_value": {
				Type:        schema.TypeSet,
				MinItems:    1,
				Required:    true,
				Description: "Value of the custom property. For `string`, `single_select`, `true_false`, and `url` property types, this should be a single value. For `multi_select` property types, this can be multiple values.",
				Elem: &schema.Schema{
					Type:             schema.TypeString,
					ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
				},
			},
		},
	}
}

func resourceGithubRepositoryCustomPropertyDiff(ctx context.Context, d *schema.ResourceDiff, _ any) error {
	tflog.Debug(ctx, "Diffing GitHub repository custom property")

	propertyTypeVal, _ := d.Get("property_type").(string)
	propertyType := github.PropertyValueType(propertyTypeVal)
	propertyValueVal, _ := d.Get("property_value").(*schema.Set)
	propertyValue := expandStringList(propertyValueVal.List())
	propertyCount := len(propertyValue)

	// The propertyValue can either be a list of strings or a string
	switch propertyType {
	case github.PropertyValueTypeMultiSelect:
		if propertyCount < 1 {
			return fmt.Errorf("custom property type %v requires at least one value", propertyType)
		}
	default:
		if propertyCount != 1 {
			return fmt.Errorf("custom property type %v requires exactly one value", propertyType)
		}
	}

	return nil
}

func resourceGithubRepositoryCustomPropertyCreate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, _ := d.Get("repository").(string)
	propertyName, _ := d.Get("property_name").(string)
	propertyTypeVal, _ := d.Get("property_type").(string)
	propertyType := github.PropertyValueType(propertyTypeVal)
	propertyValueVal, _ := d.Get("property_value").(*schema.Set)
	propertyValue := expandStringList(propertyValueVal.List())

	customProperty := github.CustomPropertyValue{
		PropertyName: propertyName,
	}

	// The propertyValue can either be a list of strings or a string
	switch propertyType {
	case github.PropertyValueTypeMultiSelect:
		customProperty.Value = propertyValue
	default:
		customProperty.Value = propertyValue[0]
	}

	_, err := client.Repositories.CreateOrUpdateCustomProperties(ctx, owner, repoName, []*github.CustomPropertyValue{&customProperty})
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := buildID(owner, repoName, propertyName)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return diag.FromErr(err)
	}
	repoID := int(repo.GetID())

	if err := d.Set("repository_id", repoID); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryCustomPropertyRead(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, _ := d.Get("repository").(string)
	propertyName, _ := d.Get("property_name").(string)

	props, _, err := client.Repositories.GetAllCustomPropertyValues(ctx, owner, repoName)
	if err != nil {
		if err, ok := errors.AsType[*github.ErrorResponse](err); ok && err.Response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	var property *github.CustomPropertyValue
	for _, prop := range props {
		if prop.PropertyName == propertyName {
			property = prop
			break
		}
	}

	if property == nil {
		d.SetId("")
		return nil
	}

	propertyValue, err := parseRepositoryCustomPropertyValueToStringSlice(property)
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := buildID(owner, repoName, propertyName)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	if err := d.Set("property_value", propertyValue); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryCustomPropertyUpdate(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, _ := d.Get("repository").(string)
	propertyName, _ := d.Get("property_name").(string)
	propertyTypeVal, _ := d.Get("property_type").(string)
	propertyType := github.PropertyValueType(propertyTypeVal)
	propertyValueVal, _ := d.Get("property_value").(*schema.Set)
	propertyValue := expandStringList(propertyValueVal.List())

	customProperty := github.CustomPropertyValue{
		PropertyName: propertyName,
	}

	// The propertyValue can either be a list of strings or a string
	switch propertyType {
	case github.PropertyValueTypeMultiSelect:
		customProperty.Value = propertyValue
	default:
		customProperty.Value = propertyValue[0]
	}

	_, err := client.Repositories.CreateOrUpdateCustomProperties(ctx, owner, repoName, []*github.CustomPropertyValue{&customProperty})
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := buildID(owner, repoName, propertyName)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	return nil
}

func resourceGithubRepositoryCustomPropertyDelete(ctx context.Context, d *schema.ResourceData, m any) diag.Diagnostics {
	meta, _ := m.(*Owner)
	client := meta.v3client
	owner := meta.name

	repoName, _ := d.Get("repository").(string)
	propertyName, _ := d.Get("property_name").(string)

	customProperty := github.CustomPropertyValue{
		PropertyName: propertyName,
		Value:        nil,
	}

	_, err := client.Repositories.CreateOrUpdateCustomProperties(ctx, owner, repoName, []*github.CustomPropertyValue{&customProperty})
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubRepositoryCustomPropertyImport(ctx context.Context, d *schema.ResourceData, m any) ([]*schema.ResourceData, error) {
	meta, _ := m.(*Owner)
	client := meta.v3client

	owner, repoName, propertyName, err := parseID3(d.Id())
	if err != nil {
		return nil, err
	}

	if !strings.EqualFold(owner, meta.name) {
		return nil, errors.New("owner in id must match authenticated owner")
	}

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return nil, err
	}
	repoID := int(repo.GetID())

	cp, _, err := client.Organizations.GetCustomProperty(ctx, owner, propertyName)
	if err != nil {
		return nil, err
	}

	if err := d.Set("repository", repoName); err != nil {
		return nil, err
	}
	if err := d.Set("repository_id", repoID); err != nil {
		return nil, err
	}
	if err := d.Set("property_type", cp.GetValueType()); err != nil {
		return nil, err
	}
	if err := d.Set("property_name", propertyName); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
