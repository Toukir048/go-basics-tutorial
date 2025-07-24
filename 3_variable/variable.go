package main

import "fmt"

func main() {
	//string
	//var name string = "Alex"
	var name = "Alice Bob"  // go lang can infer type by following given value
	fmt.Println(name)

	var age int  = 30 
	fmt.Println(age)

	var isAdult = true
	fmt.Println(isAdult)

	//shorthand declaration
	city := "Dhaka" //  := is shorthand for var city string 
	fmt.Println(city)

	//where to use var name data type

	var roll int

	roll = 101
	fmt.Println(roll)


}