package main

import "gae-backend-storage/api/route"

func main() {
	app, err := InitializeApp()
	if err != nil {
		return
	}

	setup := route.Setup(app.Executor)
	//gin.SetMode(gin.DebugMode)
	setup.Run(app.Env.ServerAddress)
}
