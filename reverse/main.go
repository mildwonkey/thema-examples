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

	// translate to 1.0
	i10, _ := i00.Translate(thema.SV(1, 0))

	// back to 0.0 panics
	_ = translateTo(i10, thema.SV(0, 0))

	// so does 0.1
	_ = translateTo(i10, thema.SV(0, 1))

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

func translateTo(inst *thema.Instance, version thema.SyntacticVersion) *thema.Instance {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic: %v\n", err)
		}
	}()
	inst, _ = inst.Translate(version)
	return inst
}
