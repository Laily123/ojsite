package routers

import (
	"github.com/duguying/ojsite/controllers"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Debug()

	e.Get("/", controllers.Index)

	return e

	//User()
	//Problem()
	//Teacher()
	//Api()
	//Teach()
}
