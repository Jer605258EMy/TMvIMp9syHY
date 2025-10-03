// 代码生成时间: 2025-10-03 18:12:11
package main

import (
    "context"
    "fmt"
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

// ModelTrainingProgress is the message type for training progress updates.
type ModelTrainingProgress struct {
    Epoch int32 `protobuf:"varint,1,opt,name=epoch,proto3" json:"epoch,omitempty"`
    // Add more fields as needed, e.g., accuracy, loss, etc.
}

// TrainingMonitorService is the server API for ModelTrainingMonitor service.
type TrainingMonitorService struct{}

// StreamTrainingProgress is a streaming RPC method that receives updates about the training progress.
func (s *TrainingMonitorService) StreamTrainingProgress(stream TrainingMonitor_StreamTrainingProgressServer) error {
    for {
        msg, err := stream.Recv()
        if err != nil {
            if err == io.EOF {
                return nil // end of stream
            }
            return status.Errorf(codes.InvalidArgument, "failed to receive message: %v", err)
        }
        // Log the received message or handle it as needed
        fmt.Printf("Received training progress: Epoch %d
", msg.GetEpoch())
    }
}

// Register and start the GRPC server.
func startServer() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Server listening on port 50051")

    grpcServer := grpc.NewServer()
    // Register the service with the server
    RegisterTrainingMonitorServer(grpcServer, &TrainingMonitorService{})
    grpcServer.Serve(lis)
}

func main() {
    startServer()
}
