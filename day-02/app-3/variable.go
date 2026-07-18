package main

import (
	"fmt"
)

var (
	name string
	age  int
)

func calc() {
	fmt.Print("Calculate ...")
	var s1, s2 int
	s1 = 5
	s2 = 3
	sum := s1 + s2
	dif := s1 - s2
	pro := s1 * s2
	quo := float64(s1) / float64(s2)
	idi := s1 % s2
	mod := s1 / s2

	fmt.Printf("Sum of %d and %d is %d \n", s1, s2, sum)
	fmt.Printf("Dif of %d and %d is %d \n", s1, s2, dif)
	fmt.Printf("Pro of %d and %d is %d \n", s1, s2, pro)
	fmt.Printf("Quo of %d and %d is %.2f \n", s1, s2, quo)
	fmt.Printf("Idi of %d and %d is %d \n", s1, s2, idi)
	fmt.Printf("Mod of %d and %d is %d \n", s1, s2, mod)
}
