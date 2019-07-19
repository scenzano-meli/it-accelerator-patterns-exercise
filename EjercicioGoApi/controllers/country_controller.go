package controllers

import (
	"../services"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	countryID = "countryId"
)

func GetCountriesFromApi( c *gin.Context) {

	countriesList, apiError :=  services.GetCountriesList()
	if apiError != nil {
		return
	}

	c.JSON(http.StatusOK, countriesList)
}

func GetCountryFromApi( c *gin.Context) {

	id := c.Param(countryID)

	country, apiError :=  services.GetCountry(id)
	if apiError != nil {
		return
	}

	c.JSON(http.StatusOK, country)
}