package weather

import (
    "encoding/json"
    "time"
    "os"
)

// Return the latest data from Yr.
func Yr() YrData {
    f, err := os.Open("responses/yr.json")
    check(err)
    defer f.Close()
    var data YrData
    err = json.NewDecoder(f).Decode(&data)
    check(err)
    return data
}

// Automatically generated struct
// https://mholt.github.io/json-to-go/
type YrData struct {
    Weatherdata struct {
        Credit struct {
            Link struct {
                Text string `json:"text"`
                URL string `json:"url"`
            } `json:"link"`
        } `json:"credit"`
        Forecast struct {
            Tabular struct {
                Time []struct {
                    From string `json:"from"`
                    Period int `json:"period"`
                    Precipitation struct {
                        Value int `json:"value"`
                    } `json:"precipitation"`
                    Pressure struct {
                        Unit string `json:"unit"`
                        Value float64 `json:"value"`
                    } `json:"pressure"`
                    Symbol struct {
                        Name string `json:"name"`
                        Number int `json:"number"`
                        NumberEx int `json:"numberEx"`
                    } `json:"symbol"`
                    Temperature struct {
                        Unit string `json:"unit"`
                        Value int `json:"value"`
                    } `json:"temperature"`
                    To string `json:"to"`
                    WindDirection struct {
                        Code string `json:"code"`
                        Deg float64 `json:"deg"`
                        Name string `json:"name"`
                    } `json:"windDirection"`
                    WindSpeed struct {
                        Mps float64 `json:"mps"`
                        Name string `json:"name"`
                    } `json:"windSpeed"`
                } `json:"time"`
            } `json:"tabular"`
            Text struct {
                Location struct {
                    Name string `json:"name"`
                    Time []struct {
                        Body struct {
                            Text string `json:"#text"`
                        } `json:"body"`
                        From string `json:"from"`
                        Title string `json:"title"`
                        To string `json:"to"`
                    } `json:"time"`
                } `json:"location"`
            } `json:"text"`
        } `json:"forecast"`
        Links struct {
            Link []struct {
                ID string `json:"id"`
                URL string `json:"url"`
            } `json:"link"`
        } `json:"links"`
        Location struct {
            Country string `json:"country"`
            Location struct {
                Altitude int `json:"altitude"`
                Geobase string `json:"geobase"`
                Geobaseid int `json:"geobaseid"`
                Latitude float64 `json:"latitude"`
                Longitude float64 `json:"longitude"`
            } `json:"location"`
            Name string `json:"name"`
            Timezone struct {
                ID string `json:"id"`
                UtcoffsetMinutes int `json:"utcoffsetMinutes"`
            } `json:"timezone"`
            Type string `json:"type"`
        } `json:"location"`
        Meta struct {
            Lastupdate string `json:"lastupdate"`
            Nextupdate string `json:"nextupdate"`
        } `json:"meta"`
        Observations struct {
            Weatherstation []struct {
                Distance int `json:"distance"`
                Lat float64 `json:"lat"`
                Lon float64 `json:"lon"`
                Name string `json:"name"`
                Source string `json:"source"`
                Stno int `json:"stno"`
                Sttype string `json:"sttype"`
                Temperature struct {
                    Time time.Time `json:"time"`
                    Unit string `json:"unit"`
                    Value float64 `json:"value"`
                } `json:"temperature"`
                WindDirection struct {
                    Code string `json:"code"`
                    Deg float64 `json:"deg"`
                    Name string `json:"name"`
                    Time time.Time `json:"time"`
                } `json:"windDirection,omitempty"`
                WindSpeed struct {
                    Mps float64 `json:"mps"`
                    Name string `json:"name"`
                    Time time.Time `json:"time"`
                } `json:"windSpeed,omitempty"`
            } `json:"weatherstation"`
        } `json:"observations"`
        Sun struct {
            Rise string `json:"rise"`
            Set string `json:"set"`
        } `json:"sun"`
    } `json:"weatherdata"`
}
