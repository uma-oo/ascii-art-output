package main

import (
	"fmt"
	"log"
	"os"

	function "ascii/functions"
)

func main() {
	args := os.Args[1:]

	switch len(args) {
	case 1:
		if !function.CheckFlagFormat(args[0]) {
			fmt.Print(function.Fs(args[0]))
		} else {
			function.Output(function.ParseFlag(args[0]), "")
			log.Fatalf("File created\nFlag '--output=' given but lacks more data")
			
		}

	case 2:
		if function.CheckFlagFormat(args[0]) {
			function.Output(function.ParseFlag(args[0]), args[1])
		} else {
			fmt.Print(function.Fs(args[0], args[1]))
		}
	case 3:
		function.Output(function.ParseFlag(args[0]), args[1], args[2])
	default:
		log.Fatal("There is no arguments! Please provide a one!")

	}
}
