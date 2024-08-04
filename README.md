# Golang gRPC User Service

This project implements a gRPC-based user service in Golang, providing functionalities to manage user details and perform searches based on specific criteria.

## Table of Contents

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Building the Application](#building-the-application)
- [Running the Application](#running-the-application)
- [Accessing gRPC Service Endpoints](#accessing-grpc-service-endpoints)
- [Testing](#testing)

## Overview

The gRPC User Service simulates a database by maintaining user details in memory. It provides gRPC endpoints for fetching user details by ID, retrieving user details by a list of IDs, and searching for users based on specific criteria.

## Prerequisites

Before running the application, ensure you have the following installed:

- Docker
- Go (if building from source)


## Run the server.go from root of directory
``` bash
go run server.go
```

## Run the client from mock-gRPC-Project/client/client.go
```bash
go run client.go
```
 

## Testing
``` bash
go run server_test.go
```
 

## Accessing gRPC Service Endpoints

```bash
package main

import (
    "context"
    "log"
    "time"

    "google.golang.org/grpc"
    pb "path/to/your/proto/package"
)

func main() {
    // Create a connection to the server
    conn, err := grpc.Dial(":50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Failed to dial server: %v", err)
    }
    defer conn.Close()

    client := pb.NewUserServiceClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    // Example: Sending GetUserByIdRequest
    getUserByIdRequest := &pb.GetUserByIdRequest{Id: 1}
    getUserByIdResponse, err := client.GetUserById(ctx, getUserByIdRequest)
    if err != nil {
        log.Fatalf("GetUserById failed: %v", err)
    }
    log.Printf("GetUserById Response: %v", getUserByIdResponse)

    // Example: Sending GetUsersListRequest
    getUsersListRequest := &pb.GetUsersListRequest{Ids: []int32{1, 2, 3}}
    getUsersListResponse, err := client.GetUsersList(ctx, getUsersListRequest)
    if err != nil {
        log.Fatalf("GetUsersList failed: %v", err)
    }
    log.Printf("GetUsersList Response: %v", getUsersListResponse)

    // Example: Sending SearchByCriteriaRequest
    searchByCriteriaRequest := &pb.SearchByCriteriaRequest{
        City:      "LA",
        IsMarried: true,
    }
    searchByCriteriaResponse, err := client.SearchByCriteria(ctx, searchByCriteriaRequest)
    if err != nil {
        log.Fatalf("SearchByCriteria failed: %v", err)
    }
    log.Printf("SearchByCriteria Response: %v", searchByCriteriaResponse)
}

```
