package main

import (
	"embed"
	"fmt"
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
	exampleJSON, _ := os.ReadFile("example.json")
	exdata, _ := vmux.NewJSONCodec("example.json").Decode(ctx, exampleJSON)
	lin, _ := Lineage(rt)
	sch00, _ := lin.Schema(thema.SV(0, 0)) // we wouldn't normally hardcode this; just for the example

	// The first "instance" of example.json, conforming to the schema at 0.0
	i00, _ := sch00.Validate(exdata)

	// now translate to 1.0, where we changed the datasource string field to a struct.
	// This is a case where Translate arguably should return an error, but it doesn't.
	i10, _ := i00.Translate(thema.SV(1, 0))
	if i10 == nil {
		fmt.Println("i00 translated to nil")
	}

	// both of these panic: #Translate.out.steps.hi: index out of range [2] with length 2
	// translate 0.0 to 2.0
	i20, _ := i00.Translate(thema.SV(2, 0))
	if i20 == nil {
		fmt.Println("i00 translated to 2.0 returned nil")
	}

	// translate 1.0 to 2.0
	i20, _ = i10.Translate(thema.SV(2, 0))
	if i20 == nil {
		fmt.Println("i00 translated to 2.0 returned nil")
	}
}

// don't you judge me earl
func print(thing interface{}) {
	if thing != nil {
		switch thing := thing.(type) {
		case error:
			fmt.Println(thing)
		default:
			fmt.Printf("%#v\n", thing)
		}
	}
}

func exitIf(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
