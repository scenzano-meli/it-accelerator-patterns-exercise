package domains

import (
	"../utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Site struct {
	Id 						string `json:"id"`
	Name					string `json:"name"`
	CountryId				string `json:"country_id"`

	SalesFeeMode			string `json:"sale_fees_mode"`
	MercadoPagoVersion		int `json:"mercadopago_version"`
	DefaultCurrencyId		string `json:"default_currency_id"`

	ImmediatePayment		string `json:"immediate_payment"`
	PaymentMethodsIds		[]interface{}	`json:"payment_method_ids"`

	Settings struct{
		IdentificationTypes	[]interface{} `json:"identification_types"`
		TaxPayerTypes		[]interface{} `json:"taxpayer_types"`
		//IdentificationTypesRules interface{} `json:"identification_types_rules"`
	} `json:"settings"`
}

func (site *Site) Get() *utils.ApiError{

	if site.Id == "" {
		return &utils.ApiError{
			Message: "El id de site esta vacío",
			Status: http.StatusBadRequest,
		}
	}

	url := fmt.Sprintf("%s%s", utils.UrlSites, site.Id)

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

	if err := json.Unmarshal(data, &site); err != nil{
		return &utils.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	return nil
}


