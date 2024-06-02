package main

import "fmt"

func main() {
	fmt.Println("this file is for the pointers...")
	fmt.Print("normal declaration of pointers :")
	var tan *int
	fmt.Println("the value of pointer is", tan)

	fmt.Println("initialise of pointer refernce to some memory location :")

	marks := 34
	ptr := &marks
	fmt.Println("printing ptr:", ptr)
	fmt.Println("value of ptr::", *ptr)

}
