package main

import "fmt"
import "flag"

var infile = flag.String("i", "infile", "File input")

func main() {
	flag.Parse()
	if infile != nil {
		fmt.Println(*infile)
	}
}
