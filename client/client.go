package main

import (
	"errors"
	"fmt"
	"io"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	hello "github.com/fukushimaf929/grpc-sample/hello"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := hello.NewHelloServiceClient(conn)

	if err := SayHelloStream(c); err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
}

func SayHello(c hello.HelloServiceClient) error {
	req := &hello.Message{
		Body: "Hello From Client! SayHello",
	}
	res, err := c.SayHello(context.Background(), req)
	if err != nil {
		return err
	}
	fmt.Println(res.Body)
	return err
}

func SayHelloStream(c hello.HelloServiceClient) error {
	req := &hello.Message{
		Body: "Hello From Client! SayHelloStream",
	}
	stream, err := c.SayHelloStream(context.Background(), req)
	if err != nil {
		return err
	}

	for {
		res, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("メッセージ全てオックタ")
			break
		}
		if err != nil {
			return err
		}
		fmt.Println(res.Body)
	}
	return nil
}
