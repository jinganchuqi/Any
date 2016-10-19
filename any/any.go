package any

import (
	"net/http"
	"reflect"
	"strings"
)

type App struct {
	Ctx
	routers Routers
}

/**
	项目初始化
 */
func (app *App) Run() {
	app.initEnv()
	app.initHttp()
}

/**
 加载
 */
func (app *App) LoadRoute(routersMap Routers) *App {
	app.routers = routersMap
	return app
}

/**
  响应请求
 */
func (app *App) callHttp(res http.ResponseWriter, req *http.Request) {
	app.Ctx.HttpContext = HttpContext{Response:res,Request:req}
	isFound := false
	switch req.Method {
	case "GET":
		isFound = app.switchMethod(app.routers.GetMap)
	case "POST":
		isFound = app.switchMethod(app.routers.PostMap)
	}
	if isFound == false {
		http.NotFound(res, req)
	}
}

/**
 路由调度
 */
func (app *App) switchMethod(routerMap map[string]interface{}) bool {
	if len(routerMap) < 1 {
		return false
	}
	method := strings.Trim(app.Ctx.HttpContext.Request.URL.Path, "/")
	for k, v := range routerMap {
		path := strings.Split(k, "@")
		if method == strings.Trim(path[0], "/") {

			vt := reflect.Indirect(reflect.ValueOf(v)).Type()
			vc := reflect.New(vt)
			init := vc.MethodByName("InitCtx")

			in := make([]reflect.Value, 1)
			in[0] = reflect.ValueOf(app.Ctx)
			init.Call(in)

			in = make([]reflect.Value, 0)
			opt := vc.MethodByName(path[1])

			if opt.IsValid() == true {
				opt.Call(in)
				return true
			}
			break
		}
	}
	return false
}

/**
  初始 Http 上下文
 */
func (app *App) initHttp(){
	http.HandleFunc("/", app.callHttp)
	err := http.ListenAndServe(":88", nil)
	checkErr(err)
}

/**
 初始化环境
 */
func (app *App) initEnv(){
	app.Ctx.EnvContext = EnvContext{RunPath:getRootDir()}
}




