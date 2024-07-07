package main

import (
	"context"
	"log"
	"testing"
	"time"

	pb "grpc/proto" // Update with your actual protobuf package name

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = ":50051"
)

func TestGetUserById(t *testing.T) {

	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for id := int32(1); id <= 10; id++ {
		testGetUserById(ctx, t, client, id)
	}
}

func TestGetUsersList(t *testing.T) {
 
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

 
	client := pb.NewUserServiceClient(conn)

 
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

 
	testGetUsersList(ctx, t, client)
}

func TestSearchByCriteria(t *testing.T) {
 
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

 
	client := pb.NewUserServiceClient(conn)

 
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

 
	testSearchByCriteria(ctx, t, client)
}

func testGetUserById(ctx context.Context, t *testing.T, client pb.UserServiceClient, id int32) {
	req := &pb.GetUserByIdRequest{Id: id}

	resp, err := client.GetUserById(ctx, req)
	if err != nil {
		t.Fatalf("GetUserById failed for ID %d: %v", id, err)
	}

	log.Printf("GetUserById Response for ID %d: %v", id, resp)
 
	if resp.User == nil {
		t.Errorf("Expected non-nil user for ID %d, but got nil", id)
	}
	if resp.User.Id != id {
		t.Errorf("Expected user ID %d, but got %d", id, resp.User.Id)
	}
}

func testGetUsersList(ctx context.Context, t *testing.T, client pb.UserServiceClient) {
	req := &pb.GetUsersListRequest{Ids: []int32{1, 2, 3}}

	resp, err := client.GetUsersList(ctx, req)
	if err != nil {
		t.Fatalf("GetUsersList failed: %v", err)
	}

	log.Printf("GetUsersList Response: %v", resp)

 
	expectedNumUsers := len(req.Ids)
	if len(resp.Users) != expectedNumUsers {
		t.Errorf("Expected %d users, but got %d", expectedNumUsers, len(resp.Users))
	}
	for i, user := range resp.Users {
		if user.Id != req.Ids[i] {
			t.Errorf("Expected user ID %d at index %d, but got %d", req.Ids[i], i, user.Id)
		}
	}
}

func testSearchByCriteria(ctx context.Context, t *testing.T, client pb.UserServiceClient) {
	req := &pb.SearchByCriteriaRequest{City: "LA", IsMarried: true}

	resp, err := client.SearchByCriteria(ctx, req)
	if err != nil {
		t.Fatalf("SearchByCriteria failed: %v", err)
	}

	log.Printf("SearchByCriteria Response: %v", resp)

 
	for _, user := range resp.Users {
		if user.City != req.City {
			t.Errorf("Expected user City %s, but got %s", req.City, user.City)
		}
		if user.Married != req.IsMarried {
			t.Errorf("Expected user Married %v, but got %v", req.IsMarried, user.Married)
		}
	}
}
