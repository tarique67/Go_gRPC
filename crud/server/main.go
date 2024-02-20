package main

import (
	"context"
	pb "crud/proto"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "postgres"
)

type userServiceServer struct {
	db *sql.DB
	pb.UnimplementedUserServiceServer
}

func (s *userServiceServer) SaveUser(ctx context.Context, req *pb.SaveUserRequest) (*pb.SaveUserResponse, error) {
	token := uuid.New().String()
	_, err := s.db.Exec("INSERT INTO users (username, password, token) VALUES ($1, $2, $3)", req.Username, req.Password, token)
	if err != nil {
		return nil, err
	}
	return &pb.SaveUserResponse{Token: token}, nil
}
func (s *userServiceServer) GetUserDetails(ctx context.Context, req *pb.UserDetailsRequest) (*pb.UserDetailsResponse, error) {
	var name string
	var age int32
	err := s.db.QueryRow("SELECT name, age FROM users WHERE token = $1", req.Token).Scan(&name, &age)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &pb.UserDetailsResponse{Name: name, Age: age}, nil
}
func (s *userServiceServer) SaveUserDetails(ctx context.Context, req *pb.SaveUserDetailRequest) (*pb.SavedUserDetailResponse, error) {
	_, err := s.db.Exec("UPDATE users SET name = $1, age = $2 WHERE token = $3", req.Name, req.Age, req.Token)
	if err != nil {
		return nil, err
	}
	return &pb.SavedUserDetailResponse{Success: true}, nil
}
func (s *userServiceServer) UpdateUserName(ctx context.Context, req *pb.UpdateUserNameRequest) (*pb.UpdateUserNameResponse, error) {
	_, err := s.db.Exec("UPDATE users SET name = $1 WHERE token = $2", req.NewName, req.Token)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserNameResponse{Success: true}, nil
}
func main() {
	DbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", DbInfo)
	if err != nil {
		log.Fatalf("failed to connect to PostgreSQL: %v", err)
	}
	defer db.Close()
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &userServiceServer{db: db})
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("gRPC server started on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
