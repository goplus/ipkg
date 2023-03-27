// export by github.com/goplus/igop/cmd/qexp

package driver

import (
	q "fyne.io/fyne/v2/driver"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name:       "driver",
		Path:       "fyne.io/fyne/v2/driver",
		Deps:       map[string]string{},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"AndroidContext": reflect.TypeOf((*q.AndroidContext)(nil)).Elem(),
			"UnknownContext": reflect.TypeOf((*q.UnknownContext)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"RunNative": reflect.ValueOf(q.RunNative),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
