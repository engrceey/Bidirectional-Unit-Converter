package main

import (
	"fmt"
	"math"
)

// Converter struct to call in the methods
type Converter struct{}

type (
	// Minutes type using float64
	Minutes float64
	// Seconds type using float64
	Seconds float64
	// Centimeter type using float64
	Centimeter float64
	// Feet type using float64
	Feet float64
	// Celsius type using float64
	Celsius float64
	// Fahrenheit type using float64
	Fahrenheit float64
	// Degree type using float64
	Degree float64
	// Radian type using float64
	Radian float64
	//Pounds type using float64
	Pounds float64
	// Kilogram type using float64
	Kilogram float64
)

// MinutesToSeconds conversion calculation
func (conv Converter) MinutesToSeconds(m Seconds) Seconds {
	sec := m * 60
	return sec
}

// SecondsToMinutes conversion calculation
func (conv Converter) SecondsToMinutes(s Minutes) Minutes {
	min := s / 60
	return min
}

// CentimeterToFeet conversion calculation
func (conv Converter) CentimeterToFeet(cm Feet) Feet {
	ft := cm / 30.48
	return ft
}

// FeetToCentimeter conversion calculation
func (conv Converter) FeetToCentimeter(ft Centimeter) Centimeter {
	cm := ft * 30.48
	return cm
}

// CelsiusToFahrenheit conversion calculation
func (conv Converter) CelsiusToFahrenheit(cel Fahrenheit) Fahrenheit {
	fah := (cel * 1.8) + 32
	return fah
}

// FahrenheitToCelsius conversion calculation
func (conv Converter) FahrenheitToCelsius(fah Celsius) Celsius {
	cel := (fah - 32) / 1.8
	return cel
}

// RadianToDegree conversion calculation
func (conv Converter) RadianToDegree(rand Degree) Degree {
	Deg := rand * (180 / math.Pi)
	return Deg
}

// DegreeToRadian conversion calculation
func (conv Converter) DegreeToRadian(deg Radian) Radian {
	rand := deg * (math.Pi / 180)
	return rand
}

// KilogramToPounds conversion calculation
func (conv Converter) KilogramToPounds(kilo Pounds) Pounds {
	pnds := kilo * 2.2046226218
	return pnds
}

// PoundsToKilo conversion calculation
func (conv Converter) PoundsToKilo(pnds Kilogram) Kilogram {
	kilo := pnds * 0.45359237
	return kilo
}

func main() {
	con := Converter{}

	// Change the Method and values for different conversions
	fmt.Printf("%.3f\n", con.PoundsToKilo(5))
}
