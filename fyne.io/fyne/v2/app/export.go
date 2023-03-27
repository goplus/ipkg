// export by github.com/goplus/igop/cmd/qexp

package app

import (
	q "fyne.io/fyne/v2/app"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "app",
		Path: "fyne.io/fyne/v2/app",
		Deps: map[string]string{
			"bytes":                                "bytes",
			"encoding/json":                        "json",
			"fmt":                                  "fmt",
			"fyne.io/fyne/v2":                      "fyne",
			"fyne.io/fyne/v2/internal":             "internal",
			"fyne.io/fyne/v2/internal/app":         "app",
			"fyne.io/fyne/v2/internal/driver/glfw": "glfw",
			"fyne.io/fyne/v2/internal/repository":  "repository",
			"fyne.io/fyne/v2/storage":              "storage",
			"fyne.io/fyne/v2/storage/repository":   "repository",
			"fyne.io/fyne/v2/theme":                "theme",
			"github.com/fsnotify/fsnotify":         "fsnotify",
			"golang.org/x/sys/execabs":             "execabs",
			"io":                                   "io",
			"net/url":                              "url",
			"os":                                   "os",
			"path/filepath":                        "filepath",
			"runtime/cgo":                          "cgo",
			"strconv":                              "strconv",
			"strings":                              "strings",
			"sync":                                 "sync",
			"sync/atomic":                          "atomic",
			"syscall":                              "syscall",
			"time":                                 "time",
			"unsafe":                               "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"SettingsSchema": reflect.TypeOf((*q.SettingsSchema)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"New":         reflect.ValueOf(q.New),
			"NewWithID":   reflect.ValueOf(q.NewWithID),
			"SetMetadata": reflect.ValueOf(q.SetMetadata),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
