package main

import "fmt"

func main() {
	fmt.Println("lets declare slices")
	// when we declare it we have to add some data to it , but in this there is no size constraint , we can extend it also
	var names = []string{"tanish", "mine", "xyz"}
	fmt.Println("names", names)

	// we cant extend it directly first we have to us  append function
	names = append(names, "bhatia", "its me")
	fmt.Println("updated names", names)

}
