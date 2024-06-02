package main

import (
	"fmt"
	"time"
)

func main() {
	present := time.Now()
	fmt.Println(present)

	fmt.Println(present.Format("01-02-2006 Monday 15:04:05"))
}
