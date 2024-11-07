package controller

import (
	"gae-backend-test/constant/request"
	"gae-backend-test/domain"
	"gae-backend-test/usecase"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"net/http"
)

type RelationController struct {
	userUsecase domain.UserUsecase
}

func NewRelationController() *RelationController {
	return &RelationController{
		userUsecase: usecase.NewUsecase(),
	}
}

func (r *RelationController) GetUserRelationNetwork(c *gin.Context) {
	usernameParam := c.Param("username")
	userRelation := r.userUsecase.QueryUserNetwork(usernameParam)
	if funk.IsEmpty(userRelation) {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Code: request.GetError})
	} else {
		c.JSON(http.StatusOK, domain.SuccessResponse{Code: request.GetSuccess, Data: userRelation})
	}

}
