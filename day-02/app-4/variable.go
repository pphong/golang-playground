package main

import (
	"fmt"
)

var (
	name string
	age  int
)

func loop() {
	fmt.Println("Loop ...")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func loop2() {
	fmt.Println("Loop2 ...")
	var i int = 10
	for i < 20 {
		fmt.Println(i)
		i++
	}
}

func loop3() {
	fmt.Println("Loop3 ...")
	var i int = 10
	for i < 100 {
		fmt.Println("________________")
		fmt.Printf("current: %d \n", i)
		fmt.Println("________________")

		fmt.Printf("stop at 95: %t \n", i == 95)
		if i == 95 {
			break
		}
		if i%5 == 0 {
			fmt.Printf("%%5: %d \n", i)
			i++
			continue
		}

		i++
	}
}
