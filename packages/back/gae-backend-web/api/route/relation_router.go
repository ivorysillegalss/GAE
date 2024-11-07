package route

import (
	"gae-backend-web/bootstrap"
	"github.com/gin-gonic/gin"
)

func RegisterRelationRouter(router *gin.RouterGroup, c *bootstrap.Controllers) {
	rlc := c.RelationController
	relationGroup := router.Group("/relation")
	{
		relationGroup.GET("/:username", rlc.GetUserRelationNetwork)
	}
}
