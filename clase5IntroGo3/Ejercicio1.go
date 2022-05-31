/*
Crear una función simple que permita multiplicar dos números y devolver el resultado.

Crear una función que reciba un número variable de enteros como argumento y devuelva el menor de ellos.

Crear una estructura Animal, y varias estructuras que tomen sus propiedades 
y representen diferentes tipos de animales.
Crear además una función que permita instanciar un Animal 
y devuelva un puntero al mismo.

Crear un a función que permita obtener de forma recursiva el factorial de un número.
*/

package main

import "fmt"

func main() {
	result := multiply(2,3)
	fmt.Println(result)
	minValue := min(5,2,3,4,5,6,7,8,9,-4)
	fmt.Println(minValue)

	dog1 := Dog{
		Animal: Animal{
			name: "Fido",
			age: 2,
			color: "Brown",
		},
		breed: "Labrador",
	}

	fmt.Println(dog1.name)
	fmt.Println(dog1.breed)

	animalPointer := createAnimal()
	fmt.Println(animalPointer.name)
	fmt.Println(animalPointer.age)

	fmt.Println(factorial(5))
}
type Animal struct{
	name string
	age int
	color string
}

type Dog struct{
	Animal
	breed string
}

func createAnimal() *Animal{
	animal := Animal{
		name: "Bob",
		age: 14,
		color: "Black",
	}
	return &animal
}

func multiply(a,b int) int{
	return a*b
}

func min(numbers ...int) int{
	minValue := numbers[0]
	for _, number := range numbers{
		if number < minValue{
			minValue = number
		}
	}
	return minValue
}

func factorial(number int) int{
	if number == 0{
		return 1
	}
	return number * factorial(number-1)
}