// package main

// import "fmt"

// func normalFunction(message string) {
// 	fmt.Println(message)
// }

// func tripleArgument(a, b int, c string) {
// 	fmt.Println(a, b, c)
// }

// func returnValue(a int) int {
// 	return a * 2
// }

// func doubleReturn(a int) (c, d int) {
// 	return a, a * 2
// }

// func numeroPar(num int) bool {
// 	if num%2 == 0 {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// func userPassword(userInput string, passwordInput int) bool {
// 	user := "Edilberto"
// 	password := 1234
// 	if user == userInput && password == passwordInput {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// func main() {
// 	// Declaración de constantes
// 	const pi float64 = 3.14
// 	const pi2 = 3.1416

// 	fmt.Println("pi:", pi)
// 	fmt.Println("pi2:", pi2)

// 	// Declaración de variables enteras
// 	base := 12          // Primera forma
// 	var altura int = 14 // Segunda forma
// 	var area int        // Tercera forma

// 	fmt.Println(base, altura, area)

// 	// Zero values
// 	var a int     // Asigna un 0
// 	var b float64 // Asigna un 0
// 	var c string  // String vacio
// 	var d bool    // Asigna un false

// 	fmt.Println(a, b, c, d)

// 	// Area cuadrado
// 	const baseCuadrado = 10
// 	areaCuadrado := baseCuadrado * baseCuadrado

// 	fmt.Println("Area cuadrado =", areaCuadrado)

// 	// Declaración de variables
// 	helloMessage := "Hello"
// 	worldMessage := "World"

// 	// Println
// 	fmt.Println(helloMessage)
// 	fmt.Println(worldMessage)

// 	// Printf
// 	nombre := "Platzi"
// 	cursos := 500

// 	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos) // Si se sabe el tipo de dato
// 	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos) // Si %v no se sabe el tipo de dato

// 	//Sprintf
// 	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)
// 	fmt.Println(message)

// 	// conocer el tipo de datos
// 	fmt.Printf("helloMessage: %T\n", helloMessage)
// 	fmt.Printf("cursos: %T\n", cursos)

// 	normalFunction("hola mundo")
// 	tripleArgument(1, 2, "hola")

// 	value := returnValue(2)

// 	fmt.Println("Value:", value)

// 	value1, _ := doubleReturn(2)

// 	fmt.Println("value1 y value 2:", value1)

// 	// For condicional
// 	for i := 0; i <= 10; i++ {
// 		fmt.Println(i)
// 	}

// 	fmt.Printf("\n")

// 	// For while
// 	counter := 0
// 	for counter < 10 {
// 		fmt.Println(counter)
// 		counter++
// 	}

// 	// For forever
// 	// counterForever := 0
// 	// for {
// 	// 	fmt.Println(counterForever)
// 	// 	counterForever++
// 	// }

// 	// For decremental
// 	// for i := 10; i >= 1; i-- {
// 	// 	fmt.Println(i)
// 	// }

// 	valor1 := 1
// 	valor2 := 2

// 	if valor1 == 1 {
// 		fmt.Println("es 1")
// 	} else {
// 		fmt.Println("No es 1")
// 	}

// 	// With and
// 	if valor1 == 1 && valor2 == 2 {
// 		fmt.Println("es verdad")
// 	}

// 	// With or
// 	if valor1 == 0 || valor2 == 2 {
// 		fmt.Println("Es verdad, OR")
// 	}

// 	fmt.Println(numeroPar(3))
// 	fmt.Println(userPassword("Edi", 123))

// 	// Switch
// 	switch modulo := 4 % 2; modulo {
// 	case 0:
// 		fmt.Println("es par")
// 	default:
// 		fmt.Println("es impar")
// 	}

// 	// Switch sin condicion
// 	value6 := 100
// 	switch {
// 	case value6 > 100:
// 		fmt.Println("es mayor a 100")
// 	case value6 < 0:
// 		fmt.Println("es menor a 0")
// 	default:
// 		fmt.Println("No condicion")
// 	}

// 	// Defer ejecuta la ultima linea de codigo de la funcon
// 	defer fmt.Println("Hola")
// 	fmt.Println("Mundo")

// 	// continue y break
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(i)
// 		//continue
// 		if i == 2 {
// 			fmt.Println("Es 2")
// 			continue
// 			//break
// 		}

// 		//break
// 		if i == 8 {
// 			fmt.Println("Break")
// 			break
// 		}
// 	}
// }

package main

import "fmt"

func primeraLineaCodigo() {
	fmt.Println("Hola Mundo")
}

func

func main() {
	primeraLineaCodigo()
}
