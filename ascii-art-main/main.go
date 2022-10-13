package main

import (
	"art"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	inputArgs := os.Args
	if !art.HashCheck("standard.txt") {
		fmt.Println("Error: Your banner file is corrupted")
		return
	}
	if len(inputArgs) != 2 {
		fmt.Println("Error: Wrong Input! (Info: go run main.go Text)")
		return
	}
	if art.CheckLetter(inputArgs[1]) {
		inData, err := ioutil.ReadFile("standard.txt")
		art.CheckErr(err)
		str1 := string(inData)
		values := art.Scan(str1)
		final := art.MakeMap(values)
		arg := art.Split(inputArgs[1])
		art.Print(arg, final)
	} else {
		fmt.Println("Error: Letter isn't latin")
	}
}
