package controllers

import (
	"../services"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	siteID = "siteID"
)

func GetSitesListFromApi( c *gin.Context) {

	sitesList, apiError :=  services.GetSitesList()
	if apiError != nil {
		return
	}

	c.JSON(http.StatusOK, sitesList)
}

func GetSiteFromApi( c *gin.Context) {

	id := c.Param(siteID)

	site, apiError :=  services.GetSite(id)
	if apiError != nil {
		return
	}

	c.JSON(http.StatusOK, site)
}
