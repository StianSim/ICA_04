package weather

import (
  "encoding/json"
  "os"
)

// Returns the last polled response from AccuWeather
func OpenWeatherMap() OpenWeatherData {
    // The response is saved on disk, so that's where we're reading from
    f, err := os.Open("responses/openweathermap.json")
    check(err)
    defer f.Close()
    var data OpenWeatherData
    // Decode the file contents as json
    err = json.NewDecoder(f).Decode(&data)
    check(err)
    return data
}

type OpenWeatherData struct {
	Coord struct {
		Lon int `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID int `json:"id"`
		Main string `json:"main"`
		Description string `json:"description"`
		Icon string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp float64 `json:"temp"`
		Pressure float64 `json:"pressure"`
		Humidity float64 `json:"humidity"`
		TempMin float64 `json:"temp_min"`
		TempMax float64 `json:"temp_max"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg float64 `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt int `json:"dt"`
	Sys struct {
		Type int `json:"type"`
		ID int `json:"id"`
		Message float64 `json:"message"`
		Country string `json:"country"`
		Sunrise int `json:"sunrise"`
		Sunset int `json:"sunset"`
	} `json:"sys"`
	ID int `json:"id"`
	Name string `json:"name"`
	Cod int `json:"cod"`
}
