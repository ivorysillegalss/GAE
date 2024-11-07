package controller

import (
	"gae-backend-web/constant/request"
	"gae-backend-web/domain"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"net/http"
)

type RelationController struct {
	userUsecase domain.UserUsecase
}

func NewRelationController(uu domain.UserUsecase) *RelationController {
	return &RelationController{userUsecase: uu}
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
