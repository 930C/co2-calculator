package main

import (
	"flag"
	"fmt"
	"log"
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

var (
	transportationMethod = flag.String("transportation-method", "", "Transporation Method")
	distance             = flag.Float64("distance", 0, "Distance")
	distanceUnit         = flag.String("unit-of-distance", "km", "Defines the unit of the distance (km or m)")
	outputUnit           = flag.String("output", "g", "Output unit (g or kg)")
)

func main() {
	flag.Parse()

	co2, err := calculateCO2(*transportationMethod, *distanceUnit, *distance)

	if err != nil {
		log.Fatalln(err)
	}

	formattedString := formatOutput(co2, *outputUnit)

	fmt.Println(formattedString)
}

// calculateCO2 calculates the CO2 emission based on the used transportationMethod, distance and it's unit.
func calculateCO2(transportationMethod, distanceUnit string, distance float64) (float64, error) {
	// Get emission rate of transportation method
	emissionRate, ok := transportationMethods[transportationMethod]

	if !ok {
		return 0, fmt.Errorf("error 404: methode %s wurde nicht gefunden", transportationMethod)
	}

	if distanceUnit == "m" {
		distance /= 1000
	}

	// calculate co2
	return distance * emissionRate, nil
}

// formatOutput formates the co2 result to a better readable string format.
func formatOutput(co2 float64, unit string) string {
	if co2 >= 1000 {
		unit = "kg"
	}

	if unit == "kg" {
		return fmt.Sprintf("Your trip caused %.1f%s of CO2-equivalent.", co2/1000, unit)
	} else {
		return fmt.Sprintf("Your trip caused %.0f%s of CO2-equivalent.", co2, unit)
	}
}
