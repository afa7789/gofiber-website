package main

import (
	"afa7789/site/cmd"
	"log"
)

func main() {
	if err := cmd.ServerExecute(); err != nil {
		log.Fatal(err)
	}
}
