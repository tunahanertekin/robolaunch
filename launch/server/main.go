package main

import (
	"log"
	"net"

	launchPb "github.com/robolaunch/robolaunch/api"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	launchPb.UnimplementedLaunchServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	launchPb.RegisterLaunchServer(s, &server{})
	log.Printf("Launch server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
