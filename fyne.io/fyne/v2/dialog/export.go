// export by github.com/goplus/igop/cmd/qexp

package dialog

import (
	q "fyne.io/fyne/v2/dialog"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "dialog",
		Path: "fyne.io/fyne/v2/dialog",
		Deps: map[string]string{
			"errors":                             "errors",
			"fmt":                                "fmt",
			"fyne.io/fyne/v2":                    "fyne",
			"fyne.io/fyne/v2/canvas":             "canvas",
			"fyne.io/fyne/v2/container":          "container",
			"fyne.io/fyne/v2/driver/desktop":     "desktop",
			"fyne.io/fyne/v2/internal/color":     "color",
			"fyne.io/fyne/v2/internal/widget":    "widget",
			"fyne.io/fyne/v2/layout":             "layout",
			"fyne.io/fyne/v2/storage":            "storage",
			"fyne.io/fyne/v2/storage/repository": "repository",
			"fyne.io/fyne/v2/theme":              "theme",
			"fyne.io/fyne/v2/widget":             "widget",
			"image":                              "image",
			"image/color":                        "color",
			"image/draw":                         "draw",
			"math":                               "math",
			"math/cmplx":                         "cmplx",
			"os":                                 "os",
			"path/filepath":                      "filepath",
			"runtime":                            "runtime",
			"strconv":                            "strconv",
			"strings":                            "strings",
			"sync/atomic":                        "atomic",
		},
		Interfaces: map[string]reflect.Type{
			"Dialog": reflect.TypeOf((*q.Dialog)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"ColorPickerDialog":      reflect.TypeOf((*q.ColorPickerDialog)(nil)).Elem(),
			"ConfirmDialog":          reflect.TypeOf((*q.ConfirmDialog)(nil)).Elem(),
			"EntryDialog":            reflect.TypeOf((*q.EntryDialog)(nil)).Elem(),
			"FileDialog":             reflect.TypeOf((*q.FileDialog)(nil)).Elem(),
			"ProgressDialog":         reflect.TypeOf((*q.ProgressDialog)(nil)).Elem(),
			"ProgressInfiniteDialog": reflect.TypeOf((*q.ProgressInfiniteDialog)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewColorPicker":      reflect.ValueOf(q.NewColorPicker),
			"NewConfirm":          reflect.ValueOf(q.NewConfirm),
			"NewCustom":           reflect.ValueOf(q.NewCustom),
			"NewCustomConfirm":    reflect.ValueOf(q.NewCustomConfirm),
			"NewEntryDialog":      reflect.ValueOf(q.NewEntryDialog),
			"NewError":            reflect.ValueOf(q.NewError),
			"NewFileOpen":         reflect.ValueOf(q.NewFileOpen),
			"NewFileSave":         reflect.ValueOf(q.NewFileSave),
			"NewFolderOpen":       reflect.ValueOf(q.NewFolderOpen),
			"NewForm":             reflect.ValueOf(q.NewForm),
			"NewInformation":      reflect.ValueOf(q.NewInformation),
			"NewProgress":         reflect.ValueOf(q.NewProgress),
			"NewProgressInfinite": reflect.ValueOf(q.NewProgressInfinite),
			"ShowColorPicker":     reflect.ValueOf(q.ShowColorPicker),
			"ShowConfirm":         reflect.ValueOf(q.ShowConfirm),
			"ShowCustom":          reflect.ValueOf(q.ShowCustom),
			"ShowCustomConfirm":   reflect.ValueOf(q.ShowCustomConfirm),
			"ShowEntryDialog":     reflect.ValueOf(q.ShowEntryDialog),
			"ShowError":           reflect.ValueOf(q.ShowError),
			"ShowFileOpen":        reflect.ValueOf(q.ShowFileOpen),
			"ShowFileSave":        reflect.ValueOf(q.ShowFileSave),
			"ShowFolderOpen":      reflect.ValueOf(q.ShowFolderOpen),
			"ShowForm":            reflect.ValueOf(q.ShowForm),
			"ShowInformation":     reflect.ValueOf(q.ShowInformation),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
