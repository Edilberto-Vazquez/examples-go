package main

import (
	"fmt"
)

func primeraLineaCodigo() {
	fmt.Println("Hola Mundo")
}

func varConstZeroval() {
	// Declaracion de constantes
	const pi float64 = 3.14 // Forma 1
	const pi2 = 3.14        // Forma 2

	fmt.Println(pi)
	fmt.Println(pi2)

	// Declaracion de variables enteras
	// Si las variables no se usan el programa no compila
	base := 12          // Primera forma
	var altura int = 14 // Segunda forma
	var area int        // Tercera fomra

	fmt.Println(base, altura, area)

	// Zero values
	var a int     // Asigna un 0
	var b float64 // Asigna un 0
	var c string  // Asigna una cadena vacia
	var d bool    // Asigna un false

	fmt.Println(a, b, c, d)

	// Area cuadrado
	areaCuadrado := base * altura

	fmt.Println(areaCuadrado)
}

func operadoresAritmeticos() {
	x := 50
	y := 10

	// Suma
	result := x + y
	fmt.Println(result)

	// Resta
	result = x - y
	fmt.Println(result)

	// Multiplicacion
	result = x * y
	fmt.Println(result)

	// Division
	result = x / y
	fmt.Println(result)

	// Modulo
	result = x % y
	fmt.Println(result)

	// Incremental
	x++
	fmt.Println(x)

	// Decremental
	x--
	fmt.Println(x)
}

func datosPrimitivos() {
	// Datos de tipo Int
	fmt.Println("Datos de tipo Int")
	fmt.Println("int = Depende del OS (35 o 64)")
	fmt.Println("int8 = 8 bits = -128 a 127")
	fmt.Println("int16 = 16 bits = -2^15 a 2^15-1")
	fmt.Println("int32 = 32 bits = -2^31 a 2^31-1")
	fmt.Println("int64 = 64 bits = -2^63 a 2^63-1\n")

	// Datos de tipo Int entero positivos
	fmt.Println("Datos de tipo Int entero positivos")
	fmt.Println("uint = Depende del OS (35 o 64)")
	fmt.Println("uint8 = 8 bits = 0 a 2^8-1")
	fmt.Println("uint16 = 16 bits = 0 a 2^16-1")
	fmt.Println("uint32 = 32 bits = 0 a 2^32-1")
	fmt.Println("uint64 = 64 bits = 0 a 2^64-1\n")

	// Decimales
	fmt.Println("Decimales")
	fmt.Println("float32 = 32 bits = +/-1.18e^-38 a +/-3.4e^38")
	fmt.Println("float32 = 32 bits = +/-2.23e^-38 a +/-1.8e^38\n")

	// Textos y boleanos
	fmt.Println("Textos y boleanos")
	fmt.Println("string = cadena vacia")
	fmt.Println("bool = true o false\n")

	// Numeros complejos
	fmt.Println("Numeros complejos")
	fmt.Println("Complex64 = Real e imaginario float32")
	fmt.Println("Complex128 = Real e imaginario float64")
	fmt.Println("Ejemplo: c := 10 + 8i")

}

func paqueteFmt() {
	// Declaracion de variables
	helloMessage := "Hello"
	worldMessage := "World"

	// Println
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	// Printf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene m??s de %d cursos\n", nombre, cursos) // Si se conoce el tipo de dato a imprimir
	fmt.Printf("%v tiene m??s de %v cursos\n", nombre, cursos) // Si no se conoce el tipo de dato a imprimir

	// Sprintf guarda un string en una variable
	message := fmt.Sprintf("%s tiene m??s de %d cursos", nombre, cursos)
	fmt.Println(message)

	// Tipo de dato
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)
}

// Ejemplo d funcion que returna 2 valores
func doubleReturn(a int) (b, c int) {
	b = a
	c = a * 2
	return b, c
}

func ciclos() {
	// For condicional
	for i := 0; i <= 10; i++ {
		fmt.Print(i, " ")
	}

	fmt.Print("\n")

	// For while
	counter := 0
	for counter <= 10 {
		fmt.Print(counter, " ")
		counter++
	}

	fmt.Print("\n")

	// For forever
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}
}

func condicionalIf() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es verdad")
	} else {
		fmt.Println("No es verdad")
	}
}

func condicionalSwitch() {
	// Switch
	switch modulo := 4 % 2; modulo {
	case 0:
		fmt.Println("es par")
	default:
		fmt.Println("es impar")
	}

	// Switch sin condicion
	value6 := 100
	switch {
	case value6 > 100:
		fmt.Println("es mayor a 100")
	case value6 < 0:
		fmt.Println("es menor a 0")
	default:
		fmt.Println("No condicion")
	}
}

func keywords() {
	// defer ejecuta la ultima funcion
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	// continue y break
	for i := 0; i <= 10; i++ {
		fmt.Println(i)

		//continue
		if i == 2 {
			fmt.Println("El ciclo continua")
			continue
		}

		if i == 8 {
			fmt.Println("El ciclo se cierra")
			break
		}
	}
}

