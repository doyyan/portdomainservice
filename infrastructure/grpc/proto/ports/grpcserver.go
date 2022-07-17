package ports

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/doyyan/portdomainservice/interfaces/controller"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type createserver struct {
	controller controller.PortController
}

func (s *createserver) mustEmbedUnimplementedCreatePortsServer() {
	//TODO implement me
	panic("implement me")
}

func (s *createserver) mustEmbedUnimplementedUpdatePortsServer() {
	//TODO implement me
	panic("implement me")
}

type updateserver struct {
	controller controller.PortController
}

func (u updateserver) UpdatePorts(ctx context.Context, request *CreateOrUpdatePortsRequest) (*CreateOrUpdatePortsResponse, error) {
	u.controller.UpdatePorts()
	return &CreateOrUpdatePortsResponse{Message: "done"}, nil
}

func (u updateserver) mustEmbedUnimplementedUpdatePortsServer() {
	//TODO implement me
	panic("implement me")
}

func NewServer(c controller.PortController) *createserver {
	return &createserver{controller: c}
}

func (s *createserver) CreatePorts(context.Context, *CreateOrUpdatePortsRequest) (*CreateOrUpdatePortsResponse, error) {
	s.controller.CreatePorts()
	return &CreateOrUpdatePortsResponse{Message: "done"}, nil
}

func (s *createserver) UpdatePorts(context.Context, *CreateOrUpdatePortsRequest) (*CreateOrUpdatePortsResponse, error) {
	s.controller.UpdatePorts()
	return &CreateOrUpdatePortsResponse{Message: "done"}, nil
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
	RegisterUpdatePortsServer(s, &updateserver{controller: c})
	RegisterCreatePortsServer(s, &createserver{controller: c})
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
	err = RegisterUpdatePortsHandler(context.Background(), gwmux, conn)
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
