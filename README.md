# [is105-ica05](https://darn.site/)

En konsolideringsside for været i Kristiansand. Ved gjennomsnittsverdier av data fra flere kilder siktes det mot mer "pålitelig" informasjon.
Informasjonen vi får fra gjennomsnittsverdiene er: Temperatur, vindretning og vindstyrke.
Værtypen har vi valgt å ta informasjon fra er Yr.no, fordi de gir oss typen på norsk.
Temperatur er målt i °C, siden det er Temperaturverdien vi benytter i Norge.
Vindstyrke er målt i m/s, som er den mest vanlige verdien for måling av vind.
Vindretningsverdiene er hentet fra [climate.umn.edu](http://climate.umn.edu/snow_fence/components/winddirectionanddegreeswithouttable3.htm)
og forteller i hvilken retning det blåser fra.

Hvis vi holder musen over de forskjellige verdiene, vil vi få informasjonen fra hver enkelt side vi bruker informasjon fra.
Kartet, som er tatt fra Google Maps, er satt til å vise Kristiansand og omegn.
Under kartet finner vi linker til de forskjellige værtjenestene, samt en link til
vårt GitHub repository.

Siden er (stort sett) å finne her: https://darn.site/

Sidene som blir spurt etter informasjon er [Yr.no](http://yr.no/), [OpenWeatherMap](http://openweathermap.org/), [Wunderground](https://www.wunderground.com/), [AccuWeather](http://www.accuweather.com/) og [DarkSky](https://darksky.net/app/).

Henting av informasjon er satt opp som goroutines i `server.go`, da for både å hedre API-restriksjoner, samt å vise et klart eksempel for concurrency.

I bunnen av dette dokumentet finnes instruksjoner for å kunne sette opp dette systemet selv.

## Systemarkitektur

Konsolideringsserveren henter periodevis data fra forskjellige API-er og mellomlagrer informasjonen. Når en bruker vil se på siden, er det data som sist ble hentet som vises.
Dermed blir det ikke unødvendig mange kall til API-ene bak om det er mange brukere som vil se på den konsoliderte siden.
Her er en illustrasjon som viser logikken, både henting og sending av data blir sendt over HTTP eller HTTPS.

![Systemarkitektur-sketch](https://raw.githubusercontent.com/crippling-depression/is105-ica05/master/assets/system-arch-sketch.png)

## Brukerhistorier

Brukerhistorier finnes i [Brukerscenario.md](https://github.com/crippling-depression/is105-ica05/blob/master/Brukerscenario.md)

## Ekstra informasjon

Yr.no's API er kun XML, men ved hjelp av et lite script i `Python` har vi konvertert det til
JSON, dog ikke for offentlig bruk da Yr.no's retningslinjer sier at data må mellomlagres i minst 10 minutter.
På grunn av dette, mellomlagres dataen vi bruker fra Yr.no i 20 minutter før fornyelse.

## Installasjonsinstruksjoner

* Lag en mappe kalt `responses`. De nedlastede json-filene vil bli lagret her for senere bruk
* Kopier `config_example.toml` til `config.toml`
* I `config.toml`, sett inn API nøkler i stedet for `<your api key>`

## Avhengigheter

Bruk `go get` for å installere biblioteker (forutsetter at `$GOPATH` er satt).

* `github.com/BurntSushi/toml`
* `github.com/go-martini/martini`
* `github.com/martini-contrib/render`
