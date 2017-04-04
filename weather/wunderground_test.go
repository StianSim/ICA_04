package weather

import (
    "testing"
    "../config"
    "net/http"
    "encoding/json"
)

// Returns an API response
func wunResponse(url string) WundergroundData {
    resp, err := http.Get(url)
    check(err)
    defer resp.Body.Close()
    var data WundergroundData
    err = json.NewDecoder(resp.Body).Decode(&data)
    // We choose not to handle the error here.
    return data
}

func TestGoodWunAPIResponse(t *testing.T) {
    r := wunResponse(config.GetConfig("../config.toml").Weather.Wunderground)
    if r.Response.Error.Type != "" {
        // Wunderground's Error fields will not be set unless there is an error.
        t.Errorf("Wunderground returned error %s", r.Response.Error.Type)
    }
}

func TestBadWunAPIResponse(t *testing.T) {
    r := wunResponse(config.GetConfig("../config.toml").Weather.Wunderground + "a")
    if r.Response.Error.Type == "" {
        // Since we invoked something weird to happen, r.Error.Type will never be " ".
        t.Errorf("Wunderground returned no error")
    }
}
