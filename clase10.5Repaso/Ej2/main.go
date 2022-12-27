package main

import (
	"fmt";
	"os";
	"encoding/json"
)

type Logger interface {
	Log(msg interface{})
}

type CalculatorCasio struct {
	register Logger
}

func (c *CalculatorCasio) Add(x,y float64) float64 {
	result := x+y
	c.register.Log(result)
	return result
}
func (CalculatorCasio) Multiply(x,y float64) float64 {
	return x*y
}
func (CalculatorCasio) Divide(x,y float64) float64 {
	return x/y
}


//Calculator is capable of making the 4 basic calc operations
type Calculator interface {
	Add(x,y float64) float64
	Multiply(x,y float64) float64
	Divide(x,y float64) float64

}
// LoggerToConsole logs inputs to the console
type LoggerToConsole struct{}

func (LoggerToConsole) 	Log(msg interface{}) {
	fmt.Println(msg)
}

// LoggerToFile logs inputs to the console
type LoggerToFile struct{}

func (LoggerToFile) Log(msg interface{}) {

	toLog, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("log.txt", toLog, 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	//var myCalc Calculator

	myLogger := LoggerToFile{}
	myCalc := CalculatorCasio{
		register: myLogger,
	}

	myCalc.Add(2,3)
	myCalc.Multiply(2,3)
	myCalc.Divide(2,3)
}