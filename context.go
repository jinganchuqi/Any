package any

import (
	"net/http"
	"html/template"
)

type HttpContext struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Params         string
}

type Ctx struct {
	//httpContext HttpContext
	Response http.ResponseWriter
	Request  *http.Request
	Params   string
}

func (ctx *Ctx) Init(httpContext HttpContext) {
	ctx.Response = httpContext.ResponseWriter
	ctx.Request = httpContext.Request
	ctx.Params = httpContext.Params
}

func (ctx *Ctx) Render(pathString string, data interface{}) {
	t, err := template.ParseFiles(pathString)
	checkErr(err)
	err = t.Execute(ctx.Response, data)
	checkErr(err)
}
