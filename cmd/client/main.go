package main

import (
	"flag"
	"fmt"
	"os"

	"example.com/m/v2/client"
)

var (
	backendAPIURL = flag.String("backend", "http://localhost:8080", "api endpoint")
	helpFlag      = flag.Bool("help", false, "information on flag usage")
)

func main() {
	flag.Parse()
	s := client.NewSwitch(*backendAPIURL)

	if *helpFlag || len(os.Args) == 1 {
		s.Help()
		return
	}

	err := s.Switch()
	if err != nil {
		fmt.Printf("cmd switch error: %v\n", err)
		os.Exit(2)
	}
}
