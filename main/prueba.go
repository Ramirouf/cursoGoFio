package main

import "fmt"
import "github.com/gin-gonic/gin"
import "clase2IntroGo"


func main(){
	fmt.Println("Hello world")
	fmt.Println(gin.Version)

	integer, str := MakeSomething("Juan", "Perez")
	MakeSomething("Ramiro", "Uffelmann")
	fmt.Println(integer, str)
}
/*
gomodcache 
gotooldir
*/

