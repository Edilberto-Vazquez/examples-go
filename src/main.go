package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	// Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.1416

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// Declaración de variables enteras
	base := 12          // Primera forma
	var altura int = 14 // Segunda forma
	var area int        // Tercera forma

	fmt.Println(base, altura, area)

	// Zero values
	var a int     // Asigna un 0
	var b float64 // Asigna un 0
	var c string  // String vacio
	var d bool    // Asigna un false

	fmt.Println(a, b, c, d)

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("Area cuadrado =", areaCuadrado)

	// Declaración de variables
	helloMessage := "Hello"
	worldMessage := "World"

	// Println
	fmt.Println(helloMessage)
	fmt.Println(worldMessage)

	// Printf
	nombre := "Platzi"
	cursos := 500

	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos) // Si se sabe el tipo de dato
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos) // Si %v no se sabe el tipo de dato

	//Sprintf
	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)
	fmt.Println(message)

	// conocer el tipo de datos
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)

	normalFunction("hola mundo")
	tripleArgument(1, 2, "hola")

	value := returnValue(2)

	fmt.Println("Value:", value)

	value1, _ := doubleReturn(2)

	fmt.Println("value1 y value 2:", value1)
}
