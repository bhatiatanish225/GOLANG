package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the number: ")
	input, _ := reader.ReadString('\n')
	fmt.Println(input)

	conv, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("no. incresed by 5 :")
		fmt.Println(conv + 5)
	}

}
