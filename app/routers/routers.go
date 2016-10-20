package routers

import (
	"any/any"
	"any/app/controllers"
)

func Routers() any.Routers {
	routers := any.Routers{}
	routers.Get("/@Index", &controllers.Index{})
	routers.Get("/test@Test", &controllers.Test{})

	routers.Get("/test/:{id}@Test", &controllers.Test{})
	routers.Get("test/blog/@Blog", &controllers.Test{})

	routers.Post("/test/upload@Upload", &controllers.Test{})
	return routers
}
