package main

import "fmt"

func main() {
	hola := "Hola"
	mundo := "Mundo"

	//Imprime con salto de linea
	fmt.Println(hola)
	fmt.Println(mundo)

	nombre := "pablillo"
	edad := 39

	//Formatea la saluda, se puede impirmir variables
	fmt.Printf("Hola, %s su edad es %d \n", nombre, edad)
	fmt.Printf("Hola, %v su edad es %v \n", nombre, edad)

	//Formatea el resultado de la expresion para guardarlo en una variable
	mensaje := fmt.Sprintf("Hola, %v su edad es %v", nombre, edad)
	fmt.Println(mensaje)

	//Imprime el tipo de la variable
	fmt.Printf("nombre %T \n", nombre)

	//Para capturar una palabra desde teclado
	fmt.Print("Ingresa otro nombre: ")
	fmt.Scanln(&nombre)

	fmt.Println(nombre)

}
