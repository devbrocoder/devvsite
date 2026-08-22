package main

import (
	"fmt"
	"os"
)

func runNginxReload() {
	var args = os.Args[1:]
	fmt.Println(args)
}
