package weather

import (
    "encoding/json"
    "os"
)

// Returns the last polled response from DarkSky
func DarkSky() DarkSkyData {
    // The response is saved on disk, so that's where we're reading from
    f, err := os.Open("responses/darksky.json")
    check(err)
    defer f.Close()
    var data DarkSkyData
    // Decode the file contents as json
    err = json.NewDecoder(f).Decode(&data)
    check(err)
    data.Currently.WindSpeedMs = mphToMs(data.Currently.WindSpeed)
    return data
}

// Automatically generated struct
// https://mholt.github.io/json-to-go/
type DarkSkyData struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone string `json:"timezone"`
	Offset int `json:"offset"`
	Currently struct {
		Time int `json:"time"`
		Summary string `json:"summary"`
		Icon string `json:"icon"`
		PrecipIntensity float64 `json:"precipIntensity"`
		PrecipProbability float64 `json:"precipProbability"`
		Temperature float64 `json:"temperature"`
		ApparentTemperature float64 `json:"apparentTemperature"`
		DewPoint float64 `json:"dewPoint"`
		Humidity float64 `json:"humidity"`
		WindSpeed float64 `json:"windSpeed"`
        WindSpeedMs float64 // Field not in JSON, intended to be assigned after parsing
		WindBearing float64 `json:"windBearing"`
		Visibility float64 `json:"visibility"`
		CloudCover float64 `json:"cloudCover"`
		Pressure float64 `json:"pressure"`
		Ozone float64 `json:"ozone"`
	} `json:"currently"`
	Hourly struct {
		Summary string `json:"summary"`
		Icon string `json:"icon"`
		Data []struct {
			Time int `json:"time"`
			Summary string `json:"summary"`
			Icon string `json:"icon"`
			PrecipIntensity float64 `json:"precipIntensity"`
			PrecipProbability float64 `json:"precipProbability"`
			Temperature float64 `json:"temperature"`
			ApparentTemperature float64 `json:"apparentTemperature"`
			DewPoint float64 `json:"dewPoint"`
			Humidity float64 `json:"humidity"`
			WindSpeed float64 `json:"windSpeed"`
			WindBearing int `json:"windBearing"`
			Visibility float64 `json:"visibility,omitempty"`
			CloudCover float64 `json:"cloudCover"`
			Pressure float64 `json:"pressure"`
			Ozone float64 `json:"ozone"`
			PrecipType string `json:"precipType,omitempty"`
		} `json:"data"`
	} `json:"hourly"`
	Daily struct {
		Summary string `json:"summary"`
		Icon string `json:"icon"`
		Data []struct {
			Time int `json:"time"`
			Summary string `json:"summary"`
			Icon string `json:"icon"`
			SunriseTime int `json:"sunriseTime"`
			SunsetTime int `json:"sunsetTime"`
			MoonPhase float64 `json:"moonPhase"`
			PrecipIntensity float64 `json:"precipIntensity"`
			PrecipIntensityMax float64 `json:"precipIntensityMax"`
			PrecipProbability float64 `json:"precipProbability"`
			TemperatureMin float64 `json:"temperatureMin"`
			TemperatureMinTime int `json:"temperatureMinTime"`
			TemperatureMax float64 `json:"temperatureMax"`
			TemperatureMaxTime int `json:"temperatureMaxTime"`
			ApparentTemperatureMin float64 `json:"apparentTemperatureMin"`
			ApparentTemperatureMinTime int `json:"apparentTemperatureMinTime"`
			ApparentTemperatureMax float64 `json:"apparentTemperatureMax"`
			ApparentTemperatureMaxTime int `json:"apparentTemperatureMaxTime"`
			DewPoint float64 `json:"dewPoint"`
			Humidity float64 `json:"humidity"`
			WindSpeed float64 `json:"windSpeed"`
			WindBearing float64 `json:"windBearing"`
			Visibility float64 `json:"visibility,omitempty"`
			CloudCover float64 `json:"cloudCover"`
			Pressure float64 `json:"pressure"`
			Ozone float64 `json:"ozone"`
			PrecipIntensityMaxTime int `json:"precipIntensityMaxTime,omitempty"`
			PrecipType string `json:"precipType,omitempty"`
		} `json:"data"`
	} `json:"daily"`
	Flags struct {
		Sources []string `json:"sources"`
		IsdStations []string `json:"isd-stations"`
		MadisStations []string `json:"madis-stations"`
		Units string `json:"units"`
	} `json:"flags"`
}
