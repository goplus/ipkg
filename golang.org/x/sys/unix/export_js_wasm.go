// export by github.com/goplus/ixgo/cmd/qexp

package unix

import (
	_ "golang.org/x/sys/unix"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "unix",
		Path: "golang.org/x/sys/unix",
		Deps: map[string]string{},
	})
}
