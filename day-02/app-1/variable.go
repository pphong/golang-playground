package main

import "fmt"

var (
	varGlo1 = "this is global var 1"
	varGlo2 = "this is global var 2"
)

func vars() {
	var varVariable string = "this is variable using var"
	fmt.Println(varVariable)
}

func varsWithColon() {
	varVariable := "this is variable using colon"
	fmt.Println(varVariable)
}

func varsWithDefault() {
	var varDefault int
	fmt.Println(varDefault)
}

func varsWithInt() {
	var varInt int = 30
	fmt.Println(varInt)
}

func varsWithGlo() {
	fmt.Println(varGlo1)
}
