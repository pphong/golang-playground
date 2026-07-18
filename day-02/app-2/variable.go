package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	name string
	age  int
)

func readInfo() {
	fmt.Print("Input your name: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		name = scanner.Text()
	}
}

func readInfoFmt() {
	fmt.Println("Input your name: ")
	fmt.Scan(&name)
	fmt.Println("Input your age: ")
	fmt.Scan(&age)
	fmt.Printf("You are %s, %d yo", name, age)
}

func readInfoFmtSprint() {
	fmt.Println("Input your name: ")
	fmt.Scan(&name)
	fmt.Println("Input your age: ")
	fmt.Scan(&age)
	msg := fmt.Sprintf("You are %s, %d yo", name, age)
	fmt.Print(msg)
}
