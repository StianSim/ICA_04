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
    server.Use(martini.Static("assets", martini.StaticOptions{Prefix: "/assets"}))
    c := config.GetConfig()

    // Start polling each of the APIs for responses in another thread.
    // The arguments passed to the loops are
    // url (from config.toml) to get a response from,
    // filename to store the response at, and timeout in minutes
    go weather.WeatherLoop(c.Weather.AccuWeather, "accuweather.json", 20)
    go weather.WeatherLoop(c.Weather.OpenWeatherMap, "openweathermap.json", 5)
    go weather.WeatherLoop(c.Weather.Wunderground, "wunderground.json", 15)
    go weather.WeatherLoop(c.Weather.Yr, "yr.json", 20)
    go weather.WeatherLoop(c.Weather.DarkSky, "darksky.json", 10)

    // Map / (the default endpoint when accessing a server) to the following response
    server.Get("/", func(r render.Render, args martini.Params) {

        // Get the consolidated data constructed from the latest data.
        p := weather.GetWeather()

        // Render the contents of "templates/index.tmpl", and pass p to it.
        r.HTML(200, "index", p)
    })
    // Finally, tell the server to enter its main loop with the previous configuration
    server.Run()
}
