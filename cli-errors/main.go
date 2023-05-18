package main

import (
	"embed"
	"fmt"
	"os"

	"cuelang.org/go/cue/cuecontext"
	"github.com/grafana/thema"
)

//go:embed example.cue cue.mod/module.cue
var LocalSchemaFS embed.FS

func main() {
	// bits and bobs to get started
	ctx := cuecontext.New()
	rt := thema.NewRuntime(ctx)

	_, err := Lineage(rt)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
