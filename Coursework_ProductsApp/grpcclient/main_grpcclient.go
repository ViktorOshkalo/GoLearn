package main

import (
	"context"
	"fmt"
	"log"
	"main/configuration"
	gc "main/grpcconnection"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("Yoo GRPC Client")

	configuration.Setup()

	addr := fmt.Sprintf("%s:%s", configuration.GrpcHost, configuration.GrpcServerPort)

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := gc.NewProductsClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	request := gc.Request{
		CatalogId: 1,
		Filter: &gc.Filter{
			Properties: map[string]string{
				"Color": "Black",
				"Size":  "M",
			},
		},
	}

	r, err := c.GetProductsByCatalogIdWithFilter(ctx, &request)
	if err != nil {
		log.Fatalf("could not answer: %v", err)
	}
	log.Printf("Products: %v", r.GetProducts())
}
