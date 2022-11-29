package ports

import (
	"context"

	"github.com/chaz8080/hex/internal/adapters/framework/driving/grpc/pb"
)

type GRPCPort interface {
	Run()
	GetAddition(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error)
	GetSubtraction(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error)
	GetMultiplication(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error)
	GetDivision(ctx context.Context, req *pb.OperationParameters) (*pb.Answer, error)
}
