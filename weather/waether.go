package weather

import (
	"demo/weather/geo"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetWeather(geo geo.GeoData, format int) (string, error) {
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}
	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()
	resp, err := http.Get(baseUrl.String())
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}
	return string(body), nil
}
