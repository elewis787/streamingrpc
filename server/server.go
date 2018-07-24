package server

import (
	"io"
	"log"

	"github.com/elewis787/streamingrpc/pb"
)

var _ pb.TransportServer = &Server{}

// Server - todo
type Server struct {
}

// Stream - todo
func (s *Server) Stream(stream pb.Transport_StreamServer) error {
	// check my map on peers
	//peerAddress := get address from ip (stream.Context())
	// if ok := map[peeraddress] { return nil }

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Println("Server recieved : ", in)
		in.Payload = append(in.Payload, []byte("-server")...)
		if err := stream.Send(in); err != nil {
			return err
		}
	}
}
