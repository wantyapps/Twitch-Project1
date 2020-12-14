package main

import (
	"os" // For CLI-Arguments and stuff like this
	"fmt" // You know what fmt is
)

func noDebugMode() {
	fmt.Println("NoDebug.")
}

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "--nodebug" || os.Args[1] == "-n" {
			noDebugMode()
		} else {
			for i := 1; i <= len(os.Args) -1; i++ {
				fmt.Println(os.Args[i])
			}
		}
	}
}
