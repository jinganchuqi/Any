package routers

import (
	"test/app/controllers"
	"test/any"
)

func Routers() any.Routers {
	routers := any.Routers{}
	routers.Get("/@Index", &controllers.Index{})
	routers.Get("/test@Test", &controllers.Test{})

	routers.Post("/upload@Upload", &controllers.Test{})
	return routers
}
