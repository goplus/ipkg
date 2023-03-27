// export by github.com/goplus/igop/cmd/qexp

package playground

import (
	q "fyne.io/fyne/v2/tools/playground"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "playground",
		Path: "fyne.io/fyne/v2/tools/playground",
		Deps: map[string]string{
			"bytes":                           "bytes",
			"encoding/base64":                 "base64",
			"fmt":                             "fmt",
			"fyne.io/fyne/v2":                 "fyne",
			"fyne.io/fyne/v2/driver/software": "software",
			"fyne.io/fyne/v2/test":            "test",
			"fyne.io/fyne/v2/theme":           "theme",
			"image":                           "image",
			"image/png":                       "png",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewSoftwareCanvas": reflect.ValueOf(q.NewSoftwareCanvas),
			"Render":            reflect.ValueOf(q.Render),
			"RenderCanvas":      reflect.ValueOf(q.RenderCanvas),
			"RenderWindow":      reflect.ValueOf(q.RenderWindow),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
