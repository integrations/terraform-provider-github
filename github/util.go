package github

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/google/go-github/v31/github"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	// https://developer.github.com/guides/traversing-with-pagination/#basics-of-pagination
	maxPerPage = 100
)

func checkOrganization(meta interface{}) error {
	if !meta.(*Owner).IsOrganization {
		return fmt.Errorf("This resource can only be used in the context of an organization, %q is a user.", meta.(*Owner).name)
	}

	return nil
}

func caseInsensitive() schema.SchemaDiffSuppressFunc {
	return func(k, old, new string, d *schema.ResourceData) bool {
		return strings.EqualFold(old, new)
	}
}

func validateValueFunc(values []string) schema.SchemaValidateFunc {
	return func(v interface{}, k string) (we []string, errors []error) {
		value := v.(string)
		valid := false
		for _, role := range values {
			if value == role {
				valid = true
				break
			}
		}

		if !valid {
			errors = append(errors, fmt.Errorf("%s is an invalid value for argument %s", value, k))
		}
		return
	}
}

// return the pieces of id `left:right` as left, right
func parseTwoPartID(id, left, right string) (string, string, error) {
	parts := strings.SplitN(id, ":", 2)
	if len(parts) != 2 {
		return "", "", fmt.Errorf("Unexpected ID format (%q). Expected %s:%s", id, left, right)
	}

	return parts[0], parts[1], nil
}

// format the strings into an id `a:b`
func buildTwoPartID(a, b string) string {
	return fmt.Sprintf("%s:%s", a, b)
}

func expandStringList(configured []interface{}) []string {
	vs := make([]string, 0, len(configured))
	for _, v := range configured {
		val, ok := v.(string)
		if ok && val != "" {
			vs = append(vs, val)
		}
	}
	return vs
}

func flattenStringList(v []string) []interface{} {
	c := make([]interface{}, 0, len(v))
	for _, s := range v {
		c = append(c, s)
	}
	return c
}

func unconvertibleIdErr(id string, err error) *unconvertibleIdError {
	return &unconvertibleIdError{OriginalId: id, OriginalError: err}
}

type unconvertibleIdError struct {
	OriginalId    string
	OriginalError error
}

func (e *unconvertibleIdError) Error() string {
	return fmt.Sprintf("Unexpected ID format (%q), expected numerical ID. %s",
		e.OriginalId, e.OriginalError.Error())
}

func validateTeamIDFunc(v interface{}, keyName string) (we []string, errors []error) {
	teamIDString, ok := v.(string)
	if !ok {
		return nil, []error{fmt.Errorf("expected type of %s to be string", keyName)}
	}
	// Check that the team ID can be converted to an int
	if _, err := strconv.ParseInt(teamIDString, 10, 64); err != nil {
		return nil, []error{unconvertibleIdErr(teamIDString, err)}
	}

	return
}

func splitRepoFilePath(path string) (string, string) {
	parts := strings.Split(path, "/")
	return parts[0], strings.Join(parts[1:], "/")
}

// getFileCommit will return the commit in which the file specified was last updated.
func getFileCommit(client *github.Client, org, repo, file, branch string) (*github.RepositoryCommit, error) {
	ctx := context.WithValue(context.Background(), ctxId, fmt.Sprintf("%s/%s", repo, file))
	opts := &github.CommitsListOptions{
		SHA: branch,
	}
	allCommits := []*github.RepositoryCommit{}
	for {
		commits, resp, err := client.Repositories.ListCommits(ctx, org, repo, opts)
		if err != nil {
			return nil, err
		}

		allCommits = append(allCommits, commits...)

		if resp.NextPage == 0 {
			break
		}

		opts.Page = resp.NextPage
	}

	for _, c := range allCommits {
		sha := c.GetSHA()

		// Skip merge commits
		if strings.Contains(c.Commit.GetMessage(), "Merge branch") {
			continue
		}

		rc, _, err := client.Repositories.GetCommit(ctx, org, repo, sha)
		if err != nil {
			return nil, err
		}

		for _, f := range rc.Files {
			if f.GetFilename() == file && f.GetStatus() != "removed" {
				log.Printf("[DEBUG] Found file: %s in commit: %s", file, sha)
				return rc, nil
			}
		}
	}

	return nil, fmt.Errorf("Cannot find file %s in repo %s/%s", file, org, repo)
}
