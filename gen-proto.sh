#!/bin/bash
protoc --go_out=. --go-grpc_out=. proto/orders.proto
go mod tidy