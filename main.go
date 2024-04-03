package main

import (
	"flag"
	"fmt"
)

var transportationMethods = map[string]float64{
	"small-diesel-car":         142,
	"small-petrol-car":         154,
	"small-plugin-hybrid-car":  73,
	"small-electric-car":       50,
	"medium-diesel-car":        171,
	"medium-petrol-car":        192,
	"medium-plugin-hybrid-car": 110,
	"medium-electric-car":      58,
	"large-diesel-car":         209,
	"large-petrol-car":         282,
	"large-plugin-hybrid-car":  126,
	"large-electric-car":       73,
	"bus":                      27,
	"train":                    6,
}

func main() {
	var (
		transportationMethod = flag.String("transportation-method", "", "Transporation Method")
		distance             = flag.Float64("distance", 0, "Distance")
		distanceUnit         = flag.String("unit-of-distance", "km", "Defines the unit of the distance (km or m)")
		outputUnit           = flag.String("output", "g", "Output unit (g or kg)")
	)

	flag.Parse()

	// Get emission rate of transportation method
	emission := transportationMethods[*transportationMethod]

	if *distanceUnit == "m" {
		*distance /= 1000
	}

	// calculate co2
	co2 := *distance * emission

	if co2 >= 1000 {
		*outputUnit = "kg"
	}

	var formattedString string
	if *outputUnit == "kg" {
		formattedString = fmt.Sprintf("Your trip caused %.1f%s of CO2-equivalent.", co2/1000, *outputUnit)
	} else {
		formattedString = fmt.Sprintf("Your trip caused %.0f%s of CO2-equivalent.", co2, *outputUnit)
	}

	fmt.Println(formattedString)
}
