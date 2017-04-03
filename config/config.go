package config

import (
    "log"
    "github.com/BurntSushi/toml"
    "io/ioutil"
)

func GetConfig() tomlConfig {
    c, err := ioutil.ReadFile("config.toml")
    check(err)
    var config tomlConfig
    _, err = toml.Decode(string(c), &config)
    check(err)
    return config
}

type tomlConfig struct {
    Weather weatherInfo
}


type weatherInfo struct {
    Wunderground string
    OpenWeatherMap string
    AccuWeather string
    Yr string
		DarkSky string
}

// Generic error handling
func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}
