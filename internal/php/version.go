package main

import (
	"fmt"
	"os"
)

func runPhpVersion() {
	var args = os.Args[1:]
	fmt.Println(args)
}
