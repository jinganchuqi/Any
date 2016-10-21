package controllers

import (
	"any/any"
)

type Index struct {
	any.Ctx
}

func (ctx *Index) Index() {
	res := struct {
		Title   string
		Content string
	}{
		Title:   "Hello,World",
		Content: "小时候，老师经常会问我们的理想，有的想成为科学家，有的想成为大富豪，有的想成为航天员，而我从小就想成为一个有趣的人。",
	}

	//fmt.Println(res.Title)
	//ctx.MakeJson(res)
	ctx.Render("error", res)
	//ctx.Render("error",&res)
	//fmt.Fprint(ctx.Response,res.Content)
}

func (ctx *Index) Test() {

}
