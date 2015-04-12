package main

/*
Author : Devicharan

*/
import (
    "google.golang.org/grpc"
    "log"
    pb "basicwebapp/proto"
    "golang.org/x/net/context"
)

const (
    serverAddress = "localhost:50053"
)
func main() {

    var opts []grpc.DialOption

    conn, err := grpc.Dial(serverAddress, opts...)

    if err !=nil {
        log.Fatal("fail to dial %v", err)
    }

    defer conn.Close()

    client := pb.NewRecommendationServiceClient(conn)


    stream, err := client.GetRecommendations(context.Background(), &pb.Product{Name:"Mouse", Id:1, Vendor:"Microsoft"})

    if err !=nil {
        log.Fatal("Failed getting product %v", err)
    }

    log.Println(stream)



}
