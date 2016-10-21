package any

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type HttpContext struct {
	Response http.ResponseWriter
	Request  *http.Request
	Params   string
}

type EnvContext struct {
	RunPath string
}

type ErrContext struct {
}

type Ctx struct {
	HttpContext
	EnvContext
	Err ErrContext
}

/**
  初始化
*/
func (ctx *Ctx) InitCtx(c *Ctx) {
	ctx.Response = c.Response
	ctx.Request = c.Request
	ctx.Params = c.Params
	ctx.RunPath = c.RunPath
}

/**
渲染模板
*/
func (ctx *Ctx) Render(pathString string, data interface{}) {
	t, err := template.ParseFiles(ctx.RunPath + "/resource/tpl/" + pathString + ".html")
	checkErr(err)
	err = t.Execute(ctx.Response, data)
	checkErr(err)
}

func (ctx *Ctx) MakeJson(data interface{}) {
	jsonString, err := json.Marshal(data)
	checkErr(err)
	//ctx.Response.Header().Add("Content-Type","application/json; charset=utf-8")
	fmt.Fprint(ctx.Response, string(jsonString))
}

func (app *Ctx) CheckErr(err error) {
	if err != nil {
		err.Error()
	}
}

/**

 */
func (ctx *Ctx) ThrowHttp(errMsg string, code int) {
	//http.Error(ctx.Response,errMsg,http.StatusText(code))
	fmt.Println(errMsg, code)
}

func (e *Ctx) ThrowConsole(exceptMsg string) {
	fmt.Println(e)
}

func (e *Ctx) ThrowLog(exceptMsg string) {

}
