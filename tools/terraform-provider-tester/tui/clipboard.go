package tui

import (
	"os"

	osc52 "github.com/aymanbagabas/go-osc52/v2"
)

// OscCopy copies s to the clipboard via OSC 52. Writing the escape sequence to
// stdout is SSH-safe and harmless on non-TTY stdout.
func OscCopy(s string) error {
	_, err := osc52.New(s).WriteTo(os.Stdout)
	return err
}
