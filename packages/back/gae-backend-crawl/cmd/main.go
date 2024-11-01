package main

func main() {
	app, err := InitializeApp()
	if err != nil {
		return
	}
	app.CrawlExecutor.Setup()
}
