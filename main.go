package main

import (
	"afa7789/site/cmd"
	"afa7789/site/internal/domain"
	"flag"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

var flags domain.Flags

func init() {
	flags.Port = flag.Int("port", 8080, "port number to listen")
	tls := flag.Bool("TLS", false, "a bool")
	flag.Parse() // add this line
	flags.TLS = tls

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("env loading error", err)
	}
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
