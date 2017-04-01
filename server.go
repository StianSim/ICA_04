package main

import (
    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"
    "./weather"
)


func main() {
    // Create an instance of Martini in the server variable
    server := martini.Classic()
    // Tell the server to use an instance of the renderer from the render package
    server.Use(render.Renderer())

    // Start polling AccuWeather for responses in another thread
    go weather.AccuWeatherLoop()

    // Map / (the default endpoint when accessing a server) to the following response
    server.Get("/", func(r render.Render, args martini.Params) {

        // Get the latest polled response from AccuWeather and store it in p.
        // Since the AccuWeatherData type is a list of structs, and we only want the first element,
        // [0] is used to get that element at the first index.
        p := weather.AccuWeatherResp()[0]

        // Render the contents of "templates/index.tmpl", and pass p to it.
        r.HTML(200, "index", p)
    })
    // Finally, tell the server to enter its main loop with the previous configuration
    server.Run()
}
