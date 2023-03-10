package main

import (
	"fmt"
)

func main() {
	var (
		i       = 21
		F       = 15
		j       = false
		persen  = "%"
		russian = '\u042F'
		decimal = 123.456000
	)

	fmt.Printf("%d \n", i)
	fmt.Printf("%T \n", i)
	fmt.Printf("%s \n", persen)
	fmt.Printf("%t \n", !j)
	fmt.Printf("%t \n", j)
	fmt.Printf("%c \n", russian)
	fmt.Printf("%d \n", i)
	fmt.Printf("%o \n", i)
	fmt.Printf("%x \n", F)
	fmt.Printf("%X \n", F)
	fmt.Printf("%U \n", russian)
	fmt.Printf("%f \n", decimal)
	fmt.Printf("%E \n", decimal)
}
