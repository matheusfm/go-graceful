package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	server := &http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			log.Println("handler: Received the request")
			time.Sleep(3 * time.Second)
			log.Println("handler: Fulfilling the request after 3 seconds")
			fmt.Fprint(rw, "Hello World!")
		}),
	}

	// Start server
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("main: Failed to start server: %v", err)
		}
	}()
	log.Println("main: Server started")

	// Wait for SIGINT or SIGTERM signal to gracefully shutdown the server with a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	s := <-quit
	log.Printf("main: Signal received: %v", s)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("main: Failed to gracefully shutdown: %v", err)
	}
	log.Println("main: Server was shutdown gracefully")
}
