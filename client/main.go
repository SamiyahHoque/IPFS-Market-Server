package main

import(
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	pb "marketServer" 
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	address := "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	
	if err != nil {		
		log.Fatalf("Failed to connect %v", err)
	}

	defer conn.Close()

	client := pb.NewMarketClient(conn)
	name := "world"

	//allows you to change name in program args
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//call server method
	r, err := client.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("Could not say hello :( %v", err)
	}

	log.Printf("Greeting: %s", r.GetMessage())

}