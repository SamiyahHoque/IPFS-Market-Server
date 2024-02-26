package main

import(
	"context"
	"log"
	// "os"
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

	//make time out for this client to limit 
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//testing server method calls
	log.Printf("\n\nTesting List all Offers \n")
	r1, err := client.ListAllOffers(ctx, &pb.ListOffersRequest{})
	if err != nil {
		log.Fatalf("could not list offers: %v", err)
	}
	log.Printf("RESULT: %v", r1.Offers)

	log.Printf("\n\nTesting query offers \n")
	r2, err := client.QueryOffers(ctx, &pb.QueryOffersRequest{CID: "123"})
	if err != nil {
		log.Fatalf("could not query offers: %v", err)
	}
	log.Printf("RESULT: %v", r2.Offers)

	log.Printf("\n\nTesting Post Offer \n")
	newBoffer := &pb.Boffer{
		CID: "TESTING",
		IP: "255.255.255.255",
		Port: 420,
		Price: 9999,
	}
	r3, err := client.PostOffer(ctx, &pb.PostOfferRequest{Offer: newBoffer})
	if err != nil {
		log.Fatalf("could not post offer %v, %v", r3, err)
	}
	log.Printf("Added the offer: %v", newBoffer)
	log.Printf("Querying for this offer with CID... %v", newBoffer.CID)

	r4, err := client.QueryOffers(ctx, &pb.QueryOffersRequest{CID: "TESTING"})
	if err != nil {
		log.Fatalf("could not query newly posted offer: %v", err)
	}
	log.Printf("RESULT: %v", r4.Offers)


}