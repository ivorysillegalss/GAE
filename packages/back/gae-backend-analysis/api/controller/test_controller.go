package controller

import (
	"fmt"
	"gae-backend-analysis/bootstrap"
	"gae-backend-analysis/constant/request"
	"gae-backend-analysis/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
)

type TestController struct {
	env *bootstrap.Env
}

func (t *TestController) TestApi(c *gin.Context) {
	sprintf := fmt.Sprintf("Authorization: Bearer " + t.env.ApiOpenaiSecretKey)
	cmd := exec.Command("curl", "-i", "https://api.openai.com/v1/models",
		"-H", sprintf)
	// 获取命令输出
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}
	c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "开启聊天失败", Code: request.StartChatError, Data: output})
}

func NewTestController(env *bootstrap.Env) *TestController {
	return &TestController{env: env}
}
