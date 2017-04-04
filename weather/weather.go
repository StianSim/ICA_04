package weather

import (
    "net/http"
    "log"
    "time"
    "io/ioutil"
    "strconv"
    "math"
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
func degreeToName(deg float64) string {
    deg = math.Mod(deg + 360, 360) // Make sure the degree is within 360
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

// Convert Kilometers per hour to meters per second
func kphToMs(kph float64) float64 {
  return kph / 3.6
}

// Convert Miles per hour to meters per second
func mphToMs(mph float64) float64 {
    return mph * 0.44704
}

// Return the average value of
// a slice of float64 arguments.
func average(u ...float64) float64 {
  var total float64
  for _, i := range u {
    total += i // For each of the elements in u, add the value to the total
  }
  return total / float64(len(u)) // Then return the average of all the values
}

// Convert Fahrenheit to Celsius
func fahrenheitToCelsius(f float64) float64 {
	return (f - 32.0) / 1.8
}

// A helper function to parse a string to float,
// which raises a fatal error if it encounters
// an error.
func stringToFloat(s string) float64 {
    f, err := strconv.ParseFloat(s, 64)
    check(err)
    return f
}

// Returns a "unified" struct with consolidated data
// As well as the individual sources' data in their respective
// fields.
func GetWeather() Weather {
    // Get the latest polled data from each of the APIs
    acw := AccuWeather()
    owm := OpenWeatherMap()
    wun := Wunderground()
    yr := Yr()
    dsk := DarkSky()

    return Weather {
        Location: location {
            Name: owm.Name, // Name of the location
            Lat: owm.Coord.Lat, // Latitude
            Lon: owm.Coord.Lon, // Longitude
        },
        // Conditions as per Yr.no's descriptions, chosen purely because it is in Norwegian
        Conditions: yr.Weatherdata.Forecast.Tabular.Time[0].Symbol.Name,
        // Get the average temperature from all of the API's latest readings
        Temperature: average(
            acw[0].Temperature.Metric.Value,
            owm.Main.Temp,
            wun.CurrentObservation.TempC,
            dsk.Currently.TemperatureC,
            stringToFloat(yr.Weatherdata.Observations.Weatherstation[0].Temperature.Value),
        ),
        // Get the average wind speed in meters per second from the APIs that carry
        // that information
        WindSpeed: average(
            owm.Wind.Speed,
            wun.CurrentObservation.WindMs,
            stringToFloat(yr.Weatherdata.Observations.Weatherstation[0].WindSpeed.Mps),
            dsk.Currently.WindSpeedMs,
        ),

        // Average the wind direction in degrees and get the Norwegian name for the direction
        WindDirection: degreeToName(average(
            owm.Wind.Deg,
            wun.CurrentObservation.WindDegrees,
            stringToFloat(yr.Weatherdata.Observations.Weatherstation[0].WindDirection.Deg),
            dsk.Currently.WindBearing,
        )),
        // Store the individual data
        AccuWeather: acw,
        OpenWeatherMap: owm,
        Wunderground: wun,
        Yr: yr,
        DarkSky: dsk,
    }

}

type Weather struct {
    Location location
    Conditions string
    Temperature float64
    WindSpeed float64
    WindDirection string
    AccuWeather AccuWeatherData
    OpenWeatherMap OpenWeatherData
    Wunderground WundergroundData
    Yr YrData
    DarkSky DarkSkyData
}

type location struct {
    Name string
    Lat float64
    Lon float64
}
