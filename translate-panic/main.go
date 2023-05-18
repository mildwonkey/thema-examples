package main

import (
	"embed"
	"fmt"
	"io/ioutil"
	"os"

	"cuelang.org/go/cue/cuecontext"
	"github.com/grafana/thema"
	"github.com/grafana/thema/vmux"
)

//go:embed example.cue cue.mod/module.cue
var LocalSchemaFS embed.FS

func main() {
	// bits and bobs to get started
	ctx := cuecontext.New()
	rt := thema.NewRuntime(ctx)
	exampleJSON, _ := ioutil.ReadFile("example.json")
	exdata, _ := vmux.NewJSONCodec("example.json").Decode(ctx, exampleJSON)

	lin, err := Lineage(rt)
	exitIf(err)
	sch00, err := lin.Schema(thema.SV(0, 0)) // we wouldn't normally hardcode this; just for the example
	exitIf(err)

	// The first "instance" of example.json, conforming to the schema at 0.0
	i00, err := sch00.Validate(exdata)
	exitIf(err)

	// Here's where we run into trouble: translate to a schema version that
	// doesn't exist There's no error to check for, so we get a panic after
	// calling Translate on a non-existent schema version.
	_, _ = i00.Translate(thema.SV(0, 1)) // note the translate does not return an error

}

func exitIf(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
