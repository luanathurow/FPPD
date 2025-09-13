package main

import "fmt"

//fmt é uma lib p print, error, etc...

func main() {

	fmt.Println(sum(1, 2))
	fmt.Println("Hello World")

	//var name string
	//name = "luana"
	name := "Luana" //ele já identifica que é uma string
	age := 22       //ele já identifica que é um int
	//dev := true     //ele já identifica que é um boo

	fmt.Println(name)

	if age > 18 {
		fmt.Println("Você é adulto!")
	} else {
		fmt.Println("Você é criança!")
	}

	//LAÇOS:
	sum := 0
	for i := 0; i <= 10; i++ {
		fmt.Println("Valor de I: ", i)
		sum += i
	}
	fmt.Println(sum)

	//array
	frutas := []string{"morango", "banana"}

	for l, frutas := range frutas {
		fmt.Println("Fruta:", frutas)
		fmt.Println("Posicao", l)
	}
}

func sum(num1, num2 int) int {
	return num1 + num2
}
