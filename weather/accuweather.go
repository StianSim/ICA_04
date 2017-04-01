package weather

import (
    "encoding/json"
    "time"
    "os"
)

// Returns the last polled response from AccuWeather
func AccuWeatherResp() AccuWeatherData {
    // The response is saved on disk, so that's where we're reading from
    f, err := os.Open("responses/accuweather.json")
    check(err)
    defer f.Close()
    var data AccuWeatherData
    // Decode the file contents as json
    err = json.NewDecoder(f).Decode(&data)
    check(err)
    return data
}

// Automatically generated struct
// https://mholt.github.io/json-to-go/
type AccuWeatherData []struct {
    LocalObservationDateTime time.Time `json:"LocalObservationDateTime"`
    EpochTime int `json:"EpochTime"`
    WeatherText string `json:"WeatherText"`
    WeatherIcon int `json:"WeatherIcon"`
    IsDayTime bool `json:"IsDayTime"`
    Temperature struct {
        Metric struct {
            Value float64 `json:"Value"`
            Unit string `json:"Unit"`
            UnitType int `json:"UnitType"`
        } `json:"Metric"`
        Imperial struct {
            Value float64 `json:"Value"`
            Unit string `json:"Unit"`
            UnitType int `json:"UnitType"`
        } `json:"Imperial"`
    } `json:"Temperature"`
    MobileLink string `json:"MobileLink"`
    Link string `json:"Link"`
}
