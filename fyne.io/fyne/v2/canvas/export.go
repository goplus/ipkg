// export by github.com/goplus/igop/cmd/qexp

package canvas

import (
	q "fyne.io/fyne/v2/canvas"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "canvas",
		Path: "fyne.io/fyne/v2/canvas",
		Deps: map[string]string{
			"fyne.io/fyne/v2":         "fyne",
			"fyne.io/fyne/v2/storage": "storage",
			"image":                   "image",
			"image/color":             "color",
			"image/draw":              "draw",
			"io":                      "io",
			"io/ioutil":               "ioutil",
			"math":                    "math",
			"path/filepath":           "filepath",
			"sync":                    "sync",
			"time":                    "time",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Circle":         reflect.TypeOf((*q.Circle)(nil)).Elem(),
			"Image":          reflect.TypeOf((*q.Image)(nil)).Elem(),
			"ImageFill":      reflect.TypeOf((*q.ImageFill)(nil)).Elem(),
			"ImageScale":     reflect.TypeOf((*q.ImageScale)(nil)).Elem(),
			"Line":           reflect.TypeOf((*q.Line)(nil)).Elem(),
			"LinearGradient": reflect.TypeOf((*q.LinearGradient)(nil)).Elem(),
			"RadialGradient": reflect.TypeOf((*q.RadialGradient)(nil)).Elem(),
			"Raster":         reflect.TypeOf((*q.Raster)(nil)).Elem(),
			"Rectangle":      reflect.TypeOf((*q.Rectangle)(nil)).Elem(),
			"Text":           reflect.TypeOf((*q.Text)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewCircle":             reflect.ValueOf(q.NewCircle),
			"NewColorRGBAAnimation": reflect.ValueOf(q.NewColorRGBAAnimation),
			"NewHorizontalGradient": reflect.ValueOf(q.NewHorizontalGradient),
			"NewImageFromFile":      reflect.ValueOf(q.NewImageFromFile),
			"NewImageFromImage":     reflect.ValueOf(q.NewImageFromImage),
			"NewImageFromReader":    reflect.ValueOf(q.NewImageFromReader),
			"NewImageFromResource":  reflect.ValueOf(q.NewImageFromResource),
			"NewImageFromURI":       reflect.ValueOf(q.NewImageFromURI),
			"NewLine":               reflect.ValueOf(q.NewLine),
			"NewLinearGradient":     reflect.ValueOf(q.NewLinearGradient),
			"NewPositionAnimation":  reflect.ValueOf(q.NewPositionAnimation),
			"NewRadialGradient":     reflect.ValueOf(q.NewRadialGradient),
			"NewRaster":             reflect.ValueOf(q.NewRaster),
			"NewRasterFromImage":    reflect.ValueOf(q.NewRasterFromImage),
			"NewRasterWithPixels":   reflect.ValueOf(q.NewRasterWithPixels),
			"NewRectangle":          reflect.ValueOf(q.NewRectangle),
			"NewSizeAnimation":      reflect.ValueOf(q.NewSizeAnimation),
			"NewText":               reflect.ValueOf(q.NewText),
			"NewVerticalGradient":   reflect.ValueOf(q.NewVerticalGradient),
			"Refresh":               reflect.ValueOf(q.Refresh),
		},
		TypedConsts: map[string]igop.TypedConst{
			"DurationShort":     {reflect.TypeOf(q.DurationShort), constant.MakeInt64(int64(q.DurationShort))},
			"DurationStandard":  {reflect.TypeOf(q.DurationStandard), constant.MakeInt64(int64(q.DurationStandard))},
			"ImageFillContain":  {reflect.TypeOf(q.ImageFillContain), constant.MakeInt64(int64(q.ImageFillContain))},
			"ImageFillOriginal": {reflect.TypeOf(q.ImageFillOriginal), constant.MakeInt64(int64(q.ImageFillOriginal))},
			"ImageFillStretch":  {reflect.TypeOf(q.ImageFillStretch), constant.MakeInt64(int64(q.ImageFillStretch))},
			"ImageScaleFastest": {reflect.TypeOf(q.ImageScaleFastest), constant.MakeInt64(int64(q.ImageScaleFastest))},
			"ImageScalePixels":  {reflect.TypeOf(q.ImageScalePixels), constant.MakeInt64(int64(q.ImageScalePixels))},
			"ImageScaleSmooth":  {reflect.TypeOf(q.ImageScaleSmooth), constant.MakeInt64(int64(q.ImageScaleSmooth))},
		},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
