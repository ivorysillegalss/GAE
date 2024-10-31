package main

import "gae-backend-crawl/crawl"

func main() {
	app, err := InitializeApp()
	if err != nil {
		return
	}
	crawl.Setup(app.Databases, app.PoolsFactory, app.Env, app.Bloom)
}
