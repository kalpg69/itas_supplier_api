package app

import (
	"log"
	"net"

	supplierpb "github.com/kalpg69/itas_supplier_api/api"
	"github.com/kalpg69/itas_supplier_api/controller"
	"github.com/kalpg69/itas_supplier_api/database/mysqlclient"
	"github.com/kalpg69/itas_supplier_api/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartApplication(address string) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("failed to listen")
	}

	s := grpc.NewServer()
	srv := controller.NewController(service.NewService(mysqlclient.NewMySQLClient("itasAdmin", "Itas@123", "localhost", "supplierdb")))
	supplierpb.RegisterSupplierServer(s, srv)
	reflection.Register(s)
	if err = s.Serve(lis); err != nil {
		log.Fatal("failed to serve")
	}
}
