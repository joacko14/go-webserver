package geoFinder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func FindMe(long float64, lat float64) (country string, city string) {
	API_KEY := os.Getenv("GEO_KEY")
	if API_KEY == "" {
		fmt.Println("API_KEY is not set")
		return "", ""
	}
	// fetching  opencagedata api
	url := fmt.Sprintf("https://api.opencagedata.com/geocode/v1/json?q=%f+%f&key=%s", long, lat, API_KEY)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Ha ocurrdio un error", err)
		return "", ""
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Ha ocurrdio un error al leer la respueta", err)
		return "", ""
	}

	type Result struct {
		Results []struct {
			Components struct {
				Country string `json:"country"`
				City    string `json:"city"`
			} `json:"components"`
		} `json:"results"`
	}

	var result Result
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Ha ocurrdio un error al deserializar la respuesta", err)
		return "", ""
	}

	if len(result.Results) > 0 {
		country = result.Results[0].Components.Country
		city = result.Results[0].Components.City
	}

	return country, city
}
