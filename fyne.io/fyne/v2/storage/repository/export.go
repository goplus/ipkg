// export by github.com/goplus/igop/cmd/qexp

package repository

import (
	q "fyne.io/fyne/v2/storage/repository"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "repository",
		Path: "fyne.io/fyne/v2/storage/repository",
		Deps: map[string]string{
			"bufio":                 "bufio",
			"errors":                "errors",
			"fmt":                   "fmt",
			"fyne.io/fyne/v2":       "fyne",
			"github.com/fredbi/uri": "uri",
			"io":                    "io",
			"mime":                  "mime",
			"path/filepath":         "filepath",
			"runtime":               "runtime",
			"strings":               "strings",
			"unicode/utf8":          "utf8",
		},
		Interfaces: map[string]reflect.Type{
			"CopyableRepository":     reflect.TypeOf((*q.CopyableRepository)(nil)).Elem(),
			"CustomURIRepository":    reflect.TypeOf((*q.CustomURIRepository)(nil)).Elem(),
			"HierarchicalRepository": reflect.TypeOf((*q.HierarchicalRepository)(nil)).Elem(),
			"ListableRepository":     reflect.TypeOf((*q.ListableRepository)(nil)).Elem(),
			"MovableRepository":      reflect.TypeOf((*q.MovableRepository)(nil)).Elem(),
			"Repository":             reflect.TypeOf((*q.Repository)(nil)).Elem(),
			"WritableRepository":     reflect.TypeOf((*q.WritableRepository)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrOperationNotSupported": reflect.ValueOf(&q.ErrOperationNotSupported),
			"ErrURIRoot":               reflect.ValueOf(&q.ErrURIRoot),
		},
		Funcs: map[string]reflect.Value{
			"ForScheme":     reflect.ValueOf(q.ForScheme),
			"ForURI":        reflect.ValueOf(q.ForURI),
			"GenericChild":  reflect.ValueOf(q.GenericChild),
			"GenericCopy":   reflect.ValueOf(q.GenericCopy),
			"GenericMove":   reflect.ValueOf(q.GenericMove),
			"GenericParent": reflect.ValueOf(q.GenericParent),
			"NewFileURI":    reflect.ValueOf(q.NewFileURI),
			"ParseURI":      reflect.ValueOf(q.ParseURI),
			"Register":      reflect.ValueOf(q.Register),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
