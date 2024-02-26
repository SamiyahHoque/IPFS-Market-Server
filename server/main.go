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

	// get all properties from the parameters the user passed
	addedBoffer := in.GetOffer()

	// add it to the appropiate CID slice
	boffers := offerTable[addedBoffer.GetCID()] 
	boffers = append(boffers, boffer {
		IP:    addedBoffer.IP,
		Port:  addedBoffer.Port,
		Price: addedBoffer.Price,
	})

	return &pb.PostOfferResponse{}, nil
}
func (s *server) listAllOffers (ctx context.Context, in *pb.ListOffersRequest) (*pb.ListOffersResponse, error) {

	// slice of pointers of type Boffer. we will store ALL offers here
	var allOffers []*pb.Boffer

	// go into every single CID in offerTable map, in which each CID is linked to a slice of boffers
	for CID := range offerTable {

		for _, boffer := range offerTable[CID] {

			// create a new boffer element with all properties from the boffer we are looking at
			// after that, add to the allOffers array
			allOffers = append(allOffers, &pb.Boffer {
				IP: boffer.IP,
				Port: boffer.Port,
				Price: boffer.Price,
			})

		}

	}

	return &pb.ListOffersResponse{Offers: allOffers}, nil

}
func (s *server) queryBids (ctx context.Context, in *pb.QueryBidsRequest) (*pb.QueryBidsResponse, error) {
	/* To work on:
		I just copied the format of queryOffers(). Still need to test if it works
	*/

	cid := in.CID

    boffers, exists := bidTable[cid]
	if !exists || len(boffers) == 0{
        return &pb.QueryBidsResponse{}, nil
    }

	var offers []*pb.Boffer
    for _, boffer := range boffers {
		offers = append(offers, &pb.Boffer{
            IP:    boffer.IP,
            Port:  boffer.Port,
            Price: boffer.Price,
        })
	}

    return &pb.QueryBidsResponse{Offers: offers}, nil
}

func (s *server) postBid (ctx context.Context, in *pb.PostBidRequest) (*pb.PostBidResponse, error) {
	/* To work on:
		Because its a server and we may be getting simultaneous requests from clients, we should handle updates to
		the bidTable in a more formal manner (similar to CSE 320 hw 4/5 - process or thread based multitasking)
	*/

	addedBoffer := in.Offer

	boffers, exists := bidTable[addedBoffer.CID]
	if !exists || len(boffers) == 0{
        return &pb.PostBidResponse{}, nil
    }

    boffers = append(boffers, boffer{
        IP:    addedBoffer.IP,
        Port:  addedBoffer.Port,
        Price: addedBoffer.Price,
    })

    bidTable[addedBoffer.GetCID()] = boffers

    return &pb.PostBidResponse{}, nil
}

func (s *server) listBids(ctx context.Context, in *pb.ListBidRequest) (*pb.ListBidResponse, error) {
    var allBids []*pb.Boffer

    for _, bids := range bidTable {
        for _, bid := range bids {
            allBids = append(allBids, &pb.Boffer{
                IP:    bid.IP,
                Port:  bid.Port,
                Price: bid.Price,
            })
        }
    }

    return &pb.ListBidResponse{Offers: allBids}, nil
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
	fillDummyBidData()

	//print data 
	log.Println("Offers:")
    for cid, offers := range offerTable {
        log.Printf("CID: %s, Total Offers: %d", cid, len(offers))
        for i, offer := range offers {
            log.Printf("\tOffer %d: IP: %s, Port: %d, Price: %d", i+1, offer.IP, offer.Port, offer.Price)
        }
    }
    log.Println("Bids:")
    for cid, bids := range bidTable {
        log.Printf("CID: %s, Total Bids: %d", cid, len(bids))
        for i, bid := range bids {
            log.Printf("\tBid %d: IP: %s, Port: %d, Price: %d", i+1, bid.IP, bid.Port, bid.Price)
        }
    }

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


func fillDummyBidData() {
    for key := range bidTable {
        delete(bidTable, key)
    }
    bidTable["123"] = []boffer{
        {"192.168.1.100", 8080, 95},
        {"10.10.10.10", 3030, 115},
        {"192.168.2.200", 9090, 105},
    }
    bidTable["abc"] = []boffer{
        {"10.0.0.2", 80, 35},
        {"192.168.100.101", 3000, 45},
        {"172.16.0.2", 7070, 55},
    }
    bidTable["f1e4c2"] = []boffer{
        {"10.20.30.40", 8080, 200},
        {"192.168.1.1", 8080, 210},
        {"172.16.0.1", 443, 190},
        {"10.30.40.50", 4040, 185},
    }
    bidTable["456def"] = []boffer{
        {"192.168.3.3", 3030, 120},
        {"10.0.2.2", 2020, 130},
    }
    bidTable["789ghi"] = []boffer{
        {"192.168.4.4", 4040, 140},
        {"10.0.3.3", 3030, 150},
        {"172.16.1.1", 5050, 160},
    }
}
