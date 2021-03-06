package client

import (
	"io"
	"log"

	"github.com/elewis787/streamingrpc/pb"
	"golang.org/x/net/context"
)

// Client todo
type Client struct {
	Transport pb.TransportClient
}

// Stream - todo
func (c *Client) Stream(ctx context.Context, msg []byte) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			stream, err := c.Transport.Stream(ctx)
			if err != nil {
				return err
			}
			defer stream.CloseSend()
			if err := stream.Send(&pb.Packet{Payload: msg}); err != nil {
				return err
			}

			in, err := stream.Recv()
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return err
			}
			log.Println("Hello Client got : ", in)
		}
	}
}
