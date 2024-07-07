package main

import (
	"context"
	"log"
	"time"

	pb "grpc/proto" // Update with your actual protobuf package name

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Dial the gRPC server with insecure transport credentials
	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	// Create a client instance
	client := pb.NewUserServiceClient(conn)

	// Context with a timeout of 1 second
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Example: GetUserById RPC call
	userByID(ctx, client, 1)

	// Example: GetUsersList RPC call
	userListByID(ctx, client, []int32{1, 2, 3})

	// Example: SearchByCriteria RPC call
	searchUsers(ctx, client, "LA", true)
}

// userByID fetches a user by ID using GetUserById RPC
func userByID(ctx context.Context, client pb.UserServiceClient, userID int32) {
	req := &pb.GetUserByIdRequest{Id: userID}
	res, err := client.GetUserById(ctx, req)
	if err != nil {
		log.Fatalf("GetUserById failed: %v", err)
	}
	log.Printf("GetUserById Response: %v", res)
}

// userListByID fetches users by a list of IDs using GetUsersList RPC
func userListByID(ctx context.Context, client pb.UserServiceClient, userIDs []int32) {
	req := &pb.GetUsersListRequest{Ids: userIDs}
	res, err := client.GetUsersList(ctx, req)
	if err != nil {
		log.Fatalf("GetUsersList failed: %v", err)
	}
	log.Printf("GetUsersList Response: %v", res)
}

// searchUsers searches users by criteria using SearchByCriteria RPC
func searchUsers(ctx context.Context, client pb.UserServiceClient, city string, isMarried bool) {
	req := &pb.SearchByCriteriaRequest{City: city, IsMarried: isMarried}
	res, err := client.SearchByCriteria(ctx, req)
	if err != nil {
		log.Fatalf("SearchByCriteria failed: %v", err)
	}
	log.Printf("SearchByCriteria Response: %v", res)
}
