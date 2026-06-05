package main

import (
	"fmt"
	"os"

	"github.com/chrischangcode/auto-pr-fixer-demo/pkg/token"
	"github.com/chrischangcode/auto-pr-fixer-demo/pkg/version"
)

func main() {
	v := version.Get()
	fmt.Printf("auto-pr-fixer-demo %s (go: %s, built: %s, platform: %s)\n",
		v.Version, v.GoVersion, v.BuildTime, v.Platform)

	if len(os.Args) > 1 && os.Args[1] == "token" {
		tok, err := token.FetchToken()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(tok)
	}
}