func arraysSlices() {
	// Array
	var array [4]int // tambien pueden ser elementos de tipo string float o boleanos
	array[0] = 1
	array[1] = 2
	fmt.Println("Array")
	fmt.Println(array, len(array), cap(array), array[0:2])

	// Slice
	fmt.Println("Slice")
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Metodos en el slice
	fmt.Println("Metodos slice")
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append
	fmt.Println("Agregar elementos con Append")
	slice = append(slice, 7)
	fmt.Println(slice)
	// agregar nueva lista
	fmt.Println("Agregar otra lista con Append")
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}

func sliceRange() {
	slice := []string{"hola", "que", "hace"}

	for i, valor := range slice {
		fmt.Println(i, valor)
	}
}

func isPalindromo(text string) {
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}

func maps() {
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	// Recorrer un map
	for i, v := range m {
		fmt.Println(i, v)
	}

	// Encontrar un valor
	value, ok := m["Jose"]
	fmt.Println(value, ok)
}

type car struct {
	brand string
	year  int
}

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) pink() {
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2
}

func (myPC pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y es una %s:", myPC.ram, myPC.disk, myPC.brand)
}

//-------interface------

type cuadrado struct {
	base float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

type rectangulo struct {
	base   float64
	altura float64
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

type figuras2D interface {
	area() float64
}

func calcular(f figuras2D) {
	fmt.Println("area:", f.area())
}

// func say(text string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println(text)
// }

func say(text string, c chan<- string) {
	c <- text

	//text = <- c
}

func message(text string, c chan string) {
	c <- text
}

func main() {
	// primeraLineaCodigo()
	// varConstZeroval()
	// operadoresAritmeticos()
	// datosPrimitivos()
	// paqueteFmt()

	// Funcion con doble retorno
	// b, c := doubleReturn(2) // Desempaqueta los 2 valores
	// b, _ = doubleReturn(2)  // Solo desempaqueta un valor
	// fmt.Println(b, c)

	// ciclos()
	// condicionalIf()
	// condicionalSwitch()
	// keywords()
	// arraysSlices()
	// sliceRange()
	// isPalindromo(strings.ToLower("Ama"))
	// maps()

	// Struct parecidas a las clases
	// myCar := car{brand: "Ford", year: 2020}
	// fmt.Println(myCar)

	// Otra manera
	// var otherCar car
	// otherCar.brand = "Ferrari"
	// fmt.Println(otherCar.brand)
	// fmt.Println(otherCar.year)

	// Structs
	// var myCar mypackage.CarPublic
	// myCar.Brand = "Ferrari"
	// myCar.Year = 2020
	// fmt.Println(myCar)

	// mypackage.PrintMessage("Hola platzi")

	// ---------Punteros---------
	// a := 50
	// b := &a

	// fmt.Println(b)
	// fmt.Println(*b)

	// *b = 100

	// fmt.Println(a)

	// myPC := pc{ram: 16, disk: 200, brand: "msi"}

	// fmt.Println(myPC)

	// myPC.pink()

	// fmt.Println(myPC)
	// myPC.duplicateRAM()

	// fmt.Println(myPC)
	// myPC.duplicateRAM()

	// fmt.Println(myPC)

	// -------Ejemplo punteros-------
	// myPC := mypackage.New(16, 1000, "Dell")

	// fmt.Println(myPC)

	// myPC.DuuplicateRAM()

	// fmt.Println(myPC)

	// myPC := pc{ram: 16, disk: 1000, brand: "Dell"}

	// fmt.Println(myPC)

	// interfaces
	// myCuadrado := cuadrado{base: 2}
	// myRectangulo := rectangulo{base: 2, altura: 4}

	// calcular(myCuadrado)
	// calcular(myRectangulo)

	// Lista interfaces
	// myInterface := []interface{}{"hola", 12, 4.90}
	// fmt.Println(myInterface)

	//----------go ruotines waitgroup----------
	// var wg sync.WaitGroup

	// fmt.Println("Hello")
	// wg.Add(1)
	// go say("world", &wg)

	// wg.Wait()

	// go func(text string) {
	// 	println(text)
	// }("Adios")

	// time.Sleep(time.Second * 1) // no es eficiente agregar tiempos de espera

	// ---------chanels---------
	// c := make(chan string, 1)

	// fmt.Println("hello")

	// go say("Bye", c)

	// fmt.Println(<-c)

	c := make(chan string, 2)
	c <- "mensaje1"
	c <- "mensaje2"

	fmt.Println(len(c), cap(c))

	// range y close
	close(c)

	// c <- "mensaje3"

	for message := range c {
		println(message)
	}

	// Select
	email1 := make(chan string)
	email2 := make(chan string)

	go message("mensaje", email1)
	go message("mensaje", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("email recibido de email1", m1)
		case m2 := <-email2:
			fmt.Println("email recibido de email2", m2)
		}
	}
}
