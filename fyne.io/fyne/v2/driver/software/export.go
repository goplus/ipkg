// export by github.com/goplus/igop/cmd/qexp

package software

import (
	q "fyne.io/fyne/v2/driver/software"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "software",
		Path: "fyne.io/fyne/v2/driver/software",
		Deps: map[string]string{
			"fyne.io/fyne/v2":                           "fyne",
			"fyne.io/fyne/v2/internal/app":              "app",
			"fyne.io/fyne/v2/internal/painter/software": "software",
			"fyne.io/fyne/v2/test":                      "test",
			"image":                                     "image",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewCanvas":            reflect.ValueOf(q.NewCanvas),
			"NewTransparentCanvas": reflect.ValueOf(q.NewTransparentCanvas),
			"Render":               reflect.ValueOf(q.Render),
			"RenderCanvas":         reflect.ValueOf(q.RenderCanvas),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
