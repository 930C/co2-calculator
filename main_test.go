package main

import "testing"

func TestCalculateCO2(t *testing.T) {
	cases := []struct {
		name         string
		method       string
		distance     float64
		distanceUnit string
		outputUnit   string
		expectedCO2  float64
	}{
		{
			name:         "Large Petrol Car 1800.5 km",
			method:       "large-petrol-car",
			distance:     1800.5,
			distanceUnit: "km",
			outputUnit:   "",
			expectedCO2:  507741,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			co2, err := calculateCO2(tc.method, tc.distanceUnit, tc.distance)

			if err != nil {
				t.Error(err)
			}

			if co2 != tc.expectedCO2 {
				t.Errorf("Wrong CO2 value, expected %f but got %f", tc.expectedCO2, co2)
			}
		})
	}
}
