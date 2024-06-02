package main

import "fmt"

func main() {
	fmt.Println("here we will discus array ")

	var fruits [5]string

	fruits[0] = "apple"
	fruits[1] = "mango"
	fruits[2] = "orange"
	fruits[3] = "kivi"
	fmt.Println(fruits)
	fmt.Println(len(fruits))

	fmt.Println("making array adn direct putting the values in it..")
	var veglist = [3]string{"tomato", "potato", "onion"}
	fmt.Println(veglist)

}
