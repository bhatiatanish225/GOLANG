package main

import "fmt"

// if we want to declare anything publicly we can make the variable name capital

const Isloggedin bool = true

func variables() {
	var ussername string = "Tanish Bhatia"
	var website = "learncodeonline.in"
	fmt.Println(ussername)
	fmt.Printf("Variable is of %T type", ussername)
	fmt.Println("")
	fmt.Println(website)
	fmt.Printf("Variable is of %T type", website)
	// if use the keyword var and does not use variable type we cant initialise it again
	// if we vant use keyword var , we can initialise it again , can change the varibale name any time by changing the value
	nofusserss := 20000
	fmt.Println(nofusserss)

}

func main() {
	variables()
}
