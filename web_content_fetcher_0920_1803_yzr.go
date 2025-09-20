// 代码生成时间: 2025-09-20 18:03:08
package main

import (
    "context"
    "fmt"
    "io/ioutil"
    "net/http"
    "log"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

// WebContentFetcherService is a service for fetching web content via gRPC
type WebContentFetcherService struct{}

// FetchWebContent is a gRPC method to fetch the content of a webpage
func (s *WebContentFetcherService) FetchWebContent(ctx context.Context, req *FetchWebContentRequest) (*FetchWebContentResponse, error) {
    // Check for a valid URL
    if req.Url == "" {
        return nil, status.Errorf(codes.InvalidArgument, "empty URL")
    }

    // Fetch the webpage content
    resp, err := http.Get(req.Url)
    if err != nil {
        return nil, status.Errorf(codes.Internal, "failed to fetch web content: %v", err)
    }
    defer resp.Body.Close()

    // Read the body of the response
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, status.Errorf(codes.Internal, "failed to read response body: %v", err)
    }

    // Return the fetched content
    return &FetchWebContentResponse{Content: string(body)}, nil
}

// FetchWebContentRequest is the request message for fetching web content
type FetchWebContentRequest struct {
    Url string
}

// FetchWebContentResponse is the response message for fetching web content
type FetchWebContentResponse struct {
    Content string
}

func main() {
    // Set up the gRPC server
    listener, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    server := grpc.NewServer()
    pb.RegisterWebContentFetcherServiceServer(server, &WebContentFetcherService{})

    // Start the server
    if err := server.Serve(listener); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// The protocol buffer definitions for the service would typically be in a separate file,
// but for simplicity, they are included as comments here.
//
// message FetchWebContentRequest {
//   string url = 1;
// }
//
// message FetchWebContentResponse {
//   string content = 1;
// }
//
// service WebContentFetcherService {
//   rpc FetchWebContent(FetchWebContentRequest) returns (FetchWebContentResponse);
// }