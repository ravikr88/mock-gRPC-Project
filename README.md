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

## Building the Application

To build the Docker image for this application:

```bash
docker build -t grpc-user-service .
```

To run the container

```bash
docker run -d -p 50051:50051 grpc-user-service
```

To test the server.go

```bash
go test -v ./...
```

```
Replace placeholders like `path/to/your/proto/package` with the actual path to your protobuf package.
```
