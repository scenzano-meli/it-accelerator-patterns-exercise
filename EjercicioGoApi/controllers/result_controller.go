package controllers

import (
	"../utils"
	"../services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUserResultFromApi(c *gin.Context) {

	userID := c.Param("userID")

	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		apiError := &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusBadRequest,
		}
		c.JSON(apiError.Status, apiError.Message)

		return
	}

	result, apiError :=  services.GetResult(id)
	if apiError != nil {
		return
	}

	c.JSON(http.StatusOK, result)
}
