package github

import (
	"context"
	"log"
	"net/http"

	"github.com/google/go-github/v84/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceGithubEnterpriseCustomProperties() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGithubEnterpriseCustomPropertiesCreate,
		ReadContext:   resourceGithubEnterpriseCustomPropertiesRead,
		UpdateContext: resourceGithubEnterpriseCustomPropertiesUpdate,
		DeleteContext: resourceGithubEnterpriseCustomPropertiesDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"enterprise_slug": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The slug of the enterprise.",
			},
			"property_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the custom property.",
			},
			"value_type": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The type of the value for the property. Can be one of: 'string', 'single_select', 'multi_select', 'true_false', 'url'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{string(github.PropertyValueTypeString), string(github.PropertyValueTypeSingleSelect), string(github.PropertyValueTypeMultiSelect), string(github.PropertyValueTypeTrueFalse), string(github.PropertyValueTypeURL)}, false)),
			},
			"required": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether the custom property is required.",
			},
			"default_values": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "The default value(s) of the custom property. For 'multi_select' properties, multiple values may be specified. For all other types, provide a single value.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A short description of the custom property.",
			},
			"allowed_values": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "An ordered list of allowed values for the property. Only applicable to 'single_select' and 'multi_select' types.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"values_editable_by": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Description:      "Who can edit the values of the property. Can be one of: 'org_actors', 'org_and_repo_actors'. Defaults to 'org_actors'.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"org_actors", "org_and_repo_actors"}, false)),
			},
		},
	}
}

func resourceGithubEnterpriseCustomPropertiesCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	enterpriseSlug := d.Get("enterprise_slug").(string)
	propertyName := d.Get("property_name").(string)

	property := buildEnterpriseCustomProperty(d)

	_, _, err := client.Enterprise.CreateOrUpdateCustomProperty(ctx, enterpriseSlug, propertyName, property)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(buildTwoPartID(enterpriseSlug, propertyName))
	return resourceGithubEnterpriseCustomPropertiesRead(ctx, d, meta)
}

func resourceGithubEnterpriseCustomPropertiesRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	enterpriseSlug, propertyName, err := parseTwoPartID(d.Id(), "enterprise_slug", "property_name")
	if err != nil {
		return diag.FromErr(err)
	}

	property, resp, err := client.Enterprise.GetCustomProperty(ctx, enterpriseSlug, propertyName)
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			log.Printf("[INFO] Removing enterprise custom property %s/%s from state because it no longer exists", enterpriseSlug, propertyName)
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	if err := d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("property_name", property.GetPropertyName()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("value_type", string(property.ValueType)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("required", property.GetRequired()); err != nil {
		return diag.FromErr(err)
	}

	var defaultValues []string
	if property.ValueType == github.PropertyValueTypeMultiSelect {
		if vals, ok := property.DefaultValueStrings(); ok {
			defaultValues = vals
		}
	} else {
		if val, ok := property.DefaultValueString(); ok {
			defaultValues = []string{val}
		}
	}
	if err := d.Set("default_values", defaultValues); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", property.GetDescription()); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("allowed_values", property.AllowedValues); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("values_editable_by", property.GetValuesEditableBy()); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubEnterpriseCustomPropertiesUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	enterpriseSlug, propertyName, err := parseTwoPartID(d.Id(), "enterprise_slug", "property_name")
	if err != nil {
		return diag.FromErr(err)
	}

	property := buildEnterpriseCustomProperty(d)

	_, _, err = client.Enterprise.CreateOrUpdateCustomProperty(ctx, enterpriseSlug, propertyName, property)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceGithubEnterpriseCustomPropertiesRead(ctx, d, meta)
}

func resourceGithubEnterpriseCustomPropertiesDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	client := meta.(*Owner).v3client

	enterpriseSlug, propertyName, err := parseTwoPartID(d.Id(), "enterprise_slug", "property_name")
	if err != nil {
		return diag.FromErr(err)
	}

	resp, err := client.Enterprise.RemoveCustomProperty(ctx, enterpriseSlug, propertyName)
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return nil
		}
		return diag.FromErr(err)
	}

	return nil
}

func resourceGithubEnterpriseCustomPropertiesImport(d *schema.ResourceData, meta any) ([]*schema.ResourceData, error) {
	enterpriseSlug, propertyName, err := parseTwoPartID(d.Id(), "enterprise_slug", "property_name")
	if err != nil {
		return nil, err
	}

	if err := d.Set("enterprise_slug", enterpriseSlug); err != nil {
		return nil, err
	}
	if err := d.Set("property_name", propertyName); err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}

func buildEnterpriseCustomProperty(d *schema.ResourceData) *github.CustomProperty {
	propertyName := d.Get("property_name").(string)
	valueType := github.PropertyValueType(d.Get("value_type").(string))
	required := d.Get("required").(bool)
	description := d.Get("description").(string)

	rawAllowedValues := d.Get("allowed_values").([]any)
	allowedValues := make([]string, 0, len(rawAllowedValues))
	for _, v := range rawAllowedValues {
		allowedValues = append(allowedValues, v.(string))
	}

	property := &github.CustomProperty{
		PropertyName:  &propertyName,
		ValueType:     valueType,
		Required:      &required,
		Description:   &description,
		AllowedValues: allowedValues,
	}

	rawDefaultValues := d.Get("default_values").([]any)
	defaultValues := make([]string, 0, len(rawDefaultValues))
	for _, v := range rawDefaultValues {
		defaultValues = append(defaultValues, v.(string))
	}
	if len(defaultValues) > 0 {
		if valueType == github.PropertyValueTypeMultiSelect {
			property.DefaultValue = defaultValues
		} else {
			property.DefaultValue = defaultValues[0]
		}
	}

	if val, ok := d.GetOk("values_editable_by"); ok {
		str := val.(string)
		property.ValuesEditableBy = &str
	}

	return property
}
