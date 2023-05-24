//
// Code generated by grafana-app-sdk. DO NOT EDIT.
//

package customkind

import (
	"embed"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"github.com/grafana/thema"
	"github.com/grafana/thema/load"
	"github.com/grafana/thema/vmux"
)

var rt = thema.NewRuntime(cuecontext.New())
var lineage thema.Lineage
var typedLatest thema.TypedSchema[*LineageType]
var muxer vmux.TypedMux[*LineageType]

// LineageType must be used instead of Object for thema.BindType, as the bound struct must exactly match the Schema,
// and Object has extra fields for metadata (which necessarily cannot be rendered in the Lineage/Schema,
// but must exist for the Object). This is essentially an "intermediate step" struct.
type LineageType struct {
	Spec     Spec     `json:"spec"`
	Status   Status   `json:"status"`
	Metadata Metadata `json:"metadata"`
}

func init() {
	var err error
	if lineage, err = Lineage(rt); err != nil {
		panic(err)
	}
	if typedLatest, err = thema.BindType(lineage.Latest(), &LineageType{}); err != nil {
		panic(err)
	}
	muxer = vmux.NewTypedMux[*LineageType](typedLatest, vmux.NewJSONCodec("input"))
}

//go:embed customkind_lineage.cue cue.mod/*
var modFS embed.FS

func loadLineage(lib *thema.Runtime) (cue.Value, error) {
	inst, err := load.InstanceWithThema(modFS, ".")
	if err != nil {
		return cue.Value{}, err
	}

	val := lib.Context().BuildInstance(inst)
	return val.LookupPath(cue.ParsePath("customkind")), nil
}

// Lineage constructs a Go handle representing the Customkind Object lineage,
// which includes the spec and all subresources.
func Lineage(rt *thema.Runtime, opts ...thema.BindOption) (thema.Lineage, error) {
	linval, err := loadLineage(rt)
	if err != nil {
		return nil, err
	}
	return thema.BindLineage(linval, rt, opts...)
}

var _ thema.LineageFactory = Lineage // Ensure our factory fulfills the type
