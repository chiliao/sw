package server

import (
	"net"

	examples "github.com/pensando/grpc-gateway/examples/examplepb"
	"google.golang.org/grpc"
)

func Run() error {
	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	examples.RegisterEchoServiceServer(s, newEchoServer())
	examples.RegisterFlowCombinationServer(s, newFlowCombinationServer())

	abe := newABitOfEverythingServer()
	examples.RegisterABitOfEverythingServiceServer(s, abe)
	examples.RegisterStreamServiceServer(s, abe)

	s.Serve(l)
	return nil
}
