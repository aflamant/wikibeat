package main

import (
	"os"

	"github.com/aflamant/wikibeat/cmd"

	_ "github.com/aflamant/wikibeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
