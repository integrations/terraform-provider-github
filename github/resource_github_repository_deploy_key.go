package github

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/go-github/v28/github"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGithubRepositoryDeployKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceGithubRepositoryDeployKeyCreate,
		Read:   resourceGithubRepositoryDeployKeyRead,
		Delete: resourceGithubRepositoryDeployKeyDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		// Deploy keys are defined immutable in the API. Updating results in force new.
		Schema: map[string]*schema.Schema{
			"key": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: suppressDeployKeyDiff,
			},
			"read_only": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Default:  true,
			},
			"repository": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"title": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"etag": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceGithubRepositoryDeployKeyCreate(d *schema.ResourceData, meta interface{}) error {
	orgName, err := getOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).client

	repoName := d.Get("repository").(string)
	key := d.Get("key").(string)
	title := d.Get("title").(string)
	readOnly := d.Get("read_only").(bool)
	ctx := prepareResourceContext(d)

	log.Printf("[DEBUG] Creating repository deploy key: %s (%s/%s)", title, orgName, repoName)
	resultKey, _, err := client.Repositories.CreateKey(ctx, orgName, repoName, &github.Key{
		Key:      github.String(key),
		Title:    github.String(title),
		ReadOnly: github.Bool(readOnly),
	})

	if err != nil {
		return err
	}

	d.SetId(buildTwoPartID(repoName, strconv.FormatInt(*resultKey.ID, 10)))

	return resourceGithubRepositoryDeployKeyRead(d, meta)
}

func resourceGithubRepositoryDeployKeyRead(d *schema.ResourceData, meta interface{}) error {
	orgName, err := getOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).client

	repoName, idString, err := parseTwoPartID(d.Id(), "repository", "ID")
	if err != nil {
		return err
	}

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return unconvertibleIdErr(idString, err)
	}

	ctx := prepareResourceContext(d)
	log.Printf("[DEBUG] Reading repository deploy key: %s (%s/%s)", d.Id(), orgName, repoName)
	key, resp, err := client.Repositories.GetKey(ctx, orgName, repoName, id)
	switch apires, apierr := apiResult(resp, err); apires {
	case APINotModified:
		return nil
	case APINotFound:
		log.Printf("[WARN] Removing repository deploy key %s from state because it no longer exists in GitHub", d.Id())
		d.SetId("")
		return nil
	case APIError:
		return apierr
	default:
		d.Set("etag", resp.Header.Get("ETag"))
		d.Set("key", key.Key)
		d.Set("read_only", key.ReadOnly)
		d.Set("repository", repoName)
		d.Set("title", key.Title)

		return nil
	}
}

func resourceGithubRepositoryDeployKeyDelete(d *schema.ResourceData, meta interface{}) error {
	orgName, err := getOrganization(meta)
	if err != nil {
		return err
	}

	client := meta.(*Organization).client

	repoName, idString, err := parseTwoPartID(d.Id(), "repository", "ID")
	if err != nil {
		return err
	}

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return unconvertibleIdErr(idString, err)
	}

	ctx := prepareResourceContext(d)

	log.Printf("[DEBUG] Deleting repository deploy key: %s (%s/%s)", idString, orgName, repoName)
	_, err = client.Repositories.DeleteKey(ctx, orgName, repoName, id)
	if err != nil {
		return err
	}

	return err
}

func suppressDeployKeyDiff(k, oldV, newV string, d *schema.ResourceData) bool {
	newV = strings.TrimSpace(newV)
	keyRe := regexp.MustCompile(`^([a-z0-9-]+ [^\s]+)( [^\s]+)?$`)
	newTrimmed := keyRe.ReplaceAllString(newV, "$1")

	return oldV == newTrimmed
}
