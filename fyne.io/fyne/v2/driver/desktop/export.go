// export by github.com/goplus/igop/cmd/qexp

package desktop

import (
	q "fyne.io/fyne/v2/driver/desktop"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "desktop",
		Path: "fyne.io/fyne/v2/driver/desktop",
		Deps: map[string]string{
			"fyne.io/fyne/v2": "fyne",
			"image":           "image",
			"runtime":         "runtime",
			"strings":         "strings",
		},
		Interfaces: map[string]reflect.Type{
			"App":        reflect.TypeOf((*q.App)(nil)).Elem(),
			"Canvas":     reflect.TypeOf((*q.Canvas)(nil)).Elem(),
			"Cursor":     reflect.TypeOf((*q.Cursor)(nil)).Elem(),
			"Cursorable": reflect.TypeOf((*q.Cursorable)(nil)).Elem(),
			"Driver":     reflect.TypeOf((*q.Driver)(nil)).Elem(),
			"Hoverable":  reflect.TypeOf((*q.Hoverable)(nil)).Elem(),
			"Keyable":    reflect.TypeOf((*q.Keyable)(nil)).Elem(),
			"Mouseable":  reflect.TypeOf((*q.Mouseable)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"CustomShortcut": reflect.TypeOf((*q.CustomShortcut)(nil)).Elem(),
			"MouseButton":    reflect.TypeOf((*q.MouseButton)(nil)).Elem(),
			"MouseEvent":     reflect.TypeOf((*q.MouseEvent)(nil)).Elem(),
			"StandardCursor": reflect.TypeOf((*q.StandardCursor)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{
			"Modifier": reflect.TypeOf((*q.Modifier)(nil)).Elem(),
		},
		Vars:  map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{},
		TypedConsts: map[string]igop.TypedConst{
			"AltModifier":          {reflect.TypeOf(q.AltModifier), constant.MakeInt64(int64(q.AltModifier))},
			"ControlModifier":      {reflect.TypeOf(q.ControlModifier), constant.MakeInt64(int64(q.ControlModifier))},
			"CrosshairCursor":      {reflect.TypeOf(q.CrosshairCursor), constant.MakeInt64(int64(q.CrosshairCursor))},
			"DefaultCursor":        {reflect.TypeOf(q.DefaultCursor), constant.MakeInt64(int64(q.DefaultCursor))},
			"HResizeCursor":        {reflect.TypeOf(q.HResizeCursor), constant.MakeInt64(int64(q.HResizeCursor))},
			"HiddenCursor":         {reflect.TypeOf(q.HiddenCursor), constant.MakeInt64(int64(q.HiddenCursor))},
			"KeyAltLeft":           {reflect.TypeOf(q.KeyAltLeft), constant.MakeString(string(q.KeyAltLeft))},
			"KeyAltRight":          {reflect.TypeOf(q.KeyAltRight), constant.MakeString(string(q.KeyAltRight))},
			"KeyCapsLock":          {reflect.TypeOf(q.KeyCapsLock), constant.MakeString(string(q.KeyCapsLock))},
			"KeyControlLeft":       {reflect.TypeOf(q.KeyControlLeft), constant.MakeString(string(q.KeyControlLeft))},
			"KeyControlRight":      {reflect.TypeOf(q.KeyControlRight), constant.MakeString(string(q.KeyControlRight))},
			"KeyMenu":              {reflect.TypeOf(q.KeyMenu), constant.MakeString(string(q.KeyMenu))},
			"KeyNone":              {reflect.TypeOf(q.KeyNone), constant.MakeString(string(q.KeyNone))},
			"KeyPrintScreen":       {reflect.TypeOf(q.KeyPrintScreen), constant.MakeString(string(q.KeyPrintScreen))},
			"KeyShiftLeft":         {reflect.TypeOf(q.KeyShiftLeft), constant.MakeString(string(q.KeyShiftLeft))},
			"KeyShiftRight":        {reflect.TypeOf(q.KeyShiftRight), constant.MakeString(string(q.KeyShiftRight))},
			"KeySuperLeft":         {reflect.TypeOf(q.KeySuperLeft), constant.MakeString(string(q.KeySuperLeft))},
			"KeySuperRight":        {reflect.TypeOf(q.KeySuperRight), constant.MakeString(string(q.KeySuperRight))},
			"LeftMouseButton":      {reflect.TypeOf(q.LeftMouseButton), constant.MakeInt64(int64(q.LeftMouseButton))},
			"MouseButtonPrimary":   {reflect.TypeOf(q.MouseButtonPrimary), constant.MakeInt64(int64(q.MouseButtonPrimary))},
			"MouseButtonSecondary": {reflect.TypeOf(q.MouseButtonSecondary), constant.MakeInt64(int64(q.MouseButtonSecondary))},
			"MouseButtonTertiary":  {reflect.TypeOf(q.MouseButtonTertiary), constant.MakeInt64(int64(q.MouseButtonTertiary))},
			"PointerCursor":        {reflect.TypeOf(q.PointerCursor), constant.MakeInt64(int64(q.PointerCursor))},
			"RightMouseButton":     {reflect.TypeOf(q.RightMouseButton), constant.MakeInt64(int64(q.RightMouseButton))},
			"ShiftModifier":        {reflect.TypeOf(q.ShiftModifier), constant.MakeInt64(int64(q.ShiftModifier))},
			"SuperModifier":        {reflect.TypeOf(q.SuperModifier), constant.MakeInt64(int64(q.SuperModifier))},
			"TextCursor":           {reflect.TypeOf(q.TextCursor), constant.MakeInt64(int64(q.TextCursor))},
			"VResizeCursor":        {reflect.TypeOf(q.VResizeCursor), constant.MakeInt64(int64(q.VResizeCursor))},
		},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
