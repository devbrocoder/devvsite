package main

import (
	"fmt"
	"os"
)

func runSystemCommand() {
	var args = os.Args[1:]
	fmt.Println(args)
}
