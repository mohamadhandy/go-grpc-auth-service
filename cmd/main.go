package main

import (
	"go-grpc-auth-svc/pkg/config"
	"go-grpc-auth-svc/pkg/db"
	"go-grpc-auth-svc/pkg/pb"
	"go-grpc-auth-svc/pkg/services"
	"go-grpc-auth-svc/pkg/utils"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	h := db.Init(c.DBUrl)
	jwt := utils.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "go-grpc-auth-svc",
		ExpirationHours: 24 * 365,
	}
	listen, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("Failed to listing", err)
	}

	s := services.Server{
		H:   h,
		Jwt: jwt,
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
