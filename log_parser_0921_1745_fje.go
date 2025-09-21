// 代码生成时间: 2025-09-21 17:45:42
package main

import (
    "context"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "strings"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    pb "path/to/your/protobuf/definitions" // Replace with the actual path to your protobuf definitions
)

// LogParserService defines the gRPC service for log parsing
type LogParserService struct {
    pb.UnimplementedLogParserServer

    // Include any internal state or dependencies required by the service
}

// ParseLog is called by clients to parse a log file
func (s *LogParserService) ParseLog(ctx context.Context, in *pb.ParseLogRequest) (*pb.ParseLogResponse, error) {
    // Verify that the log file path is provided
    if in.GetFilePath() == "" {
        return nil, status.Errorf(codes.InvalidArgument, "File path must be provided")
    }

    // Read the log file content
    fileContent, err := ioutil.ReadFile(in.GetFilePath())
    if err != nil {
        // Log the error and return a gRPC error with more context
        log.Printf("Error reading file: %v", err)
        return nil, status.Errorf(codes.Internal, "Failed to read the log file")
    }

    // Parse the log file content (This is a placeholder for actual parsing logic)
    // For simplicity, this example just returns the file content as is
    return &pb.ParseLogResponse{Data: string(fileContent)}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Log Parser Service is running on port 50051")

    // Create a new gRPC server
    s := grpc.NewServer()
    pb.RegisterLogParserServer(s, &LogParserService{})
    reflection.Register(s) // Enables gRPC reflection for service discovery

    // Start the server
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
