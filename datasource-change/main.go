package main

import (
	"embed"
	"fmt"
	"io/ioutil"
	"os"

	"cuelang.org/go/cue"
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
	lin, _ := Lineage(rt)
	sch00, _ := lin.Schema(thema.SV(0, 0)) // we wouldn't normally hardcode this; just for the example

	// The first "instance" of example.json, conforming to the schema at 0.0
	i00, _ := sch00.Validate(exdata)
	datasourceStr, _ := i00.Underlying().LookupPath(cue.ParsePath("datasource")).String()
	fmt.Printf("original datasource string: %q\n", datasourceStr) // "foo"

	// now translate to 1.0, where we changed the datasource string field to a struct.
	// This is a case where Translate arguably should return an error, but it doesn't.
	i10, _ := i00.Translate(thema.SV(1, 0))
	if i10 == nil {
		fmt.Println("i00 translated to nil")
	}

	// We can get to the error by inspecting the underlying cue.Value. if `type`
	// is optional, this returns "not found" (makes sense!); if it is required,
	// we get:
	// #Translate.out.result.result.datasource.type: non-concrete value string
	dsType, err := i10.Underlying().LookupPath(cue.ParsePath("datasource.type")).String()
	if err != nil {
		fmt.Println(err)
	} else {
		print(dsType)
	}

	// The decoded values also look fine/expected! If type is required, we get a
	// zero-value datasource (type = ""); otherwise a nil string pointer (*string)(nil).
	var example Example
	i10.Underlying().Decode(&example)
	print(example.Datasource.Type)
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
