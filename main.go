package main

import (
	"amarbank/cmd"
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalln("application cannot start", err)
	}
}
