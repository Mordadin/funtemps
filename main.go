package main

import (
	"flag"
	"fmt"

	"conv_test.go/conv"
)

var fahr float64
var out string
var funfacts string

func init() {

	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
}

func main() {

	flag.Parse()

	fmt.Println("Enter the temperature in Fahrenheit:")
	fmt.Scanln(&fahr)

	fmt.Println("Enter the output unit (C or K):")
	fmt.Scanln(&out)

	// Convert Fahrenheit to Celsius or Kelvin

	if out == "C" {
		celsius := conv.FarhenheitToCelsius(fahr)
		fmt.Printf("%.2f degrees Fahrenheit is %.2f degrees Celsius\n", fahr, celsius)
	} else if out == "K" {
		kelvin := conv.FarhenheitToKelvin(fahr)
		fmt.Printf("%.2f degrees Fahrenheit is %.2f Kelvin\n", fahr, kelvin)
	} else {
		fmt.Println("Invalid output unit specified")
	}

	if out == "C" && isFlagPassed("F") {
		// Kalle opp funksjonen FahrenheitToCelsius(fahr), som da
		// skal returnere °C
		fmt.Println("0°F er -17.78°C")
	}

}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
