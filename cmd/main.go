package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"google.golang.org/grpc"
	"net"
	"testProject/internal/api"
	"testProject/internal/repository"
	"testProject/internal/service"
	pb "testProject/protos/gen/go"

	"log"
)

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=trxxlzz dbname=test_db sslmode=disable")
	if err != nil {
		log.Fatal("Cannot connect with database", err)
	}

	if err := goose.Up(db, "./migrations"); err != nil {
		log.Fatal("Cannot apply migration", err)
	}

	log.Println("Connected to database!")

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
