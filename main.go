package main

import (
	"github.com/duguying/ojsite/routers"
	"github.com/labstack/echo/engine/standard"
)

func main() {
	e := routers.Init()

	e.Run(standard.New(":8088"))

}
