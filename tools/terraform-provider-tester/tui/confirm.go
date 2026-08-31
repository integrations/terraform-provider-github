package tui

func phraseMatches(input, expected string) bool { return input == expected }
func sweepPhrase(owner string) string           { return "SWEEP " + owner }
func fileIssuePhrase(short string) string       { return "FILE " + short }
