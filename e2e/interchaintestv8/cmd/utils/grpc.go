package utils

import (
	"fmt"
	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

func GetTLSGRPC(addr string) (*grpc.ClientConn, error) {
	creds := insecure.NewCredentials()

	// Establish a secure connection with the gRPC server
	conn, err := grpc.Dial(addr, grpc.
		WithTransportCredentials(creds))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to grpc client with addr: %s: %w", addr, err)
	}

	return conn, nil
}
