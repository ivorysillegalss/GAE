package route

import (
	"gae-backend-test/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRelationRouter(router *gin.RouterGroup) {
	rlc := controller.NewRelationController()
	relationGroup := router.Group("/relation")
	{
		relationGroup.GET("/:username", rlc.GetUserRelationNetwork)
	}
}
