package weather

import (
    "testing"
    "../config"
    "net/http"
    "encoding/json"
)

// Returns an API response
func dskResponse(url string) DarkSkyData {
    resp, err := http.Get(url)
    check(err)
    defer resp.Body.Close()
    var data DarkSkyData
    err = json.NewDecoder(resp.Body).Decode(&data)
    return data
}

func TestGoodDskAPIResponse(t *testing.T) {
    r := dskResponse(config.GetConfig("../config.toml").Weather.DarkSky)
    if r.Currently.Time == 0 {
        // r will be initialized, and r.Currently.Time will never be 0 unless something
        // weird has happened.
        t.Errorf("DarkSky.Currently.Time returned %d, expected something else", r.Currently.Time)
    }
}

func TestBadDskAPIResponse(t *testing.T) {
    r := dskResponse(config.GetConfig("../config.toml").Weather.DarkSky + "a")
    if r.Currently.Time != 0 {
        // Since we invoked something weird to happen, r.Currently.Time will be 0 here.
        t.Errorf("DarkSky.Currently.Time returned %d, expected 0", r.Currently.Time)
    }
}
