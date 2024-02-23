package main

import (
	"log"
	"net"
	"context"

	pb "marketServer"
	"google.golang.org/grpc"
)

// Adjust the import path
type server struct {
	pb.UnimplementedMarketServer
}

// the (s *server) binds the function to the server
func (s *server) queryOffers (ctx context.Context, in *pb.QueryOffersRequest) (*pb.QueryOffersResponse, error) {
	cid := in.GetCID()

    boffers, exists := offerTable[cid]
	if !exists || len(boffers) == 0{
        return &pb.QueryOffersResponse{}, nil
    }

	var offers []*pb.Boffer
    for _, boffer := range boffers {
		offers = append(offers, &pb.Boffer{
            IP:    boffer.IP,
            Port:  boffer.Port,
            Price: boffer.Price,
        })
	}

    return &pb.QueryOffersResponse{Offers: offers}, nil
}
func (s *server) postOffer (ctx context.Context, in *pb.PostOfferRequest) (*pb.PostOfferResponse, error) {
	return &pb.PostOfferResponse{}, nil
}
func (s *server) listAllOffers (ctx context.Context, in *pb.ListOffersRequest) (*pb.ListOffersResponse, error) {
	return &pb.ListOffersResponse{}, nil
}
func (s *server) queryBids (ctx context.Context, in *pb.QueryBidsRequest) (*pb.QueryBidsResponse, error) {
	return &pb.QueryBidsResponse{}, nil
}
func (s *server) postBid (ctx context.Context, in *pb.PostBidRequest) (*pb.PostBidResponse, error) {
	return &pb.PostBidResponse{}, nil
}
func (s *server) listBids (ctx context.Context, in *pb.ListBidRequest) (*pb.ListBidResponse, error) {
	return &pb.ListBidResponse{}, nil
}


type boffer struct{
	IP 		string
	Port 	int32
	Price 	int32
}

var offerTable = make(map[string][]boffer)
var bidTable = make(map[string][]boffer)

func main(){
	clearAndFillDummyData()
	
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to set up listening port: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMarketServer(s, &server{})

	log.Printf("Server Listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	
}

func clearAndFillDummyData(){
	for key := range offerTable{
		delete(offerTable, key)
	}

	offerTable["123"] = []boffer{
		{"192.168.0.1", 80, 100},
		{"172.100.01.12", 3030, 120},
		{"192.168.0.2", 80, 90},
	}
	offerTable["abc"] = []boffer{
		{"255.255.0.1", 80, 30},
		{"169.200.021.122", 3000, 40},
		{"100.33.30.21", 8888, 50},
		{"100.33.30.21", 8080, 60},
		{"100.33.30.21", 2525, 70},
	}

	offerTable["f1e4c2"] = []boffer{
		{"10.20.30.40", 8080, 200},
		{"192.168.1.1", 8080, 210},
		{"172.16.0.1", 443, 190},
	}
	
	offerTable["200bec4f00"] = []boffer{
		{"10.0.0.1", 80, 25},
		{"192.168.100.100", 80, 35},
		{"10.10.10.10", 22, 45},
	}
}
