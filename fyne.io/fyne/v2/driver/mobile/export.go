// export by github.com/goplus/igop/cmd/qexp

package mobile

import (
	q "fyne.io/fyne/v2/driver/mobile"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "mobile",
		Path: "fyne.io/fyne/v2/driver/mobile",
		Deps: map[string]string{
			"fyne.io/fyne/v2": "fyne",
		},
		Interfaces: map[string]reflect.Type{
			"Device":       reflect.TypeOf((*q.Device)(nil)).Elem(),
			"Keyboardable": reflect.TypeOf((*q.Keyboardable)(nil)).Elem(),
			"Touchable":    reflect.TypeOf((*q.Touchable)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"KeyboardType": reflect.TypeOf((*q.KeyboardType)(nil)).Elem(),
			"TouchEvent":   reflect.TypeOf((*q.TouchEvent)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs:      map[string]reflect.Value{},
		TypedConsts: map[string]igop.TypedConst{
			"DefaultKeyboard":    {reflect.TypeOf(q.DefaultKeyboard), constant.MakeInt64(int64(q.DefaultKeyboard))},
			"NumberKeyboard":     {reflect.TypeOf(q.NumberKeyboard), constant.MakeInt64(int64(q.NumberKeyboard))},
			"PasswordKeyboard":   {reflect.TypeOf(q.PasswordKeyboard), constant.MakeInt64(int64(q.PasswordKeyboard))},
			"SingleLineKeyboard": {reflect.TypeOf(q.SingleLineKeyboard), constant.MakeInt64(int64(q.SingleLineKeyboard))},
		},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
