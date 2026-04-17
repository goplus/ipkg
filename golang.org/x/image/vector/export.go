// export by github.com/goplus/ixgo/cmd/qexp

package vector

import (
	q "golang.org/x/image/vector"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "vector",
		Path: "golang.org/x/image/vector",
		Deps: map[string]string{
			"image":       "image",
			"image/color": "color",
			"image/draw":  "draw",
			"math":        "math",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Rasterizer": reflect.TypeOf((*q.Rasterizer)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewRasterizer": reflect.ValueOf(q.NewRasterizer),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
