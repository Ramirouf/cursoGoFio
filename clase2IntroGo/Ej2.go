package main
import "fmt"
/*
Una empresa de meteorología quiere tener una aplicación donde pueda tener la temperatura y humedad y presión atmosférica de distintos lugares. 
1. Declara 3 variables especificando el tipo de dato, como valor deben tener la temperatura, humedad y presión de donde te encuentres. 
2. Imprime los valores de las variables en consola. 
3. ¿Qué tipo de dato le asignarías a las variables?

*/
func main(){
	var (
		temperature float32
		humidity float32
		pressure float32
	)
	temperature, humidity, pressure = 30, 90, 1000
	fmt.Println(temperature, humidity, pressure)
}