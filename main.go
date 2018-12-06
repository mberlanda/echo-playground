package main

import (
	"github.com/mberlanda/echo-playground/route"
)

func main() {
	// Hello, World!
	e := route.Init()
	e.Logger.Fatal(e.Start(":1323"))
}
