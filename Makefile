build:
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/jittakal/go-micro-sample/ pkg/echo/service/echo.proto
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/jittakal/go-micro-sample/ pkg/blog/service/user.proto
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/jittakal/go-micro-sample/ pkg/blog/service/article.proto
