package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/saddam-satria/grpc/model/greet"
	"google.golang.org/grpc"
)

func main() {
	opts := []grpc.DialOption{
			grpc.WithInsecure(), 
	}

	conn, err :=  grpc.NewClient("localhost:5001",opts...)
	if err != nil {
		log.Fatalf("Could not connect to service: %v", err)
		return
	}
	defer conn.Close()

	greetClientService := greet.NewGreetClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := greetClientService.SayHello(ctx, &greet.Hello{Name:"saddam"})
	if err != nil {
		log.Fatalf("Could not create request: %v", err)
	}

	fmt.Println(response.Message)
}