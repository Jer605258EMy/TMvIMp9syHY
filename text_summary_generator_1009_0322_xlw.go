// 代码生成时间: 2025-10-09 03:22:22
// text_summary_generator.go
// This GRPC service provides a text summary generator functionality.

package main

import (
    "context"
    "fmt"
    "log"
    "net"
    "os"
    "os/signal"
    "syscall"
    "time"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "google.golang.org/protobuf/types/known/emptypb"

    "path/to/your/pb/textsummary"  // Import your generated protobuf package
)

// server defines the GRPC server.
type server struct {
    textsummary.UnimplementedTextSummaryServiceServer
}

// SummarizeText takes the text and returns a summary.
func (s *server) SummarizeText(ctx context.Context, in *textsummary.SummarizeTextRequest) (*textsummary.SummarizeTextResponse, error) {
    // Perform the text summarization logic here.
    // For now, just returning a placeholder response.
    summary := "This is a summary of the text provided."
    return &textsummary.SummarizeTextResponse{Summary: summary}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Server listening on port 50051")

    grpcServer := grpc.NewServer()
    textsummary.RegisterTextSummaryServiceServer(grpcServer, &server{})
    reflection.Register(grpcServer)

    // Graceful shutdown on interrupt.
    go func() {
        sigs := make(chan os.Signal, 1)
        signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
        <-sigs
        log.Println("Shutting down gRPC server...")
        grpcServer.GracefulStop()
    }()

    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}