package weather

import (
    "net/http"
    "encoding/json"
    "time"
    "io/ioutil"
    "os"
    "log"
    "../config"
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

// A loop that regularly polls AccuWeather for weather information
func AccuWeatherLoop() {
    for {
        accuWeather()
        log.Println("Updated AccuWeather")
        // Sleep for 20 minutes to safely fit inside AccuWeather's API rate limits
        time.Sleep(20 * time.Minute)
    }
}

func accuWeather() {
    c := config.GetConfig()
    resp, err := http.Get(c.Weather.AccuWeather)
    check(err)
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    err = ioutil.WriteFile("responses/accuweather.json", body, 0644)
    check(err)
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
