package __

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
)

type Server struct {
	UnimplementedHelloServiceServer
}

func (s *Server) SayHello(ctx context.Context, req *Message) (*Message, error) {
	fmt.Println(req.Body)
	return &Message{
		Body: "Hello From Server!",
	}, nil
}

func (s *Server) SayHelloStream(req *Message, stream HelloService_SayHelloStreamServer) error {
	fmt.Println(req.Body)
	for i := 0; i < 10; i++ {
		if err := stream.Send(&Message{
			Body: fmt.Sprintf("[%d] Hello From Server!, %s!", i, req.GetBody()),
		}); err != nil {
			return err
		}
		time.Sleep(time.Second * 2)
	}
	return nil
}
