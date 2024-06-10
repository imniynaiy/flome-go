package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/theoriz0/flome-go/internal/log"
	"github.com/theoriz0/flome-go/internal/model"
	"github.com/theoriz0/flome-go/internal/service"
)

func AddUserRoutes(rg *gin.RouterGroup) {
	rg.POST("/login", userLoginController)
}

func userLoginController(c *gin.Context) {
	var user model.User
	err := c.ShouldBindBodyWithJSON(&user)
	if err != nil {
		log.Error("Failed to parse login user", log.String("err", err.Error()))
		c.Status(http.StatusBadRequest)
		return
	}
	token, err := service.Login(&user)
	if err != nil {
		log.Error("Failed to login", log.String("err", err.Error()))
		c.Status(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, &model.LoginResp{Token: token})
}
