你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
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
 
