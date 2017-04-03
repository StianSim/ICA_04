# is105-ica05

En konsolideringsside for været i Kristiansand. Ved gjennomsnittsverdier av data fra flere kilder siktes det mot mer "pålitelig" informasjon.

Sidene som blir spurt etter informasjon er [Yr.no](http://yr.no/), [OpenWeatherMap](http://openweathermap.org/), [Wunderground](https://www.wunderground.com/), [AccuWeather](http://www.accuweather.com/) og [DarkSky](https://darksky.net/app/).

Henting av informasjon er satt opp som goroutines i `server.go`, da for både å hedre API-restriksjoner, samt å vise et klart eksempel for concurrency.

## Installasjonsinstruksjoner

* Lag en mappe kalt `responses`. De nedlastede json-filene vil bli lagret her for senere bruk
* Kopier `config_example.toml` til `config.toml`
* I `config.toml`, sett inn API nøkler i stedet for `<your api key>`

## Avhengigheter

Bruk `go get` for å installere biblioteker (forutsetter at `$GOPATH` er satt).

* `github.com/BurntSushi/toml`
* `github.com/go-martini/martini`
* `github.com/martini-contrib/render`

## Ekstra informasjon

Yr.no's API er kun XML, men ved hjelp av et lite script i `Python` har vi konvertert det til
JSON, dog ikke for offentlig bruk da Yr.no's retningslinjer sier at data må mellomlagres i minst 10 minutter.
På grunn av dette, mellomlagres dataen vi bruker fra Yr.no i 20 minutter før fornyelse.
