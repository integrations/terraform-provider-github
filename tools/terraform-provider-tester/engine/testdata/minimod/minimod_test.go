package minimod_test

import (
	"testing"
	"time"
)

// TestAlwaysPasses is a fixture test that always passes.
func TestAlwaysPasses(t *testing.T) {}

// TestAlwaysFails is a fixture test that always fails.
func TestAlwaysFails(t *testing.T) {
	t.Fatal("intentional failure for harness fixture test")
}

// TestLongRunning sleeps long enough to be killed by context cancellation tests.
func TestLongRunning(t *testing.T) {
	time.Sleep(60 * time.Second)
}
