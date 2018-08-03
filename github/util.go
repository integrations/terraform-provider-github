package github

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
)

const (
	// https://developer.github.com/guides/traversing-with-pagination/#basics-of-pagination
	maxPerPage = 100
)

func toGithubID(id string) int64 {
	githubID, _ := strconv.ParseInt(id, 10, 64)
	return githubID
}

func fromGithubID(id *int64) string {
	return strconv.FormatInt(*id, 10)
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

// return the pieces of id `a:b` as a, b
func parseTwoPartID(id string) (string, string, error) {
	parts := strings.SplitN(id, ":", 2)
	if len(parts) != 2 {
		return "", "", fmt.Errorf("Unexpected ID format (%q). Expected organization:name", id)
	}

	return parts[0], parts[1], nil
}

// format the strings into an id `a:b`
func buildTwoPartID(a, b *string) string {
	return fmt.Sprintf("%s:%s", *a, *b)
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
