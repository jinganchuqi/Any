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
			"userName": "子非鱼",
			"userAge" :33,
		},
	})
}

func (ctx *Index) Test() {

}
