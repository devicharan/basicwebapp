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
    serverAddress = "localhost:50052"
)
func main() {

    var opts []grpc.DialOption

    conn, err := grpc.Dial(serverAddress, opts...)

    if err !=nil {
        log.Fatal("fail to dial %v", err)
    }

    defer conn.Close()

    client := pb.NewCatalogServiceClient(conn)


    stream, err := client.GetProductCatalog(context.Background(), &pb.CatalogRequest{&pb.Category{CategoryName:"Computer"}})

    if err !=nil {
        log.Fatal("Failed getting product %v", err)
    }

    log.Println(stream)



}
