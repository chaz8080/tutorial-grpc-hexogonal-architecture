package main

import (
	"log"
	"os"

	"github.com/chaz8080/hex/internal/adapters/app/api"
	"github.com/chaz8080/hex/internal/adapters/core/arithmetic"
	"github.com/chaz8080/hex/internal/adapters/framework/driven/db"
	"github.com/chaz8080/hex/internal/adapters/framework/driving/grpc"
	"github.com/chaz8080/hex/internal/ports"
)

func main() {
	var err error

	// driving ports
	//   domain layer
	var coreAdapter ports.ArithemticPort
	//   application layer
	var appAdapter ports.APIPort
	//   framework layer
	var gRPCAdapter ports.GRPCPort

	// driven ports
	var dbaseAdapter ports.DbPort

	DB_DRIVER := os.Getenv("DB_DRIVER")
	DS_NAME := os.Getenv("DS_NAME")

	dbaseAdapter, err = db.NewAdapter(DB_DRIVER, DS_NAME)
	if err != nil {
		log.Fatalf("failed to get db adapter: %v", err)
	}
	defer dbaseAdapter.CloseDbConnection()

	coreAdapter = arithmetic.NewAdapter()
	appAdapter = api.NewAdapter(dbaseAdapter, coreAdapter)

	gRPCAdapter = grpc.NewAdapter(appAdapter)
	gRPCAdapter.Run()
}
