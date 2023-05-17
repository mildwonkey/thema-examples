package main

import (
	"embed"
	"fmt"
	"io/ioutil"
	"os"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"github.com/grafana/thema"
	"github.com/grafana/thema/load"
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

	// A secondary issue (which will have it's own exmaple elsewhere): I'm
	// having trouble getting the generated Lineage() function working (with or
	// without a lineage path in the cue file), so this is the hacky manual
	// version.
	bi, err := load.InstanceWithThema(LocalSchemaFS, "")
	exitIf(err)
	v := ctx.BuildInstance(bi)
	// get the lineage and a schema, and validate the example.json data
	lineage := v.LookupPath(cue.ParsePath("lineage"))
	lin, err := thema.BindLineage(lineage, rt)
	exitIf(err)
	sch00, err := lin.Schema(thema.SV(0, 0)) // we wouldn't normally hardcode this; just for the example
	exitIf(err)

	// The first "instance" of example.json, conforming to the schema at 0.0
	i00, err := sch00.Validate(exdata)
	exitIf(err)
	origTitleStr, err := i00.Underlying().LookupPath(cue.ParsePath("title")).String()
	exitIf(err)
	fmt.Printf("original title string: %q\n", origTitleStr) // "foo"

	// Here's where we run into trouble: translate to 0.1.
	i01, _ := i00.Translate(thema.SV(0, 1)) // note the translate does not return an error
	// There's no change to the title, so I expect that the output is the same as the above/
	_, err = i01.Underlying().LookupPath(cue.ParsePath("title")).String()
	// but instead ...
	print(err) // #Translate.out.result.result: field not found: title

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
