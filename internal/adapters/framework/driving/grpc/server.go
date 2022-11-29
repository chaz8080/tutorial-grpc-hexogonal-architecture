package grpc

import (
	"log"
	"net"

	"github.com/chaz8080/hex/internal/adapters/framework/driving/grpc/pb"
	"github.com/chaz8080/hex/internal/ports"

	"google.golang.org/grpc"
)

type Adapter struct {
	api ports.APIPort
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

func (grpca Adapter) Run() {
	port := ":9000"
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", port, err)
	}

	arithmaticServiceServer := grpca
	grpcServer := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(grpcServer, arithmaticServiceServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc server: %v", err)
	}
}
