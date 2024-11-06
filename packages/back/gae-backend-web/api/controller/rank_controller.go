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
	userUsecase domain.UserUsecase
}

func NewRankController(ru domain.RankUsecase, uu domain.UserUsecase) *RankController {
	return &RankController{rankUsecase: ru, userUsecase: uu}
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
	usernameParam := c.Param("username")
	contributorsInfo := r.userUsecase.QueryUserSpecificInfo(usernameParam)
	if funk.IsEmpty(contributorsInfo) || len(*contributorsInfo) == 0 {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Code: request.GetError})
	} else {
		c.JSON(http.StatusOK, domain.SuccessResponse{Code: request.GetSuccess, Data: contributorsInfo})
	}
}

func (r *RankController) GetRankByNation(c *gin.Context) {
	nationParam := c.Param("nation")
	scoreParam := c.Param("score")
	score, _ := strconv.Atoi(scoreParam)
	queryNation := r.userUsecase.QueryByNation(nationParam, score)
	if funk.IsEmpty(queryNation) || len(*queryNation) == 0 {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Code: request.GetError})
	} else {
		c.JSON(http.StatusOK, domain.SuccessResponse{Code: request.GetSuccess, Data: queryNation})
	}
}

func (r *RankController) GetRankByTechStack(c *gin.Context) {
	techParam := c.Param("tech")
	scoreParam := c.Param("score")
	score, _ := strconv.Atoi(scoreParam)
	queryTech := r.userUsecase.QueryByTech(techParam, score)
	if funk.IsEmpty(queryTech) || len(*queryTech) == 0 {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Code: request.GetError})
	} else {
		c.JSON(http.StatusOK, domain.SuccessResponse{Code: request.GetSuccess, Data: queryTech})
	}
}

func (r *RankController) GetRankByGrade(c *gin.Context) {
	techParam := c.Param("level")
	scoreParam := c.Param("score")
	score, _ := strconv.Atoi(scoreParam)
	queryTech := r.userUsecase.QueryByTech(techParam, score)
	if funk.IsEmpty(queryTech) || len(*queryTech) == 0 {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Code: request.GetError})
	} else {
		c.JSON(http.StatusOK, domain.SuccessResponse{Code: request.GetSuccess, Data: queryTech})
	}
}

func (r *RankController) CompareInfo(c *gin.Context) {
}
