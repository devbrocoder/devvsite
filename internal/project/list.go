package main

import (
	"fmt"
	"os"
)

func runProjectList() {
	var args = os.Args[1:]
	fmt.Println(args)
}
