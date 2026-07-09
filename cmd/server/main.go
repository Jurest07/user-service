package main

import (
	"fmt"
	"net"
	"os"

	"github.com/Jurest07/user-service/internal/config"
	"github.com/Jurest07/user-service/internal/database"
	"github.com/Jurest07/user-service/internal/grpchandler"
	"github.com/Jurest07/user-service/internal/logger/handlers/slogpretty"
	"github.com/Jurest07/user-service/internal/repository"
	pb "github.com/Jurest07/user-service/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Ошибка загрузки конфигурации: %s", err.Error())
		os.Exit(1)
	}
	logger := slogpretty.SetupLogger(cfg.Env)
	err = database.Connect(cfg)
	if err != nil {
		logger.Error(fmt.Sprintf("Ошибка подключения к БД: %v", err.Error()))
		os.Exit(1)
	}
	defer database.Close()
	err = database.RunMigrations(cfg)
	if err != nil {
		logger.Error(fmt.Sprintf("Ошибка подключения к миграциям: %v", err.Error()))
		os.Exit(1)
	}
	userRepo := repository.NewUserRepo(database.DB)
	grpcUserHandler := grpchandler.NewUserHandler(userRepo, logger)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPCPort))
	if err != nil {
		logger.Error(fmt.Sprintf("Ошибка создания listener: %v", err.Error()))
		os.Exit(1)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, grpcUserHandler)
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		logger.Error("Ошибка запуска сервера", "error", err)
		os.Exit(1)
	}
}
