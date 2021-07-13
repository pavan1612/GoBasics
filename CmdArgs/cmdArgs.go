package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	var arg = os.Args[1]
	// Separate a file form its path
	dir, file := path.Split(arg)
	fmt.Println(dir, file)
}
