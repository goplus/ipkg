// export by github.com/goplus/ixgo/cmd/qexp

package cpu

import (
	q "golang.org/x/sys/cpu"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "cpu",
		Path: "golang.org/x/sys/cpu",
		Deps: map[string]string{
			"os":      "os",
			"runtime": "runtime",
			"strconv": "strconv",
			"strings": "strings",
			"syscall": "syscall",
			"unsafe":  "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"CacheLinePad": reflect.TypeOf((*q.CacheLinePad)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{
			"Errno": reflect.TypeOf((*q.Errno)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{
			"ARM":         reflect.ValueOf(&q.ARM),
			"ARM64":       reflect.ValueOf(&q.ARM64),
			"Initialized": reflect.ValueOf(&q.Initialized),
			"Loong64":     reflect.ValueOf(&q.Loong64),
			"MIPS64X":     reflect.ValueOf(&q.MIPS64X),
			"PPC64":       reflect.ValueOf(&q.PPC64),
			"RISCV64":     reflect.ValueOf(&q.RISCV64),
			"S390X":       reflect.ValueOf(&q.S390X),
			"X86":         reflect.ValueOf(&q.X86),
		},
		Funcs:       map[string]reflect.Value{},
		TypedConsts: map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"IsBigEndian": {Typ: "untyped bool", Value: constant.MakeBool(bool(q.IsBigEndian))},
		},
	})
}
