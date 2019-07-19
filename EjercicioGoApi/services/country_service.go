package services

import (
	"../domains"
	"../utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Countries []domains.Country

func GetCountriesList() (Countries, *utils.ApiError){

	var countriesList Countries

	res, err := http.Get(utils.UrlContries)
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

	if err := json.Unmarshal(data, &countriesList); err != nil {
		if err != nil {
			return nil, &utils.ApiError{
				Message: err.Error(),
				Status: http.StatusInternalServerError,
			}
		}
	}

	return countriesList, nil
}

func GetCountry(countryId string) (*domains.Country, *utils.ApiError){

	country := &domains.Country{
		Id: countryId,
	}

	if err := country.Get(); err != nil {
		return nil, err
	}

	return country, nil
}
