// 代码生成时间: 2025-09-18 00:57:41
package main

import (
    "context"
    "fmt"
    "log"
    "net"
    "os"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
)

// InventoryService 定义了一个库存服务
type InventoryService struct {
    products map[int]*Product // 产品存储
}

// Product 描述了一个库存产品
type Product struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
    Quantity    int    `json:"quantity"`
}

// AddProduct 向库存中添加新产品
func (s *InventoryService) AddProduct(ctx context.Context, in *Product) (*Response, error) {
    if in.ID < 1 || in.Quantity < 1 || in.Name == "" {
        return nil, fmt.Errorf("invalid product details")
    }
    s.products[in.ID] = in
    return &Response{Success: true, Message: "Product added successfully"}, nil
}

// GetProduct 根据ID获取库存中的产品
func (s *InventoryService) GetProduct(ctx context.Context, in *Request) (*Product, error) {
    product, exists := s.products[in.ID]
    if !exists {
        return nil, fmt.Errorf("product not found")
    }
    return product, nil
}

// UpdateProduct 更新库存中的产品信息
func (s *InventoryService) UpdateProduct(ctx context.Context, in *Product) (*Response, error) {
    if _, exists := s.products[in.ID]; !exists {
        return nil, fmt.Errorf("product not found")
    }
    s.products[in.ID] = in
    return &Response{Success: true, Message: "Product updated successfully"}, nil
}

// DeleteProduct 从库存中删除产品
func (s *InventoryService) DeleteProduct(ctx context.Context, in *Request) (*Response, error) {
    if _, exists := s.products[in.ID]; !exists {
        return nil, fmt.Errorf("product not found\)
    }
    delete(s.products, in.ID)
    return &Response{Success: true, Message: "Product deleted successfully"}, nil
}

// Response 是一个通用响应结构
type Response struct {
    Success bool   `json:"success"`
    Message string `json:"message"`
}

// Request 是一个通用请求结构
type Request struct {
    ID int `json:"id"`
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Server is running on :50051")

    grpcServer := grpc.NewServer()
    // 注册服务到服务器
    inventory.RegisterInventoryServiceServer(grpcServer, &InventoryService{products: make(map[int]*Product)})
    reflection.Register(grpcServer)

    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// 以下是gRPC相关的代码，需要在proto文件中定义相应的服务
// inventory.proto
// syntax = "proto3";
// package inventory;
// service InventoryService {
//     rpc AddProduct(Product) returns (Response);
//     rpc GetProduct(Request) returns (Product);
//     rpc UpdateProduct(Product) returns (Response);
//     rpc DeleteProduct(Request) returns (Response);
// }
// message Product {
//     int32 id = 1;
//     string name = 2;
//     string description = 3;
//     int32 quantity = 4;
// }
// message Request {
//     int32 id = 1;
// }
// message Response {
//     bool success = 1;
//     string message = 2;
// }
