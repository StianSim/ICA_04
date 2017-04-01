package weather

import (
    "log"
)

// A generic check for error handling
func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
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
