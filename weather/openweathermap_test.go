package weather

import (
    "testing"
    "../config"
    "net/http"
    "encoding/json"
)

// Returns an API response
func owmResponse(url string) OpenWeatherData {
    resp, err := http.Get(url)
    check(err)
    defer resp.Body.Close()
    var data OpenWeatherData
    err = json.NewDecoder(resp.Body).Decode(&data)
    return data
}

func TestGoodOwmAPIResponse(t *testing.T) {
    r := owmResponse(config.GetConfig("../config.toml").Weather.OpenWeatherMap)
    if r.Cod != 200 {
        // Regardless of bad or good responses, OpenWeatherMap always sets a value in
        // "cod". Therefore, we check if it is not 200 (HTTP Status OK).
        t.Errorf("OpenWeatherMap returned status code %d, expected 200", r.Cod)
    }
}

func TestBadOwmAPIResponse(t *testing.T) {
    r := owmResponse(config.GetConfig("../config.toml").Weather.OpenWeatherMap + "a")
    if r.Cod == 200 {
        // Since we invoked something weird to happen, r.Cod will never be 200.
        t.Errorf("OpenWeatherMap returned status code 200, expected something else")
    }
}
