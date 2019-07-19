package services

import (
	"../domains"
	"../utils"
	"fmt"
	"sync"
)

func getCountryRoutine(user domains.User, c chan *domains.Result) {

	var result domains.Result

	fmt.Println("Dentro del country")
	country := &domains.Country {
		Id: user.CountryID,
	}

	if err := country.Get(); err != nil {
		fmt.Println(err)
	}

	result.Country = country

	c <- &result
}

func getSiteRoutine(user domains.User, c chan *domains.Result) {

	var result domains.Result

	fmt.Println("Dentro del site")
	site := &domains.Site{
		Id: user.SiteID,
	}

	if err := site.Get(); err != nil {
		fmt.Println(err)
	}

	result.Site = site

	c <- &result
}


func GetResult(userId int64) (*domains.Result, *utils.ApiError) {

	var wg sync.WaitGroup
	ch := make(chan *domains.Result)
	defer close(ch)

	user := &domains.User{
		Id: userId,
	}

	if err := user.Get(); err != nil {
		return nil, err
	}

	result := &domains.Result{
		User: user,
	}

	wg.Add(1)

	go getCountryRoutine(*user, ch)

	wg.Add(1)

	go getSiteRoutine(*user, ch)

	go func() {
		for resp := range ch {
			wg.Done()
			if resp.Site != nil  {
				result.Site = resp.Site
			}

			if resp.Country != nil {
				result.Country = resp.Country
			}
		}
	}()

	wg.Wait()

	return  result, nil
}
