// export by github.com/goplus/igop/cmd/qexp

package validation

import (
	q "fyne.io/fyne/v2/data/validation"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "validation",
		Path: "fyne.io/fyne/v2/data/validation",
		Deps: map[string]string{
			"errors":          "errors",
			"fyne.io/fyne/v2": "fyne",
			"regexp":          "regexp",
			"time":            "time",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewAllStrings": reflect.ValueOf(q.NewAllStrings),
			"NewRegexp":     reflect.ValueOf(q.NewRegexp),
			"NewTime":       reflect.ValueOf(q.NewTime),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
