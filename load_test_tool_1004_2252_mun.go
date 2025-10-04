// 代码生成时间: 2025-10-04 22:52:51
// load_test_tool.go
package main

import (
    "fmt"
    "log"
    "math/rand"
    "os"
    "os/signal"
    "syscall"
    "time"

    "golang.org/x/net/context"
    "google.golang.org/grpc"
)

// Define a constant for the server address.
const (
    serverAddress = "localhost:50051"
)

// Define the LoadTestConfig structure to hold the load test configuration.
type LoadTestConfig struct {
    Concurrency int
    Duration    time.Duration
}

// Client is used to interact with the gRPC server.
type Client struct {
    conn *grpc.ClientConn
    client LoadTesterClient
}

// NewClient creates a new gRPC client connected to the server.
func NewClient(address string) (*Client, error) {
    conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        return nil, err
    }
    client := NewLoadTesterClient(conn)
    return &Client{conn, client}, nil
}

// StartLoadTest starts the load test with the given configuration.
func (c *Client) StartLoadTest(config LoadTestConfig) {
    // Create a context with a timeout for the load test duration.
    ctx, cancel := context.WithTimeout(context.Background(), config.Duration)
    defer cancel()

    // Create a channel to receive the results of the load test.
    results := make(chan *LoadTestResponse, config.Concurrency)

    // Start the load test with the specified concurrency.
    for i := 0; i < config.Concurrency; i++ {
        go func() {
            for {
                select {
                case <-ctx.Done():
                    return
                default:
                    // Call the gRPC method to perform the load test.
                    response, err := c.client.LoadTest(ctx, &LoadTestRequest{})
                    if err != nil {
                        log.Printf("Error during load test: %v", err)
                    } else {
                        // Send the result to the channel.
                        results <- response
                    }
                    // Introduce a random delay to simulate variable load.
                    time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
                }
            }
        }()
    }

    // Wait for the load test to finish or be interrupted.
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        <-sigChan
        cancel()
    }()

    // Process the load test results.
    for result := range results {
        fmt.Printf("Received response: %s\
", result.GetMessage())
    }
}

func main() {
    // Define the load test configuration.
    config := LoadTestConfig{
        Concurrency: 10, // Number of concurrent goroutines.
        Duration:    5 * time.Minute, // Duration of the load test.
    }

    // Create a new gRPC client.
    client, err := NewClient(serverAddress)
    if err != nil {
        log.Fatalf("Failed to create gRPC client: %v", err)
    }
    defer client.conn.Close()

    // Start the load test.
    client.StartLoadTest(config)
}
