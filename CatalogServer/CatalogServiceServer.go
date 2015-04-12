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

    port = ":50052"
)
/*

   This microservice brings in catalog services..

*/
type catalogserviceserver struct {}

func (s *catalogserviceserver) GetProductCatalog(ctx context.Context, in *pb.CatalogRequest) (*pb.CatalogResponse, error) {
    //CatalogService.Product{Name:"Mouse",ID:12}
    /*
       TODO: You can process the CatalogRequest here ... I just hardcoded response..

       TODO: You can connect to a database a retrieve data here

    */
    products := []*pb.Product{
        &pb.Product{Name:"Mouse", Id:1, Vendor:"Microsoft"},
        &pb.Product{Name:"Keyboard", Id:2, Vendor:"LG"},
        &pb.Product{Name:"Monitor", Id:3, Vendor:"Samsung"}}

    return &pb.CatalogResponse{Products: products}, nil
}

func main() {

    lis, err := net.Listen("tcp", port)

    if err != nil {

        log.Fatal("failed to listen : %v", err)
    }

    s := grpc.NewServer()

    pb.RegisterCatalogServiceServer(s, &catalogserviceserver{})

    s.Serve(lis)

}
