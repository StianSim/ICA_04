package weather

import (
    "encoding/json"
    "time"
    "os"
)


func Yr() YrData {
    f, err := os.Open("responses/yr.json")
    check(err)
    defer f.Close()
    var data YrData
    err = json.NewDecoder(f).Decode(&data)
    check(err)
    return data
}

type YrData struct {
    Weatherdata struct {
        Credit struct {
            Link struct {
                Text string `json:"@text"`
                URL string `json:"@url"`
            } `json:"link"`
        } `json:"credit"`
        Forecast struct {
            Tabular struct {
                Time []struct {
                    From string `json:"@from"`
                    Period string `json:"@period"`
                    To string `json:"@to"`
                    Precipitation struct {
                        Value string `json:"@value"`
                    } `json:"precipitation"`
                    Pressure struct {
                        Unit string `json:"@unit"`
                        Value string `json:"@value"`
                    } `json:"pressure"`
                    Symbol struct {
                        Name string `json:"@name"`
                        Number string `json:"@number"`
                        NumberEx string `json:"@numberEx"`
                        Var string `json:"@var"`
                    } `json:"symbol"`
                    Temperature struct {
                        Unit string `json:"@unit"`
                        Value string `json:"@value"`
                    } `json:"temperature"`
                    WindDirection struct {
                        Code string `json:"@code"`
                        Deg string `json:"@deg"`
                        Name string `json:"@name"`
                    } `json:"windDirection"`
                    WindSpeed struct {
                        Mps string `json:"@mps"`
                        Name string `json:"@name"`
                    } `json:"windSpeed"`
                } `json:"time"`
            } `json:"tabular"`
            Text struct {
                Location struct {
                    Name string `json:"@name"`
                    Time []struct {
                        From string `json:"@from"`
                        To string `json:"@to"`
                        Body struct {
                            Text string `json:"#text"`
                            Strong []string `json:"strong"`
                        } `json:"body"`
                        Title string `json:"title"`
                    } `json:"time"`
                } `json:"location"`
            } `json:"text"`
        } `json:"forecast"`
        Links struct {
            Link []struct {
                ID string `json:"@id"`
                URL string `json:"@url"`
            } `json:"link"`
        } `json:"links"`
        Location struct {
            Country string `json:"country"`
            Location struct {
                Altitude string `json:"@altitude"`
                Geobase string `json:"@geobase"`
                Geobaseid string `json:"@geobaseid"`
                Latitude string `json:"@latitude"`
                Longitude string `json:"@longitude"`
            } `json:"location"`
            Name string `json:"name"`
            Timezone struct {
                ID string `json:"@id"`
                UtcoffsetMinutes string `json:"@utcoffsetMinutes"`
            } `json:"timezone"`
            Type string `json:"type"`
        } `json:"location"`
        Meta struct {
            Lastupdate string `json:"lastupdate"`
            Nextupdate string `json:"nextupdate"`
        } `json:"meta"`
        Observations struct {
            Weatherstation []struct {
                Distance string `json:"@distance"`
                Lat string `json:"@lat"`
                Lon string `json:"@lon"`
                Name string `json:"@name"`
                Source string `json:"@source"`
                Stno string `json:"@stno"`
                Sttype string `json:"@sttype"`
                Temperature struct {
                    Time time.Time `json:"@time"`
                    Unit string `json:"@unit"`
                    Value string `json:"@value"`
                } `json:"temperature"`
                WindDirection struct {
                    Code string `json:"@code"`
                    Deg string `json:"@deg"`
                    Name string `json:"@name"`
                    Time time.Time `json:"@time"`
                } `json:"windDirection,omitempty"`
                WindSpeed struct {
                    Mps string `json:"@mps"`
                    Name string `json:"@name"`
                    Time time.Time `json:"@time"`
                } `json:"windSpeed,omitempty"`
            } `json:"weatherstation"`
        } `json:"observations"`
        Sun struct {
            Rise string `json:"@rise"`
            Set string `json:"@set"`
        } `json:"sun"`
    } `json:"weatherdata"`
}
