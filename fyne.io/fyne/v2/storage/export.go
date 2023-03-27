// export by github.com/goplus/igop/cmd/qexp

package storage

import (
	q "fyne.io/fyne/v2/storage"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "storage",
		Path: "fyne.io/fyne/v2/storage",
		Deps: map[string]string{
			"errors":                             "errors",
			"fyne.io/fyne/v2":                    "fyne",
			"fyne.io/fyne/v2/storage/repository": "repository",
			"io/ioutil":                          "ioutil",
			"strings":                            "strings",
		},
		Interfaces: map[string]reflect.Type{
			"FileFilter": reflect.TypeOf((*q.FileFilter)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"ExtensionFileFilter": reflect.TypeOf((*q.ExtensionFileFilter)(nil)).Elem(),
			"MimeTypeFileFilter":  reflect.TypeOf((*q.MimeTypeFileFilter)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrAlreadyExists": reflect.ValueOf(&q.ErrAlreadyExists),
			"ErrNotExists":     reflect.ValueOf(&q.ErrNotExists),
			"URIRootError":     reflect.ValueOf(&q.URIRootError),
		},
		Funcs: map[string]reflect.Value{
			"CanList":                reflect.ValueOf(q.CanList),
			"CanRead":                reflect.ValueOf(q.CanRead),
			"CanWrite":               reflect.ValueOf(q.CanWrite),
			"Child":                  reflect.ValueOf(q.Child),
			"Copy":                   reflect.ValueOf(q.Copy),
			"CreateListable":         reflect.ValueOf(q.CreateListable),
			"Delete":                 reflect.ValueOf(q.Delete),
			"Exists":                 reflect.ValueOf(q.Exists),
			"List":                   reflect.ValueOf(q.List),
			"ListerForURI":           reflect.ValueOf(q.ListerForURI),
			"LoadResourceFromURI":    reflect.ValueOf(q.LoadResourceFromURI),
			"Move":                   reflect.ValueOf(q.Move),
			"NewExtensionFileFilter": reflect.ValueOf(q.NewExtensionFileFilter),
			"NewFileURI":             reflect.ValueOf(q.NewFileURI),
			"NewMimeTypeFileFilter":  reflect.ValueOf(q.NewMimeTypeFileFilter),
			"NewURI":                 reflect.ValueOf(q.NewURI),
			"OpenFileFromURI":        reflect.ValueOf(q.OpenFileFromURI),
			"Parent":                 reflect.ValueOf(q.Parent),
			"ParseURI":               reflect.ValueOf(q.ParseURI),
			"Reader":                 reflect.ValueOf(q.Reader),
			"SaveFileToURI":          reflect.ValueOf(q.SaveFileToURI),
			"Writer":                 reflect.ValueOf(q.Writer),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
