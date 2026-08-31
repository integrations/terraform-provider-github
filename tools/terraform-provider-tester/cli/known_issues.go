package cli

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/github/terraform-provider-tester/engine"
	ghissues "github.com/github/terraform-provider-tester/provider/github"
)

func runKnownIssues(args []string, out, errOut io.Writer, d deps) int {
	if len(args) == 0 {
		fmt.Fprintln(errOut, "usage: known-issues list|sync|add [flags]")
		return 2
	}
	switch args[0] {
	case "list":
		return runKnownIssuesList(args[1:], out, errOut)
	case "sync":
		return runKnownIssuesSync(args[1:], out, errOut, d)
	case "add":
		return runKnownIssuesAdd(args[1:], out, errOut, d)
	default:
		fmt.Fprintf(errOut, "unknown known-issues subcommand %q\n", args[0])
		return 2
	}
}

func runKnownIssuesList(args []string, out, errOut io.Writer) int {
	fs := flag.NewFlagSet("known-issues list", flag.ContinueOnError)
	fs.SetOutput(errOut)
	var opts triageOptions
	format := formatText
	addKnownIssueListFlags(fs, &opts)
	addJSONFlag(fs, &format)
	if err := fs.Parse(args); err != nil {
		return 2
	}
	entries, err := knownIssueEntries(context.Background(), opts)
	if err != nil {
		fmt.Fprintln(errOut, "known-issues list:", err)
		return 1
	}
	if format == formatJSON {
		payload := struct {
			Version    int                      `json:"version"`
			IssuesRepo string                   `json:"issues_repo"`
			Entries    []engine.KnownIssueEntry `json:"entries"`
		}{Version: 1, IssuesRepo: opts.IssuesRepo, Entries: entries}
		enc := json.NewEncoder(out)
		enc.SetEscapeHTML(false)
		if err := enc.Encode(payload); err != nil {
			fmt.Fprintln(errOut, "writing JSON:", err)
			return 1
		}
		return 0
	}
	for _, entry := range entries {
		fmt.Fprintf(out, "%s issue #%d %s %s\n", entry.Fingerprint, entry.Issue, entry.State, strings.Join(entry.Tests, ","))
	}
	return 0
}

func runKnownIssuesSync(args []string, out, errOut io.Writer, d deps) int {
	fs := flag.NewFlagSet("known-issues sync", flag.ContinueOnError)
	fs.SetOutput(errOut)
	issuesRepo := defaultIssuesRepo
	var outPath string
	fs.StringVar(&issuesRepo, "issues-repo", defaultIssuesRepo, "GitHub issue repo owner/name")
	fs.StringVar(&outPath, "out", "", "output YAML file")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	if outPath == "" {
		fmt.Fprintln(errOut, "known-issues sync requires --out")
		return 2
	}
	svc := newKnownIssueSyncService(issuesRepo, outPath, d.newKnownIssueLister)
	result, err := svc.Sync(context.Background())
	if err != nil {
		fmt.Fprintln(errOut, "known-issues sync:", err)
		var cfgErr knownIssueListerConfigError
		if errors.As(err, &cfgErr) {
			return 2
		}
		return 1
	}
	fmt.Fprintf(out, "synced %d known issue(s) to %s\n", result.Count, result.CachePath)
	return 0
}

func runKnownIssuesAdd(args []string, out, errOut io.Writer, _ deps) int {
	fs := flag.NewFlagSet("known-issues add", flag.ContinueOnError)
	fs.SetOutput(errOut)
	var path, fingerprint, mode, testName, className string
	var issue int
	fs.StringVar(&path, "known-issues", "", "YAML known-issues file")
	fs.StringVar(&fingerprint, "fingerprint", "", "sha256 fingerprint")
	fs.IntVar(&issue, "issue", 0, "GitHub issue number")
	fs.StringVar(&mode, "mode", "known-real", "known issue mode")
	fs.StringVar(&testName, "test", "", "test name")
	fs.StringVar(&className, "class", "", "failure class")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	if path == "" || fingerprint == "" {
		fmt.Fprintln(errOut, "known-issues add requires --known-issues and --fingerprint")
		return 2
	}
	file := engine.KnownIssueFile{Version: 1, Label: "acctest-failure"}
	if existing, err := engine.LoadKnownIssuesFile(path, time.Now().UTC()); err == nil {
		file = *existing
	} else if !os.IsNotExist(err) {
		fmt.Fprintln(errOut, "known-issues add:", err)
		return 1
	}
	entry := engine.KnownIssueEntry{Fingerprint: fingerprint, Issue: issue, State: "open", Mode: mode, Origin: "local"}
	if strings.HasPrefix(fingerprint, "sha256:") && len(fingerprint) >= len("sha256:")+16 {
		entry.Short = strings.TrimPrefix(fingerprint, "sha256:")[:16]
	}
	if testName != "" {
		entry.Tests = []string{testName}
	}
	if className != "" {
		entry.Classes = []string{className}
	}
	replaced := false
	for i := range file.Entries {
		if file.Entries[i].Fingerprint == fingerprint {
			file.Entries[i] = entry
			replaced = true
			break
		}
	}
	if !replaced {
		file.Entries = append(file.Entries, entry)
	}
	if err := engine.SaveKnownIssuesFile(path, file); err != nil {
		fmt.Fprintln(errOut, "known-issues add:", err)
		return 1
	}
	fmt.Fprintf(out, "added %s to %s\n", fingerprint, path)
	return 0
}

func addKnownIssueListFlags(fs *flag.FlagSet, opts *triageOptions) {
	opts.IssuesRepo = defaultIssuesRepo
	fs.StringVar(&opts.KnownIssuesPath, "known-issues", "", "optional YAML known-issues registry")
	fs.BoolVar(&opts.KnownIssuesOffline, "known-issues-offline", false, "use only the local known-issues file")
	fs.StringVar(&opts.IssuesRepo, "issues-repo", defaultIssuesRepo, "GitHub issue repo owner/name")
}

func knownIssueEntries(ctx context.Context, opts triageOptions) ([]engine.KnownIssueEntry, error) {
	reg := &engine.KnownIssueRegistry{CachePath: opts.KnownIssuesPath, Offline: opts.KnownIssuesOffline}
	if !opts.KnownIssuesOffline {
		client, err := ghissues.NewIssueClientFromEnv(opts.IssuesRepo)
		if err != nil {
			return nil, err
		}
		reg.Live = client
	}
	return reg.Entries(ctx)
}
