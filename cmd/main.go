package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"net"
	"testProject/internal/api"
	"testProject/internal/repository"
	"testProject/internal/service"
	pb "testProject/protos/gen/go"

	"log"
)

func main() {
	db, err := sql.Open("sqlite3", "storage/database.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	repo := &repository.UserRepo{DB: db}

	userService := &service.UserService{Repo: repo}

	apiService := &api.UserService{Service: userService}

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, apiService)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	if err := grpcServer.Serve(listener); err != nil {
		log.Println("Grpc server error:", err)
	}

	fmt.Println("Server started successfully!")

}
