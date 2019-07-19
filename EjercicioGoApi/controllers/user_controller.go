package controllers

import (
	"../services"
	"../utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	paramUserID = "userID"
)

func GetUserFromApi(c *gin.Context) {

	userID := c.Param(paramUserID)

	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		apiError := &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusBadRequest,
		}
		c.JSON(apiError.Status, apiError.Message)

		return
	}

	user, apiError :=  services.GetUser(id)
	if apiError != nil {
		return
	}

	c.JSON(http.StatusOK, user)
}
