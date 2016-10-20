package controllers

import (
	"any/any"
	"fmt"
	"io"
	"os"
)

type Test struct {
	any.Ctx
}

func (ctx *Test) Test() {
	data := map[string]interface{}{
		"Title":   "Hello",
		"Content": "Hello,World1",
	}
	ctx.Render("test", data)
}

func (ctx *Test) Blog() {
	ctx.Render("blog", struct {
	}{})
}

func (ctx *Test) Upload() {
	file, handle, err := ctx.Request.FormFile("file")
	ctx.CheckErr(err)
	f, err := os.OpenFile(ctx.RunPath+"/resource/storage/"+handle.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	ctx.CheckErr(err)
	io.Copy(f, file)
	ctx.CheckErr(err)
	defer f.Close()
	defer file.Close()
	fmt.Println("upload success")
}
