package main

import (
	"fmt"
	"net"

	"github.com/santoshanand/sample-grpc-go/gen"
	"github.com/santoshanand/sample-grpc-go/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	startServer()
}

func startServer() {

	lis, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()

	// Register user service server
	gen.RegisterUserServiceServer(srv, &user.NewUserService{})

	reflection.Register(srv)

	serr := srv.Serve(lis)
	fmt.Println("Server init...")

	if serr != nil {
		panic(serr)
	}
}
