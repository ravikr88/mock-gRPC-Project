package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpc/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
	users []*pb.User
}

func (s *server) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
	for _, user := range s.users {
		if user.Id == req.Id {
			return &pb.GetUserByIdResponse{User: user}, nil
		}
	}
	return nil, fmt.Errorf("user with ID %d not found", req.Id)
}

func (s *server) GetUsersList(ctx context.Context, req *pb.GetUsersListRequest) (*pb.GetUsersListResponse, error) {
	var users []*pb.User
	for _, id := range req.Ids {
		found := false
		for _, user := range s.users {
			if user.Id == id {
				users = append(users, user)
				found = true
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("user with ID %d not found", id)
		}
	}
	return &pb.GetUsersListResponse{Users: users}, nil
}

func (s *server) SearchByCriteria(ctx context.Context, req *pb.SearchByCriteriaRequest) (*pb.SearchByCriteriaResponse, error) {
	var usersList []*pb.User

	for _, user := range s.users {
		if (req.City == user.City) &&
			(req.IsMarried == user.Married) {
			log.Printf("Found matching user: %v", user.Id)
			usersList = append(usersList, user)
		}
	}

	if len(usersList) == 0 {
		log.Printf("No users found matching criteria: City=%s, IsMarried=%v", req.City, req.IsMarried)
	}

	return &pb.SearchByCriteriaResponse{Users: usersList}, nil
}

func main() {
	users := []*pb.User{
		{Id: 1, FName: "Alice", City: "NY", Phone: 1234567890, Height: 5.8, Married: false},
		{Id: 2, FName: "Bob", City: "LA", Phone: 9876543210, Height: 6.0, Married: true},
		{Id: 3, FName: "Carol", City: "CHI", Phone: 5556667777, Height: 5.5, Married: true},
		{Id: 4, FName: "David", City: "SF", Phone: 1112223333, Height: 5.10, Married: false},
		{Id: 5, FName: "Eve", City: "SEA", Phone: 9998887777, Height: 5.7, Married: false},
		{Id: 6, FName: "Frank", City: "DEN", Phone: 3334445555, Height: 6.1, Married: true},
		{Id: 7, FName: "Grace", City: "MIA", Phone: 6667778888, Height: 5.9, Married: false},
		{Id: 8, FName: "Helen", City: "BOS", Phone: 2223334444, Height: 5.4, Married: true},
		{Id: 9, FName: "Ian", City: "AUS", Phone: 7778889999, Height: 6.2, Married: false},
		{Id: 10, FName: "Julia", City: "POR", Phone: 8889990000, Height: 5.6, Married: true},
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{users: users})

	log.Printf("Server listening on %s", lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
