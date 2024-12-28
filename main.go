package main

import (
	"context"
	"echo-base/database"
	"echo-base/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGKILL, syscall.SIGINT, syscall.SIGTERM)

	err := server.NewConfig()
	if err != nil {
		log.Println("Error load config")
		panic(err)
	}

	connector, err := database.NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	orm, err := connector.Open()
	if err != nil {
		log.Fatal(err)
	}

	address := os.Getenv(server.FlagAddress)
	if address == "" {
		address = server.DefaultAddress
	}

	go server.NewEchoEngine(orm, address).Serve()

	<-c

	log.Println("server shutdowns")

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	defer connector.Close()
}
