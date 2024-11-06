package route

import (
	"gae-backend-test/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRankRouter(router *gin.RouterGroup) {
	rc := controller.NewRankController()
	rankGroup := router.Group("/rank")
	{
		rankGroup.GET("/hot/phase", rc.GetHotRankPhase)
		rankGroup.GET("/hot/:page/:phase", rc.GetHotRank)
		rankGroup.GET("/user/info/:username", rc.GetUserRank)
		rankGroup.GET("/user/info/specific/:username", rc.GetSpecificInfo)
		rankGroup.GET("/compare/:user1/:user2", rc.CompareInfo)
		rankGroup.GET("/level/:level/:score", rc.GetRankByGrade)
		rankGroup.GET("/tech/:tech/:score", rc.GetRankByTechStack)
		rankGroup.GET("/nation/:nation/:score", rc.GetRankByNation)
	}
}
