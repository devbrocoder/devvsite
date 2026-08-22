package main

import (
	"fmt"
	"os"
)

func runPhpFpm() {
	var args = os.Args[1:]
	fmt.Println(args)
}
