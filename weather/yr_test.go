package weather

import (
    "testing"
    "../config"
    "net/http"
    "encoding/json"
)

// Returns an API response
func yrResponse(url string) YrData {
    resp, err := http.Get(url)
    check(err)
    defer resp.Body.Close()
    var data YrData
    err = json.NewDecoder(resp.Body).Decode(&data)
    return data
}

// Unfortunately, our Yr endpoint is not a true API in the
// sense that it's just a static URL that takes no parameters.
// So we can only check if that endpoint has returned something wrong.
func TestYrAPIResponse(t *testing.T) {
    r := yrResponse(config.GetConfig("../config.toml").Weather.Yr)
    if r.Weatherdata.Location.Country != "Noreg" {
        // We just check if a field has the wrong value, but
        // in reality, no fields will have been initialized anyway.
        t.Errorf("Yr.Weatherdata.Location.Country returned %s, expected Noreg", r.Weatherdata.Location.Country)
    }
}
