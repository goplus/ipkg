// export by github.com/goplus/igop/cmd/qexp

package cpu

import (
	q "golang.org/x/sys/cpu"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "cpu",
		Path: "golang.org/x/sys/cpu",
		Deps: map[string]string{
			"io/ioutil": "ioutil",
			"os":        "os",
			"runtime":   "runtime",
			"strings":   "strings",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"CacheLinePad": reflect.TypeOf((*q.CacheLinePad)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ARM":         reflect.ValueOf(&q.ARM),
			"ARM64":       reflect.ValueOf(&q.ARM64),
			"Initialized": reflect.ValueOf(&q.Initialized),
			"MIPS64X":     reflect.ValueOf(&q.MIPS64X),
			"PPC64":       reflect.ValueOf(&q.PPC64),
			"S390X":       reflect.ValueOf(&q.S390X),
			"X86":         reflect.ValueOf(&q.X86),
		},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
