package main

import (
	"go-hex-arch/internal/adapters/app/api"
	"go-hex-arch/internal/adapters/core/arithmatic"
	gRPC "go-hex-arch/internal/adapters/framework/left/grpc"
	"go-hex-arch/internal/adapters/framework/right/db"
	"go-hex-arch/internal/ports"
	"log"
	"os"
)

func main() {
	var err error

	// ports
	var arithAdapter ports.ArithmaticPort
	var dbAdapter ports.DbPort
	var appAdapter ports.APIPort
	var grpcAdapter ports.GRPCPort

	dbDriver := os.Getenv("DB_DRIVER")
	dbSource := os.Getenv("DB_DATASOURCE")

	dbAdapter, err = db.NewAdapter(dbDriver, dbSource)

	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}

	defer dbAdapter.CloseDbConnection()

	arithAdapter = arithmatic.NewAdapter()

	appAdapter = api.NewAdapter(dbAdapter, arithAdapter)
	grpcAdapter = gRPC.NewAdapter(appAdapter)

	grpcAdapter.Run()
}
