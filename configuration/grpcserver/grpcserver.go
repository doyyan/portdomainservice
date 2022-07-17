package grpcserver

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/doyyan/portdomainservice/infrastructure/grpc/proto/ports"
	"github.com/doyyan/portdomainservice/interfaces/controller"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	ports.UnimplementedCreatePortsServer
	ports.UnimplementedUpdatePortsServer
	controller controller.PortController
}

func NewServer(c controller.PortController) *server {
	return &server{controller: c}
}

func (s *server) CreatePortsServer(context.Context, *ports.CreateOrUpdatePortsRequest) (*ports.CreateOrUpdatePortsResponse, error) {
	s.controller.CreatePorts()
	return &ports.CreateOrUpdatePortsResponse{Message: "done"}, nil
}

func (s *server) UpdatePortsServer(context.Context, *ports.CreateOrUpdatePortsRequest) (*ports.CreateOrUpdatePortsResponse, error) {
	s.controller.UpdatePorts()
	return &ports.CreateOrUpdatePortsResponse{Message: "done"}, nil
}

func NewGRPCServer(c controller.PortController) {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Ports Update/Create service to the server
	ports.RegisterUpdatePortsServer(s, server{})
	ports.RegisterCreatePortsServer(s, server{})
	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Ports server
	err = ports.RegisterUpdatePortsHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
