package main


import (
    "github.com/ant0ine/go-json-rest/rest"
)


// CORS
type Country struct {
	Code string
	Name string
}
func GetAllCountries(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(
		[]Country{
			Country{
				Code: "FR",
				Name: "France",
			},
			Country{
				Code: "US",
				Name: "United States",
			},
		},
	)
}// CORS /
