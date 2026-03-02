// Package tfproviderlint is a custom linter to be used by
// golangci-lint to integrate bflad/tfproviderlint.
package tfproviderlint

import (
	"github.com/integrations/terraform-provider-github/tools/tfproviderlint/checks"

	"github.com/bflad/tfproviderlint/passes"
	"github.com/bflad/tfproviderlint/xpasses"
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("tfproviderlint", New)
}

// TfProviderLintPlugin is a custom linter plugin for golangci-lint.
type TfProviderLintPlugin struct {
	enabledChecks map[string]bool
}

type Settings struct {
	EnabledChecks []string `json:"enabled-checks" yaml:"enabled-checks"`
}

// New returns an analysis.Analyzer to use with golangci-lint.
func New(settings any) (register.LinterPlugin, error) {
	enabledChecks := map[string]bool{}

	if settings != nil {
		if settingsMap, ok := settings.(map[string]any); ok {
			if enabledChecksRaw, ok := settingsMap["enabled-checks"]; ok {
				if enabledChecksList, ok := enabledChecksRaw.([]any); ok {
					for _, check := range enabledChecksList {
						if check, ok := check.(string); ok {
							enabledChecks[check] = true
						}
					}
				}
			}
		}
	}
	return &TfProviderLintPlugin{enabledChecks: enabledChecks}, nil
}

// BuildAnalyzers builds the analyzers for the TfProviderLintPlugin.
func (t *TfProviderLintPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	var analyzers []*analysis.Analyzer
	for _, check := range checks.AllChecks {
		if t.enabledChecks[check.Name] {
			analyzers = append(analyzers, check)
		}
	}
	for _, check := range passes.AllChecks {
		if t.enabledChecks[check.Name] {
			analyzers = append(analyzers, check)
		}
	}
	for _, check := range xpasses.AllChecks {
		if t.enabledChecks[check.Name] {
			analyzers = append(analyzers, check)
		}
	}

	return analyzers, nil
}

func (t *TfProviderLintPlugin) GetLoadMode() string {
	return register.LoadModeSyntax
}
