package github

import (
	"errors"
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
func parseTwoPartID(id string) (string, string) {
	parts := strings.SplitN(id, ":", 2)
	return parts[0], parts[1]
}

// validateTwoPartID performs a quick validation of a two-part ID, designed for
// use when validation has not been previously possible, such as importing.
func validateTwoPartID(id string) error {
	if id == "" {
		return errors.New("no ID supplied. Please supply an ID format matching organization:username")
	}
	parts := strings.Split(id, ":")
	if len(parts) != 2 {
		return fmt.Errorf("incorrectly formatted ID %q. Please supply an ID format matching organization:username", id)
	}
	return nil
}

// format the strings into an id `a:b`
func buildTwoPartID(a, b *string) string {
	return fmt.Sprintf("%s:%s", *a, *b)
}
