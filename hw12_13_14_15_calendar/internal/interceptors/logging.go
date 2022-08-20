package interceptors

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

const dateLayout = "2006-01-02"

// LoggingInterceptor ...
func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	fmt.Printf("[gRPC] %s: %s --- %v\n", time.Now().Format(dateLayout), info.FullMethod, req)

	res, err := handler(ctx, req)
	if err != nil {
		fmt.Printf("[gRPC] %s: %s --- %v\n", time.Now().Format(dateLayout), info.FullMethod, err)
		return nil, err
	}

	fmt.Printf("[gRPC] %s: %s --- %v\n", time.Now().Format(dateLayout), info.FullMethod, res)

	return res, nil
}
