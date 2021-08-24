package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ory/graceful"
)

func main() {
	server := graceful.WithDefaults(&http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			log.Println("handler: Received the request")
			time.Sleep(3 * time.Second)
			log.Println("handler: Fulfilling the request after 3 seconds")
			fmt.Fprint(rw, "Hello World!")
		}),
	})

	log.Println("main: Starting the server")
	if err := graceful.Graceful(server.ListenAndServe, server.Shutdown); err != nil {
		log.Fatalln("main: Failed to gracefully shutdown")
	}
	log.Println("main: Server was shutdown gracefully")
}
