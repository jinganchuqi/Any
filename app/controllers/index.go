package controllers

import (
	"any/any"
)

type Index struct {
	any.Ctx
}

func (ctx *Index) Index() {
	ctx.MakeJson(map[string]interface{}{
		"code": 200,
		"msg":  "SUCCESS",
		"data": map[string]interface{}{
			"userName": "Jiang",
		},
	})
}

func (ctx *Index) Test() {

}
