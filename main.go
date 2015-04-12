package main

/*
Author : Devicharan

*/
import (
    "net/http"
// "github.com/russross/blackfriday"
    "log"
    pb "basicwebapp/proto"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "encoding/json"
)

const (
    recommendationServeraddress = "localhost:50053"
    catalogServerAddress = "localhost:50052"
    defaultName = "world"
)
func main() {


    http.HandleFunc("/catalog", GetProductCatalog)
    http.Handle("/", http.FileServer(http.Dir("public")))
    http.ListenAndServe(":8080", nil)

}

func GetProductCatalog(rw http.ResponseWriter, r *http.Request) {

    var opts []grpc.DialOption

    conn, err := grpc.Dial(catalogServerAddress, opts...)

    if err !=nil {
        log.Fatal("fail to dial %v", err)
    }

    defer conn.Close()

    client := pb.NewCatalogServiceClient(conn)


    stream, err := client.GetProductCatalog(context.Background(), &pb.CatalogRequest{&pb.Category{CategoryName:"Computer"}})

    if err !=nil {
        log.Fatal("Failed getting product %v", err)
    }
    result, err := json.Marshal(stream)
    rw.Write(result)

    // Make call to Recommendations MicroService now...
    GetRecommendations(rw, r)

}

func GetRecommendations(rw http.ResponseWriter, r *http.Request) {

    var opts []grpc.DialOption

    conn, err := grpc.Dial(recommendationServeraddress, opts...)

    if err !=nil {
        log.Fatal("fail to dial %v", err)
    }

    defer conn.Close()

    client := pb.NewRecommendationServiceClient(conn)


    stream, err := client.GetRecommendations(context.Background(), &pb.Product{Name:"Mouse", Id:1, Vendor:"Microsoft"})

    if err !=nil {
        log.Fatal("Failed getting recommendations %v", err)
    }
    result, err := json.Marshal(stream)
    rw.Write(result)
}
