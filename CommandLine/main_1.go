package main

import (
	"flag"
	"fmt"
)

func main() {
	// ./main_1.exe 0.txt 1.txt 2.txt
	flag.Parse()
	fmt.Println("Parsed:", flag.Parsed())       //Parsed: true
	fmt.Println("  NArg:", flag.NArg())         //  NArg: 3
	fmt.Printf(" Arg %d: %q\n", 0, flag.Arg(0)) // Arg 0: "0.txt"
	fmt.Printf(" Arg %d: %q\n", 1, flag.Arg(1)) // Arg 1: "1.txt"
	fmt.Printf(" Arg %d: %q\n", 2, flag.Arg(2)) // Arg 2: "2.txt"

	fmt.Println("-----------------")
	for i := 0; i < flag.NArg(); i++ {
		fmt.Printf(" Arg %d: %q\n", i, flag.Arg(i))
	}
}
