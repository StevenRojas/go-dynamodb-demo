package main

import(
	"fmt"
	"flag"
	"context"
	"os"
	"os/signal"
	"log"
	"syscall"
	"net/http"

	"github.com/StevenRojas/go-dynamodb-demo/server"
	"github.com/StevenRojas/go-dynamodb-demo/services"
)

func main() {
	var (
		address = flag.String("http", ":8080", "http listen address")
	)
	flag.Parse()
	ctx := context.Background()
	srv := server.NewServer()
	errorChan := make(chan error)

	endpoints := server.Endpoints{

	}

	go func() {
		c := make(chan os.Signal. 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errorChan <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		log.Println("Server is listen on port: ", *address)
		handler := server.NewHttpServer(ctx, endpoints)
		errorChan <- http.ListenAndServe(*address, handler)
	}()
	log.Fatalln(<-errorChan)
}