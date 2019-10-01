package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/StevenRojas/go-dynamodb-demo/server"
)

func main() {
	address := flag.String("http", ":8080", "http listen address")

	flag.Parse()
	ctx := context.Background()
	srv := server.GetServiceInstance()
	errChan := make(chan error)

	endpoints := server.Endpoints{
		SchemaEndpoint:            server.MakeSchemaEndpoint(srv),
		SchemaDescriptionEndpoint: server.MakeSchemaDescriptionEndpoint(srv),
		AddThemeEndpoint:          server.MakeAddThemeEndpoint(srv),
		ListThemeEndpoint:         server.MakeListThemeEndpoint(srv),
		PatchThemeEndpoint:        server.MakePatchThemeEndpoint(srv),
		QueryThemeEndpoint:        server.MakeQueryThemeEndpoint(srv),
		GetThemeEndpoint:          server.MakeGetThemeEndpoint(srv),
	}

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		log.Println("Server is running on port: ", *address)
		handler := server.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*address, handler)
	}()

	log.Fatalln(<-errChan)

}
