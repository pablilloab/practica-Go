package main

import "fmt"

func saludar(nombre string) {
	fmt.Printf("Hola %s \n", nombre)
}

//La declaracion del tipo q va a retorna se hace luego de la declaracion de parametros
func sumar(num1, num2 int) int {
	return num1 + num2
}

func main() {
	var nombre string
	fmt.Print("Ingrese su nombre: ")
	fmt.Scanln(&nombre)

	saludar(nombre)

	suma := sumar(10, 20)
	fmt.Println("La suma es ", suma)

}
