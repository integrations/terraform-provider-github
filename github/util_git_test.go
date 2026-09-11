package github

import "testing"

func TestGitBlobSHA(t *testing.T) {
	t.Parallel()

	for content, want := range map[string]string{
		"":        "e69de29bb2d1d6434b8b29ae775ad8c2e48c5391",
		"hello\n": "ce013625030ba8dba906f756967f9e9ca394464a",
	} {
		if got := gitBlobSHA(content); got != want {
			t.Errorf("gitBlobSHA(%q) = %s, want %s", content, got, want)
		}
	}
}
