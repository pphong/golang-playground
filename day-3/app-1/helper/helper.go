package helper

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	option string
	path   string
	name   string
)

func createFile() {
	if exists, err := Exists(); err == nil {
		if exists {
			fmt.Printf("this file with name '%s' is already exists! \n", name)
			return
		}
		file, err := os.OpenFile(
			path,
			os.O_CREATE|os.O_EXCL,
			0644,
		)

		if err != nil {
			fmt.Println(err)
			return
		}

		defer file.Close()
		fmt.Printf("Created '%s' successfully! \n", name)
	} else {
		fmt.Println("Something went wrong!", err)
	}
}

func createFolder() {
	if exists, err := Exists(); err == nil {
		if exists {
			fmt.Printf("this folder with name '%s' is already exists! \n", name)
			return
		}

		err := os.Mkdir(path, 0755)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Created '%s' successfully! \n", name)
	} else {
		fmt.Println("Something went wrong!", err)
	}
}

func remove() {
	if exists, err := Exists(); err == nil {
		if !exists {
			fmt.Printf("this folder/ file with name '%s' is not exists! \n", name)
			return
		}
	} else {
		fmt.Println("Something went wrong!", err)
	}
}

func Exists() (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	} else if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else {
		fmt.Println(err)
		return false, err
	}
}

func helper() {
	fmt.Println("-----------------")
	fmt.Println("PPhong Go CLI")
	fmt.Println("-----------------")
	fmt.Println("-h : this menu")
	fmt.Println("-d : create a folder (e.g: -d <fodler name>)")
	fmt.Println("-f : create an empty file (e.g -f <file name>)")
	fmt.Println("-rmf : remove a folder/ file (e.g -rmf <folder/file name>)")
}

func executeCmd() {
	pathsArr := strings.Split(path, "/")
	name = pathsArr[len(pathsArr)-1]
	switch option {
	case "-d":
		{
			createFolder()
			break
		}
	case "-f":
		{
			createFile()
			break
		}
	case "-rmf":
		{
			remove()
			break
		}
	default:
		errorThrow()
	}

}

func errorThrow() {
	fmt.Println("Syntax error")
	fmt.Println("Usage: PPc <option> <path>")
	helper()
}

func readCommand() {
	if len(os.Args) < 3 {
		if os.Args[1] == "-h" {
			helper()
		} else {
			errorThrow()
			return
		}
		return
	}
	option = os.Args[1]
	path = os.Args[2]
	fmt.Printf("Option: %s, uri: %s \n", option, path)
	fmt.Println("-----------------")
	executeCmd()
}

func Cli() {
	readCommand()
}
