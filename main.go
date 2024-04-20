package main

import (
	"HAstore/api"
	"fmt"
	"log"
)

const PORT = 5500

func main() {
	server := api.NewServer(PORT)

	fmt.Println("Starting server at ", fmt.Sprintf("http://localhost:%d/", PORT))
	log.Fatal(server.Start())
}
