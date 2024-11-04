package route

import (
	"gae-backend-web/bootstrap"
	"github.com/gin-gonic/gin"
)

func RegisterRankRouter(router *gin.RouterGroup, c *bootstrap.Controllers) {
	rc := c.RankController
	rankGroup := router.Group("/rank")
	{
		rankGroup.GET("/hot/phase", rc.GetHotRankPhase)
		rankGroup.GET("/hot/:page/:phase", rc.GetHotRank)
		rankGroup.GET("/user/info/:username", rc.GetUserRank)
		rankGroup.GET("/user/info/specific/:username", rc.GetSpecificInfo)
		rankGroup.GET("/compare/:user1/:user2", rc.CompareInfo)
	}
}
