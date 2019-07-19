package services

import (
	"../domains"
	"../utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Sites []domains.Site

func GetSitesList() (Sites, *utils.ApiError){

	var sitesList Sites

	res, err := http.Get(utils.UrlSites)
	if err != nil {
		return nil, &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	data, err :=  ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal(data, &sitesList); err != nil {
		if err != nil {
			return nil, &utils.ApiError{
				Message: err.Error(),
				Status: http.StatusInternalServerError,
			}
		}
	}

	return sitesList, nil
}

func GetSite(siteId string) (*domains.Site, *utils.ApiError){

	site := &domains.Site{
		Id: siteId,
	}

	if err := site.Get(); err != nil {
		return nil, err
	}

	return site, nil
}


