package main

import "fmt"

var varGlo = "this is global var"

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
	fmt.Println(varGlo)
}
