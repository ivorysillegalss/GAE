package bootstrap

import "github.com/gin-gonic/gin"

func Setup(app *Application) *gin.Engine {
	r := gin.Default()
	e := app.Executor
	e.CronExecutor.SetupCron()
	e.ConsumeExecutor.SetupConsume()
	e.DataExecutor.InitData()
	return r
}
