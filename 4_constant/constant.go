package main

import "fmt"

const name string = "Alexa" // constant can be declared outside of main function
func main(){
	//constant can not be changed
	const pi = 3.1416
	fmt.Println(pi)

	fmt.Println(name)

	//multiple constants
	const (
		name = "Toukir"
		id = 12321105
		isActive = true
	)

	// name = Sarder X
	fmt.Println(name, id, isActive)
}