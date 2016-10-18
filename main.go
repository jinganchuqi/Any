package main

import (
	"test/app/routers"
	"test/any"
)

func main() {
	any := any.App{}
	any.LoadRoute(routers.Routers())
	any.Run()
}
