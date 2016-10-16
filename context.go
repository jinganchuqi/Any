package any

import (
	"net/http"
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
