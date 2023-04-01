package main

import (
	"fmt"
	"go-hex-arch/internal/adapters/app/api"
	"go-hex-arch/internal/adapters/core/arithmatic"
	gRPC "go-hex-arch/internal/adapters/framework/left/grpc"
	http "go-hex-arch/internal/adapters/framework/left/http"
	"go-hex-arch/internal/adapters/framework/right/db"
	"go-hex-arch/internal/ports"
	"log"
	"os"

	"golang.org/x/sync/errgroup"
)

func main() {
	var err error

	// ports
	var arithAdapter ports.ArithmaticPort
	var dbAdapter ports.DbPort
	var appAdapter ports.APIPort
	var grpcAdapter ports.GRPCPort
	var httpAdapter ports.HTTPPort

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

	httpAdapter = http.NewAdapter(appAdapter)

	g := new(errgroup.Group)

	g.Go(func() error {
		return grpcAdapter.Run()
	})

	g.Go(func() error {
		return httpAdapter.Run()
	})

	if err := g.Wait(); err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Println("Successfully executed")
}
