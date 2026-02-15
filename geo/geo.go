package geo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CityPopulationResponce struct {
	Error bool `json:"error"`
}

type GeoData struct {
	City string `json:"city"`
}

func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		if !checkCity(city) {
			panic("Такого города нет")
		}
		return &GeoData{
			City: city,
		}, nil
	}
	responce, err := http.Get("https://ipapi.co/json")
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	if responce.StatusCode != 200 {
		return nil, fmt.Errorf("Ответ API: %d", responce.StatusCode)
	}
	defer responce.Body.Close()
	body, err := io.ReadAll(responce.Body)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	var geo GeoData
	json.Unmarshal(body, &geo)
	return &geo, nil
}

func checkCity(city string) bool {
	postBody, _ := json.Marshal(map[string]string{
		"city": city,
	})
	resp, err := http.Post("https://countriesnow.space/api/v0.1/countries/population/cities", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	var populationResponce CityPopulationResponce
	json.Unmarshal(body, &populationResponce)
	return !populationResponce.Error
}
