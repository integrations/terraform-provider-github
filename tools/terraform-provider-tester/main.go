package main

import (
	"os"

	"github.com/github/terraform-provider-tester/cli"
)

func main() {
	os.Exit(cli.Run(os.Args[1:], os.Stdout, os.Stderr))
}
