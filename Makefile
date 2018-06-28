build:
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/jittakal/go-micro-sample/ pkg/echo/service/echo.proto
