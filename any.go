package any

import (
	"log"
	"net/http"
	"reflect"
	"strings"
)

type App struct {
	RoutersMap Routers
	Ctx        HttpContext
}

func (app *App) Run() {
	http.HandleFunc("/", app.callHttp)
	err := http.ListenAndServe(":88", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func (app *App) LoadRoute(routersMap Routers) *App {
	app.RoutersMap = routersMap
	return app
}

func (app *App) callHttp(res http.ResponseWriter, req *http.Request) {
	app.Ctx = HttpContext{ResponseWriter: res, Request: req}
	isFound := false
	switch req.Method {
	case "GET":
		isFound = app.switchMethod(app.RoutersMap.GetMap)
	case "POST":
		isFound = app.switchMethod(app.RoutersMap.PostMap)

	}
	if isFound == false {
		http.NotFound(res, req)
	}
}

func (app *App) switchMethod(routerMap map[string]interface{}) bool {
	if len(routerMap) < 1 {
		return false
	}
	method := strings.Trim(app.Ctx.Request.URL.Path, "/")
	for k, v := range routerMap {
		path := strings.Split(k, "@")
		if method == strings.Trim(path[0], "/") {

			vt := reflect.Indirect(reflect.ValueOf(v)).Type()
			vc := reflect.New(vt)
			init := vc.MethodByName("Init")
			in := make([]reflect.Value, 1)

			in[0] = reflect.ValueOf(app.Ctx)
			init.Call(in)

			in = make([]reflect.Value, 0)
			method := vc.MethodByName(path[1])
			method.Call(in)
			return true
			break
		}
	}
	return false
}
