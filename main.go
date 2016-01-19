package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var foswikiHomeDir string

	flag.StringVar(&foswikiHomeDir, "dir", "", "Path to Foswiki's home directory")
	flag.Parse()

	if len(foswikiHomeDir) == 0 {
		fmt.Fprintln(os.Stderr, "ERROR: parameter dir missing")
		os.Exit(1)
	}

	if err := startProcessing(foswikiHomeDir); err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		os.Exit(1)
	}
}
