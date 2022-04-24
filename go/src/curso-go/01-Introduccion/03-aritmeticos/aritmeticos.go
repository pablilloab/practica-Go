package main

import "fmt"

func main() {
	a := 20
	b := 10

	//Suma
	result := a + b

	fmt.Println("Suma", result)

	//Resta
	result = a - b
	fmt.Println("Resta", result)

	//Multiplicacion
	result = a * b
	fmt.Println("Multiplicacion", result)

	//Division
	result = a / b
	fmt.Println("Division", result)

	//Division exacta
	var div float64 = 3.0 / 2.0
	fmt.Println("Division exacta", div)

	//Modulo o resto
	result = 3 % 2
	fmt.Println("Modulo", result)

}
