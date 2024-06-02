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

	input, _ := reader.ReadString('\n')

	fmt.Println(input)
	// here we strconv to v=convert string into int we use ParseFloat to change into number , trim space to trim the line which is
	// coming when we print input becaure we use Println
	extra, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(extra + 1)

	}

}
