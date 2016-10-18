package controllers

import (
	"test/any"
)

type Test struct {
	any.Ctx
}

func (ctx *Test) Test() {
	data := map[string]interface{}{
		"Title":"Hello",
		"Content":"Hello,World1",
	}
	ctx.Render("D:/Go/src/test/tpl/test.html", data)
}
