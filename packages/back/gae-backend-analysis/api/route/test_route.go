package route

import (
	"gae-backend-analysis/bootstrap"
	"github.com/gin-gonic/gin"
)

func RegisterTestRouter(router *gin.RouterGroup, c *bootstrap.Controllers) {
	testGroup := router.Group("/test")
	testGroup.GET("", c.TestController.TestApi)
}
