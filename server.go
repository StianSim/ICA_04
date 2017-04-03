package main

import (
    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"
    "./weather"
    "./config"
)


func main() {
    // Create an instance of Martini in the server variable
    server := martini.Classic()
    // Tell the server to use an instance of the renderer from the render package
    server.Use(render.Renderer())

    c := config.GetConfig()
    // Start polling AccuWeather for responses in another thread
    go weather.WeatherLoop(c.Weather.AccuWeather, "accuweather.json", 20)
    go weather.WeatherLoop(c.Weather.OpenWeatherMap, "openweathermap.json", 5)
    go weather.WeatherLoop(c.Weather.Wunderground, "wunderground.json", 15)
    go weather.WeatherLoop(c.Weather.Yr, "yr.json", 20)
		go weather.WeatherLoop(c.Weather.DarkSky, "darksky.json", 10)
    // Map / (the default endpoint when accessing a server) to the following response
    server.Get("/", func(r render.Render, args martini.Params) {

        // Get the latest polled response from AccuWeather and store it in p.
        // Since the AccuWeatherData type is a list of structs, and we only want the first element,
        // [0] is used to get that element at the first index.
        p := weather.GetWeather()

        // Render the contents of "templates/index.tmpl", and pass p to it.
        r.HTML(200, "index", p)
    })
    // Finally, tell the server to enter its main loop with the previous configuration
    server.Run()
}
