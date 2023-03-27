// export by github.com/goplus/igop/cmd/qexp

package layout

import (
	q "fyne.io/fyne/v2/layout"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "layout",
		Path: "fyne.io/fyne/v2/layout",
		Deps: map[string]string{
			"fyne.io/fyne/v2":        "fyne",
			"fyne.io/fyne/v2/canvas": "canvas",
			"fyne.io/fyne/v2/theme":  "theme",
			"math":                   "math",
		},
		Interfaces: map[string]reflect.Type{
			"SpacerObject": reflect.TypeOf((*q.SpacerObject)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Spacer": reflect.TypeOf((*q.Spacer)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewAdaptiveGridLayout":    reflect.ValueOf(q.NewAdaptiveGridLayout),
			"NewBorderLayout":          reflect.ValueOf(q.NewBorderLayout),
			"NewCenterLayout":          reflect.ValueOf(q.NewCenterLayout),
			"NewFormLayout":            reflect.ValueOf(q.NewFormLayout),
			"NewGridLayout":            reflect.ValueOf(q.NewGridLayout),
			"NewGridLayoutWithColumns": reflect.ValueOf(q.NewGridLayoutWithColumns),
			"NewGridLayoutWithRows":    reflect.ValueOf(q.NewGridLayoutWithRows),
			"NewGridWrapLayout":        reflect.ValueOf(q.NewGridWrapLayout),
			"NewHBoxLayout":            reflect.ValueOf(q.NewHBoxLayout),
			"NewMaxLayout":             reflect.ValueOf(q.NewMaxLayout),
			"NewPaddedLayout":          reflect.ValueOf(q.NewPaddedLayout),
			"NewSpacer":                reflect.ValueOf(q.NewSpacer),
			"NewVBoxLayout":            reflect.ValueOf(q.NewVBoxLayout),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
