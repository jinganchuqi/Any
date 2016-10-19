package any

import (
	"net/http"
	"html/template"
	"fmt"
	"encoding/json"
)


type HttpContext struct {
	Response       http.ResponseWriter
	Request        *http.Request
	Params         string
}

type EnvContext struct {
	RunPath string
}

type Ctx struct {
	HttpContext
	EnvContext
}

/**
  初始化
 */
func (ctx *Ctx) InitCtx(c Ctx) {
	ctx.Response = c.Response
	ctx.Request = c.Request
	ctx.Params = c.Params
	ctx.RunPath = c.RunPath
}

/**
 渲染模板
 */
func (ctx *Ctx) Render(pathString string, data interface{}) {
	t, err := template.ParseFiles(ctx.RunPath+"/resource/tpl/"+pathString+".html")
	checkErr(err)
	err = t.Execute(ctx.Response, data)
	checkErr(err)
}

func (ctx *Ctx) MakeJson(data interface{}){
	jsonString, err := json.Marshal(data)
	checkErr(err)
	fmt.Fprint(ctx.Response,string(jsonString))
}

func (app *Ctx) CheckErr(err error) {
	if err != nil {
		err.Error()
	}
}


