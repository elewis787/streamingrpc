package server

import (
	"io"
	"log"

	"github.com/elewis787/streamingrpc/pb"
)

var _ pb.TransportServer = &Server{}

// HelloServer - todo
type HelloServer struct {
}

// Stream - todo
func (s *HelloServer) Stream(stream pb.Hello_StreamServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		log.Println("Hello Server recieved : ", in)
		in.Payload = append(in.Payload, []byte("-HelloServer")...)
		if err := stream.Send(in); err != nil {
			return err
		}
	}
}
