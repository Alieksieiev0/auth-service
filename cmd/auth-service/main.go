package main

import (
	"flag"
	"log"

	"github.com/Alieksieiev0/auth-service/internal/services"
	"github.com/Alieksieiev0/auth-service/internal/transport/grpc"
	"github.com/Alieksieiev0/auth-service/internal/transport/rest"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/sync/errgroup"
)

func main() {
	var (
		restServerAddr = flag.String("rest-server", ":3001", "listen address of rest server")
		grpcServerAddr = flag.String("grpc-server", ":4001", "listen address of grpc server")
		grpcClientAddr = flag.String("grpc-client", ":4000", "listen address of grpc client")
		tokenService   = services.NewTokenService()
		app            = fiber.New()
		g              = new(errgroup.Group)
	)

	client, err := grpc.NewGRPCClient(*grpcClientAddr)
	if err != nil {
		log.Fatal(err)
	}

	authService := services.NewAuthService(client, tokenService)

	grpcServer := grpc.NewServer()
	g.Go(func() error {
		return grpcServer.Start(*grpcServerAddr, authService)
	})

	restServer := rest.NewServer(app)
	g.Go(func() error {
		return restServer.Start(*restServerAddr, authService)
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
