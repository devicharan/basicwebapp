#Microservices for golang webapp using grpc

#Prerequisite

The best way to start on this is going through https://github.com/grpc/grpc-common/tree/master/go

You should be able to understand grpc in go after this.

1. run catalogservier

   go build CatalogServiceServer.go //Execute this inside Catalogserver directory
   
   run the exe generated in the same directory
   
2. run recommendation server 

   go build RecommendationServiceServer.go
   
   run the exe generated in the same directory
   
3. Compile the main.go file and run it to start webserver.

4. Open localhost:8080 to view a bootstrap page . Click on submit button to see responses coming from two different microservices
 
