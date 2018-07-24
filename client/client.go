package client

import (
	"errors"
	"io"
	"log"

	"github.com/elewis787/streamingrpc/pb"
	"golang.org/x/net/context"
)

// Client todo
type Client struct {
}

// Stream - todo
func (c *Client) Stream(ctx context.Context, transport pb.TransportClient, msg []byte) error {
	stream, err := transport.Stream(ctx)
	if err != nil {
		return nil
	}

	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Println(err)
				return
			}
			log.Println("Client got : ", in)
		}
	}()

	go func() {
		for {
			p := &pb.Packet{
				Payload: msg,
			}
			if err := stream.Send(p); err != nil {
				log.Println(err)
				return
			}
		}
	}()
	select {
	case <-ctx.Done():
		stream.CloseSend()
		return errors.New("context closed")
	}
}
