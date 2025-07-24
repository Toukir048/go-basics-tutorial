package main

import 	"fmt"


// for is the only loop in go lang
func main(){
	//while loop using for 
	// i := 1
	// for i <= 5{
	// 	fmt.Println(i*2)
	// 	i++
	// }

	// infinte loop concept of while loop
	// for {
	// 	fmt.Println("helllo")
	// }

	//classic for loop
	// for i:= 1; i <= 5; i++{
	// 	//break
	// 	if i == 3{
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// for go 1.22 range are added
	for i := range 3{
		fmt.Println(i)
	}
}