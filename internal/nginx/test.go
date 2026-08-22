package main

import (
	"fmt"
	"os"
)

func runNginxTest() {
	var args = os.Args[1:]
	fmt.Println(args)
}
