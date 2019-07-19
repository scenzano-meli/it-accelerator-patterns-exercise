package main

import (
	"./controllers"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
)

//puerto de gin
const (
	port = ":8080"
)

var (
	router = gin.Default()
)

var limiter = rate.NewLimiter(2, 3)

func limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if limiter.Allow() == false {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}


func main() {

	router.GET("/users/:userID", controllers.GetUserFromApi)
	router.GET("/sites", controllers.GetSitesListFromApi)
	router.GET("/sites/:siteID", controllers.GetSiteFromApi)
	router.GET("/countries", controllers.GetCountriesFromApi)
	router.GET("/countries/:countryId", controllers.GetCountryFromApi)
	router.GET("/user_results/:userID", controllers.GetUserResultFromApi)

	http.ListenAndServe(":8080", limit(router))

}