package main

import (
	"flag"
	"fmt"
)

var flagR bool
var flagT bool
var flagV int
var flagName string
var flagAge int

func main() {
	// ./main_2.exe -r -t=false -v=7 -name=lx -age=17
	flag.BoolVar(&flagR, "r", false, "flagR")
	flag.BoolVar(&flagT, "t", false, "flagT")
	flag.IntVar(&flagV, "v", 0, "flagV")
	flag.StringVar(&flagName, "name", "", "flagName")
	flag.IntVar(&flagAge, "age", 0, "flagAge")

	flag.Parse()

	fmt.Println("Parsed:", flag.Parsed()) //Parsed: true
	fmt.Println(" NFlag:", flag.NFlag())  //  NArg: 5
	fmt.Println("     r:", flagR)         //     r: true
	fmt.Println("     t:", flagT)         //     t: false
	fmt.Println("	  v:", flagV)           //     v: 7
	fmt.Println("  name:", flagName)      //  name: lx
	fmt.Println("   age:", flagAge)       //   age: 17
}
