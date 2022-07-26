// export by github.com/goplus/igop/cmd/qexp

package execabs

import (
	q "golang.org/x/sys/execabs"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "execabs",
		Path: "golang.org/x/sys/execabs",
		Deps: map[string]string{
			"context":       "context",
			"fmt":           "fmt",
			"os/exec":       "exec",
			"path/filepath": "filepath",
			"reflect":       "reflect",
			"unsafe":        "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{
			"Cmd":       reflect.TypeOf((*q.Cmd)(nil)).Elem(),
			"Error":     reflect.TypeOf((*q.Error)(nil)).Elem(),
			"ExitError": reflect.TypeOf((*q.ExitError)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{
			"ErrNotFound": reflect.ValueOf(&q.ErrNotFound),
		},
		Funcs: map[string]reflect.Value{
			"Command":        reflect.ValueOf(q.Command),
			"CommandContext": reflect.ValueOf(q.CommandContext),
			"LookPath":       reflect.ValueOf(q.LookPath),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
