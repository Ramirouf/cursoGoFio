/*
Ejercicio 1
01 | Arrays y slices

Crear un array de 10 elementos del tipo entero.
Recorrerlo y loguear los que son impares.
Crear un slice que consista en los elementos [3,7) de dicho string.
Crear una función que reciba un slice de enteros y devuelva su promedio.

*/
package main
import "fmt"

func main(){
	var array [10]int
	var i int
	for i=0; i<10; i++{
		array[i] = i*3
	}
	//Print odd numbers
	for i=0; i<10; i++{
		if array[i]%2 != 0{
			fmt.Println(array[i])
		}
	}
	fmt.Println("Print entire array")
	for i=0; i<10;i++ {
		fmt.Println(array[i])
	}
	fmt.Println("Print slice with items [3,7) from array")
	var slice []int
	slice = array[3:7]
	fmt.Println(slice)
	fmt.Println("Average of slice: ")
	fmt.Println(average(slice))
}


func average (slice []int) float32{
	var sum float32
	var i int
	for i=0; i<len(slice);i++{
		sum += float32(slice[i])
	}
	averageValue := float32(sum/float32(len(slice)))
	return averageValue
}