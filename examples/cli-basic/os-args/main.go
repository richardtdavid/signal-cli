package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("no commands provided")
		os.Exit(2)
	}

	cmd := os.Args[1]
	//switch cmd {
	//case "address":
	//msg := "Signal CLI BASICS"
	//if len(os.Args) > 2 {
	//flags := strings.Split(os.Args[2], "=")
	//if len(flags) == 2 && flags[0] == "--msg" {
	//msg = flags[1]
	//}
	//}

	//fmt.Printf("Welcome to %v\n", msg)

	//case "help":
	//fmt.Println("help message")
	//default:
	//fmt.Printf("Unknown command: %v\n", cmd)
	//}

	// using flag sets
	switch cmd {
	case "address":
		msgCmd := flag.NewFlagSet("address", flag.ExitOnError)
		msgFlag := msgCmd.String("msg", "SIGNAL CLI BASICS", "Intro message")
		err := msgCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Printf("Welcome to %v\n", *msgFlag)
	case "help":
		fmt.Println("help message")
	default:
		fmt.Printf("Unknown command: %v\n", cmd)
	}
}
