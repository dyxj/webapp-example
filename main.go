package main

import (
	"webapp-example/app"
)

func main() {
	wa := app.Apl{}
	// Prod
	wa.InitApp("./dist")
	//panic("fake error: test close")
	wa.Run(":8080")
}
