package main

import (
	"flag"
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
		transportation = flag.String("transportation-method", "", "Transporation Method")
		distance       = flag.Float64("distance", 0, "Distance")
		distanceUnit   = flag.String("unit-of-distance", "km", "Defines the unit of the distance (km or m)")
		outputUnit     = flag.String("output", "g", "Output unit (g or kg)")
	)

	flag.Parse()
}
