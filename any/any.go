package any

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strconv"
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
	fmt.Println("----start----")
	app.initEnvContext()
	app.initHttpContext(res, req)
	isStatic, statusCode := app.callStatic()
	fmt.Println(isStatic, statusCode)
	if isStatic == false {
		switch app.Request.Method {
		case "GET":
			statusCode = app.switchMethod(app.routers.GetMap)
		case "POST":
			statusCode = app.switchMethod(app.routers.PostMap)
		}
	}
	if statusCode != 200 {
		http.Error(app.Response, strconv.Itoa(statusCode), statusCode)
	}
	fmt.Println(isStatic, statusCode)
	fmt.Println("----end----")
}

/**
 静态文件处理
返回值 是否匹配到该路径，是否有该文件
*/
func (app *App) callStatic() (bool, int) {
	staticDir := app.RunPath + "/static/"
	UrlPath := app.Request.URL.Path
	if strings.HasPrefix(UrlPath, "/static") {
		filePath := staticDir + UrlPath[len("/static"):]
		fileInfo, err := os.Stat(filePath)
		if err != nil || os.IsNotExist(err) {
			return true, 404
		}
		if fileInfo.IsDir() {
			return true, 403
		}
		http.ServeFile(app.Response, app.Request, filePath)
		return true, 200
	}
	return false, 404
}

/**
路由调度
*/
func (app *App) switchMethod(routerMap map[string]interface{}) int {
	if len(routerMap) < 1 {
		return 404
	}
	method := strings.Trim(app.Ctx.HttpContext.Request.URL.Path, "/")
	for k, v := range routerMap {
		path := strings.Split(k, "@")
		if len(path) < 1 {
			continue
		}
		if method == strings.Trim(path[0], "/") {
			if len(path) != 2 {
				return 404
			}
			vt := reflect.ValueOf(v)
			in := make([]reflect.Value, 1)
			in[0] = reflect.ValueOf(app.Ctx)
			InitCtx := vt.MethodByName("InitCtx")
			if InitCtx.IsValid() == false {
				return 404
			}
			InitCtx.Call(in)
			in = make([]reflect.Value, 0)
			requestMethod := vt.MethodByName(path[1])
			if requestMethod.IsValid() == false {
				return 404
			}
			requestMethod.Call(in)
			return 200
			break
		}
	}
	return 404
}

/**
  初始 Http 上下文
*/
func (app *App) initHttp() {
	http.HandleFunc("/", app.callHttp)
	err := http.ListenAndServe(":88", nil)
	checkErr(err)
}

/**
初始化环境
*/
func (app *App) initEnvContext() {
	app.Ctx.EnvContext = EnvContext{RunPath: getRootDir()}
}

/**
初始化 http 上下文
*/
func (app *App) initHttpContext(res http.ResponseWriter, req *http.Request) {
	app.Ctx.HttpContext = HttpContext{Response: res, Request: req}
}
