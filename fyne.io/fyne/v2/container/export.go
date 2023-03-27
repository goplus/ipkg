// export by github.com/goplus/igop/cmd/qexp

package container

import (
	q "fyne.io/fyne/v2/container"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "container",
		Path: "fyne.io/fyne/v2/container",
		Deps: map[string]string{
			"fyne.io/fyne/v2":                 "fyne",
			"fyne.io/fyne/v2/canvas":          "canvas",
			"fyne.io/fyne/v2/driver/desktop":  "desktop",
			"fyne.io/fyne/v2/internal":        "internal",
			"fyne.io/fyne/v2/internal/widget": "widget",
			"fyne.io/fyne/v2/layout":          "layout",
			"fyne.io/fyne/v2/theme":           "theme",
			"fyne.io/fyne/v2/widget":          "widget",
			"image/color":                     "color",
			"sync":                            "sync",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"AppTabs":     reflect.TypeOf((*q.AppTabs)(nil)).Elem(),
			"DocTabs":     reflect.TypeOf((*q.DocTabs)(nil)).Elem(),
			"Split":       reflect.TypeOf((*q.Split)(nil)).Elem(),
			"TabItem":     reflect.TypeOf((*q.TabItem)(nil)).Elem(),
			"TabLocation": reflect.TypeOf((*q.TabLocation)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{
			"Scroll":          reflect.TypeOf((*q.Scroll)(nil)).Elem(),
			"ScrollDirection": reflect.TypeOf((*q.ScrollDirection)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"New":                reflect.ValueOf(q.New),
			"NewAdaptiveGrid":    reflect.ValueOf(q.NewAdaptiveGrid),
			"NewAppTabs":         reflect.ValueOf(q.NewAppTabs),
			"NewBorder":          reflect.ValueOf(q.NewBorder),
			"NewCenter":          reflect.ValueOf(q.NewCenter),
			"NewDocTabs":         reflect.ValueOf(q.NewDocTabs),
			"NewGridWithColumns": reflect.ValueOf(q.NewGridWithColumns),
			"NewGridWithRows":    reflect.ValueOf(q.NewGridWithRows),
			"NewGridWrap":        reflect.ValueOf(q.NewGridWrap),
			"NewHBox":            reflect.ValueOf(q.NewHBox),
			"NewHScroll":         reflect.ValueOf(q.NewHScroll),
			"NewHSplit":          reflect.ValueOf(q.NewHSplit),
			"NewMax":             reflect.ValueOf(q.NewMax),
			"NewPadded":          reflect.ValueOf(q.NewPadded),
			"NewScroll":          reflect.ValueOf(q.NewScroll),
			"NewTabItem":         reflect.ValueOf(q.NewTabItem),
			"NewTabItemWithIcon": reflect.ValueOf(q.NewTabItemWithIcon),
			"NewVBox":            reflect.ValueOf(q.NewVBox),
			"NewVScroll":         reflect.ValueOf(q.NewVScroll),
			"NewVSplit":          reflect.ValueOf(q.NewVSplit),
			"NewWithoutLayout":   reflect.ValueOf(q.NewWithoutLayout),
		},
		TypedConsts: map[string]igop.TypedConst{
			"ScrollBoth":           {reflect.TypeOf(q.ScrollBoth), constant.MakeInt64(int64(q.ScrollBoth))},
			"ScrollHorizontalOnly": {reflect.TypeOf(q.ScrollHorizontalOnly), constant.MakeInt64(int64(q.ScrollHorizontalOnly))},
			"ScrollNone":           {reflect.TypeOf(q.ScrollNone), constant.MakeInt64(int64(q.ScrollNone))},
			"ScrollVerticalOnly":   {reflect.TypeOf(q.ScrollVerticalOnly), constant.MakeInt64(int64(q.ScrollVerticalOnly))},
			"TabLocationBottom":    {reflect.TypeOf(q.TabLocationBottom), constant.MakeInt64(int64(q.TabLocationBottom))},
			"TabLocationLeading":   {reflect.TypeOf(q.TabLocationLeading), constant.MakeInt64(int64(q.TabLocationLeading))},
			"TabLocationTop":       {reflect.TypeOf(q.TabLocationTop), constant.MakeInt64(int64(q.TabLocationTop))},
			"TabLocationTrailing":  {reflect.TypeOf(q.TabLocationTrailing), constant.MakeInt64(int64(q.TabLocationTrailing))},
		},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
