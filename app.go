package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const API_KEY = "365e0a9bdfce4a1ca3e47fc45fb8eb4e"

func main() {
	// gorillaserver.Start()
	fmt.Println(findMe(57.461459, -4.221189))
}

func findMe(long float64, lat float64) (country string, city string) {
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
