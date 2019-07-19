package domains

import (
	"../utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Country struct {
	Id					string `json:"id"`
	Name				string `json:"name"`
	Locale				string `json:"locale"`
	Currency			string `json:"currency_id"`
	DecimalSeparator	string `json:"decimal_separator"`
	ThousandsSeparator	string `json:"decimal_separator"`
	TimeZone			interface{} `json:"time_zone"`
	GeoInformation struct {
		Location struct {
			Latitude interface{} `json:"latitude"`
			Longitude interface{} `json:"longitude"`
		} `json:"location"`
	} `json:"geo_information"`
	States []interface{} `json:"states"`
}

func (country *Country) Get() *utils.ApiError{

	if country.Id == "" {
		return &utils.ApiError{
			Message: "El id de country esta vacío",
			Status: http.StatusBadRequest,
		}
	}

	url := fmt.Sprintf("%s%s", utils.UrlContries, country.Id)

	response, err := http.Get(url)
	if err != nil{
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal(data, &country); err != nil{
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	return nil
}

