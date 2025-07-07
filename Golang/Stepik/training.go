package main

import "fmt"

func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

func celsiusToFahrenheit(k celsius) fahrenheit {
	return fahrenheit(k*9.0/5.0) + 32.0
}

func celsiusToKelvin(k celsius) kelvin {
	return kelvin(k + 273.15)
}

type celsius float64
type kelvin float64
type fahrenheit float64

func main() {
	var kelvin kelvin = (127 + 273.15)
	celsius := kelvinToCelsius(kelvin)
	fahrenheit := celsiusToFahrenheit(celsius)
	kelvin = celsiusToKelvin(celsius)
	fmt.Println(celsius, "° C is ", fahrenheit, "° F", kelvin, "° F")
}
