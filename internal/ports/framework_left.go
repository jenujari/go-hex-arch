package ports

import (
	"context"
	"go-hex-arch/internal/adapters/framework/left/grpc/pb"
)

type GRPCPort interface {
	Run()
	GetAddition(ctx context.Context, req *pb.OprationParameters) (*pb.Answer, error)
	GetSubtraction(ctx context.Context, req *pb.OprationParameters) (*pb.Answer, error)
	GetMultiplication(ctx context.Context, req *pb.OprationParameters) (*pb.Answer, error)
	GetDivision(ctx context.Context, req *pb.OprationParameters) (*pb.Answer, error)
}
