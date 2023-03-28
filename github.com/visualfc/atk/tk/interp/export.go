// export by github.com/goplus/igop/cmd/qexp

package interp

import (
	q "github.com/visualfc/atk/tk/interp"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "interp",
		Path: "github.com/visualfc/atk/tk/interp",
		Deps: map[string]string{
			"errors":      "errors",
			"fmt":         "fmt",
			"image":       "image",
			"image/color": "color",
			"image/draw":  "draw",
			"os":          "os",
			"runtime/cgo": "cgo",
			"strconv":     "strconv",
			"sync":        "sync",
			"syscall":     "syscall",
			"unsafe":      "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"ActionMap":         reflect.TypeOf((*q.ActionMap)(nil)).Elem(),
			"CommandMap":        reflect.TypeOf((*q.CommandMap)(nil)).Elem(),
			"Interp":            reflect.TypeOf((*q.Interp)(nil)).Elem(),
			"ListObj":           reflect.TypeOf((*q.ListObj)(nil)).Elem(),
			"Obj":               reflect.TypeOf((*q.Obj)(nil)).Elem(),
			"Photo":             reflect.TypeOf((*q.Photo)(nil)).Elem(),
			"Tcl_QueuePosition": reflect.TypeOf((*q.Tcl_QueuePosition)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Async":         reflect.ValueOf(q.Async),
			"FindPhoto":     reflect.ValueOf(q.FindPhoto),
			"IsMainThread":  reflect.ValueOf(q.IsMainThread),
			"MainLoop":      reflect.ValueOf(q.MainLoop),
			"NewActionMap":  reflect.ValueOf(q.NewActionMap),
			"NewBoolObj":    reflect.ValueOf(q.NewBoolObj),
			"NewCommandMap": reflect.ValueOf(q.NewCommandMap),
			"NewFloat64Obj": reflect.ValueOf(q.NewFloat64Obj),
			"NewInt64Obj":   reflect.ValueOf(q.NewInt64Obj),
			"NewIntObj":     reflect.ValueOf(q.NewIntObj),
			"NewInterp":     reflect.ValueOf(q.NewInterp),
			"NewListObj":    reflect.ValueOf(q.NewListObj),
			"NewStringObj":  reflect.ValueOf(q.NewStringObj),
		},
		TypedConsts: map[string]igop.TypedConst{
			"TCL_QUEUE_HEAD": {reflect.TypeOf(q.TCL_QUEUE_HEAD), constant.MakeInt64(int64(q.TCL_QUEUE_HEAD))},
			"TCL_QUEUE_MARK": {reflect.TypeOf(q.TCL_QUEUE_MARK), constant.MakeInt64(int64(q.TCL_QUEUE_MARK))},
			"TCL_QUEUE_TAIL": {reflect.TypeOf(q.TCL_QUEUE_TAIL), constant.MakeInt64(int64(q.TCL_QUEUE_TAIL))},
		},
		UntypedConsts: map[string]igop.UntypedConst{
			"TCL_ALL_EVENTS":    {"untyped int", constant.MakeInt64(int64(q.TCL_ALL_EVENTS))},
			"TCL_DONT_WAIT":     {"untyped int", constant.MakeInt64(int64(q.TCL_DONT_WAIT))},
			"TCL_ERROR":         {"untyped int", constant.MakeInt64(int64(q.TCL_ERROR))},
			"TCL_FILE_EVENTS":   {"untyped int", constant.MakeInt64(int64(q.TCL_FILE_EVENTS))},
			"TCL_IDLE_EVENTS":   {"untyped int", constant.MakeInt64(int64(q.TCL_IDLE_EVENTS))},
			"TCL_OK":            {"untyped int", constant.MakeInt64(int64(q.TCL_OK))},
			"TCL_TIMER_EVENTS":  {"untyped int", constant.MakeInt64(int64(q.TCL_TIMER_EVENTS))},
			"TCL_WINDOW_EVENTS": {"untyped int", constant.MakeInt64(int64(q.TCL_WINDOW_EVENTS))},
		},
	})
}
