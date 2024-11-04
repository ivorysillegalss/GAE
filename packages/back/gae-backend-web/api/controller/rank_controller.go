package controller

import (
	"gae-backend-web/constant/request"
	"gae-backend-web/domain"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"net/http"
	"strconv"
)

type RankController struct {
	rankUsecase domain.RankUsecase
}

func NewRankController(ru domain.RankUsecase) *RankController {
	return &RankController{rankUsecase: ru}
}

func (r *RankController) GetHotRankPhase(c *gin.Context) {
	phase := r.rankUsecase.GetHotRankPhase()
	if funk.IsEmpty(phase) {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Code: request.GetError})
	} else {
		c.JSON(http.StatusOK, domain.SuccessResponse{Code: request.GetSuccess, Data: phase})
	}
}

func (r *RankController) GetHotRank(c *gin.Context) {
	pageParam := c.Param("page")
	phaseParam := c.Param("phase")
	page, _ := strconv.Atoi(pageParam)
	hotRank := r.rankUsecase.GetHotRank(page, phaseParam)
	if funk.IsEmpty(hotRank) {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Code: request.GetError})
	} else {
		c.JSON(http.StatusOK, domain.SuccessResponse{Code: request.GetSuccess, Data: hotRank})
	}
}

func (r *RankController) GetUserRank(c *gin.Context) {
	usernameParam := c.Param("username")
	userRank := r.rankUsecase.SearchUserRank(usernameParam)
	if funk.IsEmpty(userRank) || len(*userRank) == 0 {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Code: request.GetError})
	} else {
		c.JSON(http.StatusOK, domain.SuccessResponse{Code: request.GetSuccess, Data: userRank})
	}
}

func (r *RankController) GetSpecificInfo(c *gin.Context) {

}

func (r *RankController) CompareInfo(c *gin.Context) {

}
