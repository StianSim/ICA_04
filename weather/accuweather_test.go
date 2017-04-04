package weather

import (
    "testing"
    "../config"
    "net/http"
    "encoding/json"
)

// Returns an API response
func acwResponse(url string) AccuWeatherData {
    resp, err := http.Get(url)
    check(err)
    defer resp.Body.Close()
    var data AccuWeatherData
    err = json.NewDecoder(resp.Body).Decode(&data)
    return data
}

func TestGoodAcwAPIResponse(t *testing.T) {
    r := acwResponse(config.GetConfig("../config.toml").Weather.AccuWeather)
    if len(r) != 1 {
        // Since the AccuWeather response is a slice of structs, we check if our call
        // returns a single element in that slice.
        t.Errorf("AccuWeather length returned %d, expected 1", len(r))
    }
}

func TestBadAcwAPIResponse(t *testing.T) {
    // Should return an unitialized slice of structs
    r := acwResponse(config.GetConfig("../config.toml").Weather.AccuWeather + "a")
    if len(r) != 0 {
        t.Errorf("AccuWeather returned %d, expected 0", len(r))
    }
}
