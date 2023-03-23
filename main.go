package main

import (
	"awesome/api"
	"flag"
	"fmt"
	"log"
)

func main() {
	listenAddr := flag.String("listenaddr", ":8081", "the server address")
	flag.Parse()

	server := api.NewServer(*listenAddr)
	fmt.Println("The server running on port", *listenAddr)
	log.Fatal(server.Start())
}
