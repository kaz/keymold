package main

import (
	"fmt"
	"os"

	"github.com/kaz/keymold/cli"
)

func main() {
	if err := cli.Start(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(-1)
	}
}
