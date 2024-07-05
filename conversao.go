package main

import "fmt"

var tempKelvin float64 = 373.15
var tempCelcius float64 = tempKelvin - 273

func main() {
	fmt.Printf("A temperatura em Kelvin é igual a: %.2f. \nA temperatura em Celcius é igual a: %.2f.", tempKelvin, tempCelcius)
}
