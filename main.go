package main

import (
	"afa7789/site/cmd"
	"afa7789/site/internal/domain"
	"flag"
	"fmt"
	"log"
)

var (
	flags domain.Flags
)

func init() {
	// env = flag.String("env", "development", "current environment")
	flags.Port = flag.Int("port", 8080, "port number to listen")
}

func main() {
	flag.Parse()
	if err := cmd.ServerExecute(flags); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("")
		log.Printf("SERVER HERE: http://localhost:%d\n", *flags.Port)
	}
}
