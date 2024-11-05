package main

import "gae-backend-analysis/bootstrap"

func main() {
	app, err := InitializeApp()
	if err != nil {
		return
	}
	r := bootstrap.Setup(app)
	//gin.SetMode(gin.DebugMode)
	r.Run(app.Env.ServerAddress)
}
