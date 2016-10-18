package controllers

import (
	"fmt"
	"test/any"
	"path/filepath"
	"os"
	"os/exec"
)

type Index struct {
	any.Ctx
}

func (ctx *Index)Index() {
	fmt.Fprint(ctx.Response, "/")

	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	println(path)
}

