package main

/*
Author : Devicharan

*/

import (
    pb "basicwebapp/proto"
    "golang.org/x/net/context"
    "net"
    "log"
    "google.golang.org/grpc"


)
const (

    port = ":50053"
)

type recommendationserviceserver struct {

}

func (s *recommendationserviceserver) GetRecommendations(context.Context, *pb.Product) (*pb.RecommendationResponse, error) {

    return &pb.RecommendationResponse{[]*pb.RecommendationResponse_Recommendation{&pb.RecommendationResponse_Recommendation{Rating:4, Productid:1}}}, nil

}

func main() {

    lis, err := net.Listen("tcp", port)

    if err != nil {

        log.Fatal("failed to listen : %v", err)
    }

    s := grpc.NewServer()

    pb.RegisterRecommendationServiceServer(s, &recommendationserviceserver{})


    s.Serve(lis)

}
