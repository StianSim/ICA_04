package weather

import (
    "net/http"
    "log"
    "time"
    "io/ioutil"
)

// A generic check for error handling
func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

// Meant to be used in goroutines to periodically poll for json responses
func WeatherLoop(url string, filename string, timeout time.Duration) {
    for {
        getResponse(url, filename)
        log.Println("Updated " + filename)
        time.Sleep(timeout * time.Minute)
    }
}

// Downloads a json response and saves it in responses/filename
func getResponse(url string, filename string) {
    resp, err := http.Get(url)
    check(err)
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    err = ioutil.WriteFile("responses/" + filename, body, 0644)
    check(err)
}

// Takes a degree and returns the name for a wind direction
// Based on http://climate.umn.edu/snow_fence/components/winddirectionanddegreeswithouttable3.htm
func DegreeToName(deg float64) string {
    if deg >= 11.25 && deg < 33.75 {
        return "Nord-nordøst"
    } else if deg >= 33.75 && deg < 56.25 {
        return "Nordøst"
    } else if deg >= 56.25 && deg < 78.75 {
        return "Øst-nordøst"
    } else if deg >= 78.75 && deg < 101.25 {
        return "Øst"
    } else if deg >= 101.25 && deg < 123.75 {
        return "Øst-sørøst"
    } else if deg >= 123.75 && deg < 146.25 {
        return "Sørøst"
    } else if deg >= 146.25 && deg < 168.75 {
        return "Sør-sørøst"
    } else if deg >= 168.75 && deg < 191.25 {
        return "Sør"
    } else if deg >= 191.25 && deg < 213.75 {
        return "Sør-sørvest"
    } else if deg >= 213.75 && deg < 236.25 {
        return "Sørvest"
    } else if deg >= 236.25 && deg < 258.75 {
        return "Vest-sørvest"
    } else if deg >= 258.75 && deg < 281.25 {
        return "Vest"
    } else if deg >= 281.25 && deg < 303.75 {
        return "Vest-nordvest"
    } else if deg >= 303.75 && deg < 326.25 {
        return "Nordvest"
    } else if deg >= 326.25 && deg < 348.75 {
        return "Nord-nordvest"
    } else {
        return "Nord"
    }
}

func KphToMs(kph float64) float64 {
  return kph / 3.6
}

func Average(u ...float64) float64 {
  var total float64
  for _, i := range u {
    total += i
  }
  return total / float64(len(u))
}

func FahrenheitToCelcius(f float64) float64 {
	return (f - 32.0) / 1.8
}

// Returns a "unified" struct with consolidated data
// As well as the individual sources' data in their respective
// fields.
func GetWeather() Weather {
    acw := AccuWeather()
    owm := OpenWeatherMap()
    wun := Wunderground()
    yr := Yr()
		dsk := DarkSky()
    name := owm.Name
    lat := owm.Coord.Lat
    lon := owm.Coord.Lon
    temp := Average(acw[0].Temperature.Metric.Value, owm.Main.Temp, wun.CurrentObservation.TempC, FahrenheitToCelcius(dsk.Currently.Temperature))
    windspeed := Average(owm.Wind.Speed, KphToMs(wun.CurrentObservation.WindKph))
    winddirection := DegreeToName(Average(owm.Wind.Deg, wun.CurrentObservation.WindDegrees))
    return Weather {
        Location: location {
            Name: name,
            Lat: lat,
            Lon: lon,
        },
        Temperature: temp,
        WindSpeed: windspeed,
        WindDirection: winddirection,
        AccuWeather: acw,
        OpenWeatherMap: owm,
        Wunderground: wun,
        Yr: yr,
    }

}

type Weather struct {
    Location location
    Temperature float64
    WindSpeed float64
    WindDirection string
    AccuWeather AccuWeatherData
    OpenWeatherMap OpenWeatherData
    Wunderground WundergroundData
    Yr YrData
}

type location struct {
    Name string
    Lat float64
    Lon float64
}
