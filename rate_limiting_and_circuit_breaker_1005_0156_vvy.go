// 代码生成时间: 2025-10-05 01:56:24
package main

import (
    "context"
    "fmt"
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/metadata"
    "google.golang.org/grpc/status"

    "github.com/juju/ratelimit"
    "github.com/sony/gobreaker"
)

// Define the service we want to protect with circuit breaker
type Service interface {
    Call(ctx context.Context) error
}

// Implement the Service interface
type MyService struct {
    limit ratelimit.Limit
    window ratelimit.Window
    bucket *ratelimit.Bucket
    breaker *gobreaker.CircuitBreaker
}

func NewMyService() *MyService {
    // 10 requests per second
    limit := ratelimit.Limit(10)
    window := ratelimit.NewWindow(1 * time.Second)
    bucket := ratelimit.NewBucket(limit, window)

    // Circuit breaker configuration
    breaker := gobreaker.NewCircuitBreaker(gobreaker.Settings{
        Name:    "myService",
        Timeout: 10 * time.Second,
        Trip: func(counts gobreaker.Counts) bool {
            return counts.ConsecutiveFailures >= 3
        },
    })

    return &MyService{
        limit:  limit,
        window: window,
        bucket: bucket,
        breaker: breaker,
    }
}

func (m *MyService) Call(ctx context.Context) error {
    if !m.breaker.Allow() {
        return status.Errorf(codes.Unavailable, "Service is currently unavailable")
    }

    // If the token is not taken, it means we've reached the rate limit
    if !m.bucket.Take(1) {
        return status.Errorf(codes.ResourceExhausted, "Rate limit exceeded")
    }

    // Simulate a service call that might fail
    if rand.Intn(10) < 2 { // 20% chance to fail
        return fmt.Errorf("service call failed")
    }

    // Service call is successful
    return nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    service := NewMyService()
    // Register the service with the gRPC server
    grpcServer.Serve(lis)
    log.Printf("Server listening on %v", lis.Addr())
    
    // Run the server
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
