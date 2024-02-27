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
	log.Printf("\n========Testing Offers Methods========\n")
	log.Printf("\n\nTesting List all Offers \n")
	r1, err := client.ListAllOffers(ctx, &pb.ListOffersRequest{})
	if err != nil {
		log.Fatalf("could not list offers: %v", err)
	}
	log.Printf("RESULT: %v", r1.Offers)

	log.Printf("\n\nTesting query offers \n")
	log.Printf("Querying for offers  with CID 123")
	r2, err := client.QueryOffers(ctx, &pb.QueryOffersRequest{CID: "123"})
	if err != nil {
		log.Fatalf("could not query offers: %v", err)
	}
	log.Printf("RESULT: %v", r2.Offers)

	log.Printf("\n\nTesting Post Offer \n")
	newOffer := &pb.Boffer{
		CID: "TESTINGOFFER",
		IP: "255.255.255.255",
		Port: 420,
		Price: 9999,
	}
	r3, err := client.PostOffer(ctx, &pb.PostOfferRequest{Offer: newOffer})
	if err != nil {
		log.Fatalf("could not post offer %v, %v", r3, err)
	}
	log.Printf("Added the offer: %v", newOffer)
	log.Printf("Querying for this offer with CID... %v", newOffer.CID)

	r4, err := client.QueryOffers(ctx, &pb.QueryOffersRequest{CID: newOffer.CID})
	if err != nil {
		log.Fatalf("could not query newly posted offer: %v", err)
	}
	log.Printf("RESULT: %v", r4.Offers)

	log.Printf("\n========Testing Bid Methods========\n")
	log.Printf("\n\nTesting List all bids \n")
	r5, err := client.ListBids(ctx, &pb.ListBidRequest{})
	if err != nil {
		log.Fatalf("could not list bids: %v", err)
	}
	log.Printf("RESULT: %v", r5.Bids)

	log.Printf("\n\nTesting query bids \n")
	log.Printf("Querying for bids with CID f1e4c2")
	r6, err := client.QueryBids(ctx, &pb.QueryBidsRequest{CID: "f1e4c2"})
	if err != nil {
		log.Fatalf("could not query bids: %v", err)
	}
	log.Printf("RESULT: %v", r6.Bids)

	log.Printf("\n\nTesting Post Bid \n")
	newBid := &pb.Boffer{
		CID: "TESTINGBID",
		IP: "0.0.0.0",
		Port: 69,
		Price: 96,
	}
	r7, err := client.PostBid(ctx, &pb.PostBidRequest{Bid: newBid})
	if err != nil {
		log.Fatalf("could not post bid %v, %v", r7, err)
	}
	log.Printf("Added the bid: %v", newBid)
	log.Printf("Querying for this offer with CID... %v", newBid.CID)

	r8, err := client.QueryBids(ctx, &pb.QueryBidsRequest{CID: newBid.CID})
	if err != nil {
		log.Fatalf("could not query newly posted offer: %v", err)
	}
	log.Printf("RESULT: %v", r8.Bids)

}