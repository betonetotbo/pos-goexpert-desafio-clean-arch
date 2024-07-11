#!/bin/bash
protoc --go_out=.. --go-grpc_out=.. --proto_path=.. proto/orders.proto
go mod tidy