package main

import (
	"fmt"
	"os"
)

func runMysqlList() {
	var args = os.Args[1:]
	fmt.Println(args)
}
