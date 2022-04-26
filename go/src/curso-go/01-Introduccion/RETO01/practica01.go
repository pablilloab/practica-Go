package main

import "fmt"

func sumar(num1, num2 int) int {
	return num1 + num2
}

func main() {

	//Ejercicio 1: sumar 2 numeros
	var num1 int
	var num2 int

	fmt.Printf("Ingrese el primer numero :")
	fmt.Scan(&num1)
	fmt.Printf("Ingrese el segundo numero :")
	fmt.Scan(&num2)

	suma := sumar(num1, num2)

	fmt.Printf("La suma es %d \n", suma)

	//Ejercicio 2: calcular cociente y resto de dos numeros
	fmt.Printf("Ingrese dividendo")
	fmt.Scan(&num1)
	fmt.Printf("Ingrese divisor")
	fmt.Scan(&num2)

	cociente := num1 / num2
	resto := num1 % num2

	fmt.Printf("El cociente es %d y el resto %d \n", cociente, resto)

	var precio float64
	fmt.Print("Ingrese el valor para calcular su precio de venta")
	fmt.Scan(&precio)

	igv := precio + (precio * 0.18)

	fmt.Printf("El valor de venta es de %f ", igv)

}
