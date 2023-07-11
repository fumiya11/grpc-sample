package main

import (
	"fmt"
	"log"
	"net"

	hello "github.com/fumiya11/grpc-sample/hello"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Start GRPC Serve")
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()

	hello.RegisterHelloServiceServer(s, &hello.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
