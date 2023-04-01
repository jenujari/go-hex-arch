package ports

import (
	"context"
	"go-hex-arch/internal/adapters/framework/left/grpc/pb"
	"net/http"
)

type GRPCPort interface {
	Run() error
	GetAddition(ctx context.Context, req *pb.OprationParameters) (*pb.Answer, error)
	GetSubtraction(ctx context.Context, req *pb.OprationParameters) (*pb.Answer, error)
	GetMultiplication(ctx context.Context, req *pb.OprationParameters) (*pb.Answer, error)
	GetDivision(ctx context.Context, req *pb.OprationParameters) (*pb.Answer, error)
}

type HTTPPort interface {
	Run() error
	GetAddition(w http.ResponseWriter, r *http.Request)
	GetSubtraction(w http.ResponseWriter, r *http.Request)
	GetMultiplication(w http.ResponseWriter, r *http.Request)
	GetDivision(w http.ResponseWriter, r *http.Request)
}
