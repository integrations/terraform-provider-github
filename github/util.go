package github

import (
	"crypto/md5"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"slices"
	"sort"
	"strings"

	"github.com/google/go-github/v81/github"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	idSeparator        = ":"
	idSeperatorEscaped = `??`
)

// https://developer.github.com/guides/traversing-with-pagination/#basics-of-pagination
var maxPerPage = 100

// escapeIDPart escapes any idSeparator characters in a string.
func escapeIDPart(part string) string {
	return strings.ReplaceAll(part, idSeparator, idSeperatorEscaped)
}

// unescapeIDPart unescapes any escaped idSeparator characters in a string.
func unescapeIDPart(part string) string {
	return strings.ReplaceAll(part, idSeperatorEscaped, idSeparator)
}

// buildID joins the parts with the idSeparator.
func buildID(parts ...string) (string, error) {
	l := len(parts)
	if l == 0 {
		return "", fmt.Errorf("no parts provided to build id")
	}

	id := strings.Join(parts, idSeparator)

	if p := strings.Split(id, idSeparator); len(p) != l {
		return "", fmt.Errorf("unescaped seperators in id parts %v", parts)
	}

	return id, nil
}

// parseID splits the id by the idSeparator checking the count.
func parseID(id string, count int) ([]string, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("id is empty")
	}

	parts := strings.Split(id, idSeparator)
	if len(parts) != count {
		return nil, fmt.Errorf("unexpected ID format (%q); expected %d parts separated by %q", id, count, idSeparator)
	}

	return parts, nil
}

// parseID2 splits the id by the idSeparator into two parts.
func parseID2(id string) (string, string, error) {
	parts, err := parseID(id, 2)
	if err != nil {
		return "", "", err
	}

	return parts[0], parts[1], nil
}

// parseID3 splits the id by the idSeparator into three parts.
func parseID3(id string) (string, string, string, error) {
	parts, err := parseID(id, 3)
	if err != nil {
		return "", "", "", err
	}

	return parts[0], parts[1], parts[2], nil
}

func checkOrganization(meta any) error {
	if !meta.(*Owner).IsOrganization {
		return fmt.Errorf("this resource can only be used in the context of an organization, %q is a user", meta.(*Owner).name)
	}

	return nil
}

func caseInsensitive() schema.SchemaDiffSuppressFunc {
	return func(k, o, n string, d *schema.ResourceData) bool {
		return strings.EqualFold(o, n)
	}
}

// wrapErrors is provided to easily turn errors into diag.Diagnostics
// until we go through the provider and replace error usage.
func wrapErrors(errs []error) diag.Diagnostics {
	var diags diag.Diagnostics

	for _, err := range errs {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error",
			Detail:   err.Error(),
		})
	}

	return diags
}

// toDiagFunc is a helper that operates on Hashicorp's helper/validation functions
// and converts them to the diag.Diagnostic format
// --> nolint: oldFunc needs to be schema.SchemaValidateFunc to keep compatibility with
// the old code until all uses of schema.SchemaValidateFunc are gone.
func toDiagFunc(oldFunc schema.SchemaValidateFunc, keyName string) schema.SchemaValidateDiagFunc { //nolint:staticcheck
	return func(i any, path cty.Path) diag.Diagnostics {
		warnings, errors := oldFunc(i, keyName)
		var diags diag.Diagnostics

		for _, err := range errors {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  err.Error(),
			})
		}

		for _, warn := range warnings {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  warn,
			})
		}

		return diags
	}
}

func validateValueFunc(values []string) schema.SchemaValidateDiagFunc {
	return func(v any, k cty.Path) diag.Diagnostics {
		errs := make([]error, 0)
		value := v.(string)
		valid := slices.Contains(values, value)

		if !valid {
			errs = append(errs, fmt.Errorf("%s is an invalid value for argument %s", value, k))
		}
		return wrapErrors(errs)
	}
}

// return the pieces of id `left:right` as left, right.
func parseTwoPartID(id, left, right string) (string, string, error) {
	parts := strings.SplitN(id, ":", 2)
	if len(parts) != 2 {
		return "", "", fmt.Errorf("unexpected ID format (%q); expected %s:%s", id, left, right)
	}

	return parts[0], parts[1], nil
}

// format the strings into an id `a:b`.
func buildTwoPartID(a, b string) string {
	return fmt.Sprintf("%s:%s", a, b)
}

// return the pieces of id `left:center:right` as left, center, right.
func parseThreePartID(id, left, center, right string) (string, string, string, error) {
	parts := strings.SplitN(id, ":", 3)
	if len(parts) != 3 {
		return "", "", "", fmt.Errorf("unexpected ID format (%q). Expected %s:%s:%s", id, left, center, right)
	}

	return parts[0], parts[1], parts[2], nil
}

// format the strings into an id `a:b:c`.
func buildThreePartID(a, b, c string) string {
	return fmt.Sprintf("%s:%s:%s", a, b, c)
}

func buildChecksumID(v []string) string {
	sort.Strings(v)

	h := md5.New()
	// Hash.Write never returns an error. See https://pkg.go.dev/hash#Hash
	_, _ = h.Write([]byte(strings.Join(v, "")))
	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}

func expandStringList(configured []any) []string {
	vs := make([]string, 0, len(configured))
	for _, v := range configured {
		val, ok := v.(string)
		if ok && val != "" {
			vs = append(vs, val)
		}
	}
	return vs
}

func flattenStringList(v []string) []any {
	c := make([]any, 0, len(v))
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

func splitRepoFilePath(path string) (string, string) {
	parts := strings.Split(path, "/")
	return parts[0], strings.Join(parts[1:], "/")
}

// https://docs.github.com/en/actions/reference/encrypted-secrets#naming-your-secrets
var secretNameRegexp = regexp.MustCompile("^[a-zA-Z_][a-zA-Z0-9_]*$")

func validateSecretNameFunc(v any, path cty.Path) diag.Diagnostics {
	errs := make([]error, 0)
	name, ok := v.(string)
	if !ok {
		return wrapErrors([]error{fmt.Errorf("expected type of %s to be string", path)})
	}

	if !secretNameRegexp.MatchString(name) {
		errs = append(errs, errors.New("secret names can only contain alphanumeric characters or underscores and must not start with a number"))
	}

	if strings.HasPrefix(strings.ToUpper(name), "GITHUB_") {
		errs = append(errs, errors.New("secret names must not start with the GITHUB_ prefix"))
	}

	return wrapErrors(errs)
}

// deleteResourceOn404AndSwallow304OtherwiseReturnError will log and delete resource if error is 404 which indicates resource (or any of its ancestors)
// doesn't exist.
// resourceDescription represents a formatting string that represents the resource
// args will be passed to resourceDescription in `log.Printf`.
func deleteResourceOn404AndSwallow304OtherwiseReturnError(err error, d *schema.ResourceData, resourceDescription string, args ...any) error {
	var ghErr *github.ErrorResponse
	if errors.As(err, &ghErr) {
		if ghErr.Response.StatusCode == http.StatusNotModified {
			return nil
		}
		if ghErr.Response.StatusCode == http.StatusNotFound {
			log.Printf("[INFO] Removing "+resourceDescription+" from state because it no longer exists in GitHub",
				args...)
			d.SetId("")
			return nil
		}
	}
	return err
}
