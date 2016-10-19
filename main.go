package main

import (
	"any/any"
	"any/app/routers"
)

func main() {
	any := any.App{}
	any.LoadRoute(routers.Routers())
	any.Run()
}
