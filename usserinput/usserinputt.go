package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// usser input
	//bufio is the lib through which we can make a reader
	// first we have to make reader and then make listener

	//reader
	welcome := "welcome to usser input"

	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the number!!: ")
	// this is a comma ok syntax ,err syntax
	// this means that either there could be valid input or a erro
	// _,err another syntax in this syntax we store error not the valid input

	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for entering ", input)

}
