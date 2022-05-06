package main

import (
	"fmt"
	"github.com/go-the-way/anorm-cli/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	}
}
