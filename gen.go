//go:build ignore
// +build ignore

package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	pkgList = []string{
		"github.com/modern-go/reflect2",
		"golang.org/x/sys/execabs",
		"golang.org/x/image/vector",
		"golang.org/x/text/language",
	}
	sysList = []string{
		"golang.org/x/sys/cpu",
		"golang.org/x/sys/unix",
		"golang.org/x/sys/windows",
		"golang.org/x/sys/plan9",
	}
)

func main() {
	// ver := runtime.Version()[:6]
	var tags string
	var name string
	name = "export"
	// var fname string
	// switch ver {
	// case "go1.18":
	// 	tags = "//+build go1.18"
	// 	name = "go118_export"
	// 	fname = "go118_pkgs.go"
	// case "go1.17":
	// 	tags = "//+build go1.17,!go1.18"
	// 	name = "go117_export"
	// 	fname = "go117_pkgs.go"
	// case "go1.16":
	// 	tags = "//+build go1.16,!go1.17"
	// 	name = "go116_export"
	// 	fname = "go116_pkgs.go"
	// case "go1.15":
	// 	tags = "//+build go1.15,!go1.16"
	// 	name = "go115_export"
	// 	fname = "go115_pkgs.go"
	// case "go1.14":
	// 	tags = "//+build go1.14,!go1.15"
	// 	name = "go114_export"
	// 	fname = "go114_pkgs.go"
	// }

	pkgs := pkgList

	// log.Println(ver, name, tags)
	log.Println(pkgs)

	cmd := exec.Command("qexp", "-outdir", ".", "-addtags", tags, "-filename", name)
	cmd.Args = append(cmd.Args, pkgs...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	list := osarchList()
	for _, pkg := range sysList {
		for _, osarch := range list {
			ar := strings.Split(osarch, "_")
			if len(ar) != 2 {
				continue
			}
			log.Println(osarch)
			cmd := exec.Command("qexp", "-outdir", ".", "-addtags", tags, "-filename", name, "-contexts", osarch, pkg)
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			cmd.Env = os.Environ()
			cmd.Env = append(cmd.Env, "GOOS="+ar[0])
			cmd.Env = append(cmd.Env, "GOARCH="+ar[1])
			cmd.Env = append(cmd.Env, "GOEXPERIMENT=noregabi")
			cmd.Run()
		}
	}
}

//skip syscall
func stdList() []string {
	cmd := exec.Command("go", "list", "std")
	data, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	var ar []string
	for _, v := range strings.Split(string(data), "\n") {
		if v == "" {
			continue
		}
		if isSkipPkg(v) {
			continue
		}
		ar = append(ar, v)
	}
	return ar
}

func isSkipPkg(pkg string) bool {
	switch pkg {
	case "syscall", "unsafe":
		return true
	case "runtime/cgo", "runtime/race":
		return true
	case "time/tzdata":
		return true
	default:
		if strings.HasPrefix(pkg, "vendor/") {
			return true
		}
		for _, v := range strings.Split(pkg, "/") {
			if v == "internal" {
				return true
			}
		}
	}
	return false
}

func osarchList() []string {
	//go tool dist list
	cmd := exec.Command("go", "tool", "dist", "list")
	data, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	var ar []string
	for _, v := range strings.Split(string(data), "\n") {
		if v == "" {
			continue
		}
		ar = append(ar, strings.Replace(v, "/", "_", 1))
	}
	return ar
}
