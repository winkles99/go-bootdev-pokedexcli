package repl

import "os"

func commandExit(cfg *replConfig, opts []string) error {
	os.Exit(0)
	return nil
}