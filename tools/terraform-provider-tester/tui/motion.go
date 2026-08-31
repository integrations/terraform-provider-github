package tui

import (
	"fmt"
	"time"
)

// stallThreshold is how long the run may go without a new test result before the
// status line shows a "quiet" hint. It is intentionally conservative: a single
// acceptance test can legitimately run for minutes without emitting an event, so
// a short threshold would nag during healthy runs.
const stallThreshold = 45 * time.Second

// formatRunDuration renders a duration in the GitHub CLI's compact style
// (e.g. "42s", "1m30s", "2m"), rounded to whole seconds. Durations under one
// second render as "0s"; callers gate on >= time.Second to avoid showing it.
func formatRunDuration(d time.Duration) string {
	d = d.Round(time.Second)
	secs := int(d.Seconds())
	if secs < 60 {
		return fmt.Sprintf("%ds", secs)
	}
	mins := secs / 60
	secs %= 60
	if secs == 0 {
		return fmt.Sprintf("%dm", mins)
	}
	return fmt.Sprintf("%dm%ds", mins, secs)
}
