package main

import "gae-backend-test/route"

func main() {

	setup := route.Setup()
	//gin.SetMode(gin.DebugMode)
	setup.Run(":3000")
}
