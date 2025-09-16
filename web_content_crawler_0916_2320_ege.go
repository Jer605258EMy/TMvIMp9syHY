// 代码生成时间: 2025-09-16 23:20:40
package main

import (
    "context"
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
    "time"
)

// CrawlerService defines the service for web content crawling
type CrawlerService struct{}

// Crawl takes a URL and returns the raw HTML content of the webpage
func (s *CrawlerService) Crawl(ctx context.Context, url string) (string, error) {
    // Start a timer to measure the request duration
    startTime := time.Now()
    defer func() {
        fmt.Printf("Request took %s
", time.Since(startTime))
    }()
    
    // Make the HTTP GET request to the provided URL
    resp, err := http.Get(url)
    if err != nil {
        return "", fmt.Errorf("error fetching URL: %w", err)
    }
    defer resp.Body.Close()
    
    // Check if the response status code is not 200 OK
    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("bad status: %d", resp.StatusCode)
    }
    
    // Read the body of the response
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", fmt.Errorf("error reading response body: %w", err)
    }
    
    // Convert the body to a string
    htmlContent := string(body)
    
    // Return the raw HTML content of the webpage
    return htmlContent, nil
}

func main() {
    // Define the URL to crawl
    url := "http://example.com"

    // Create an instance of the CrawlerService
    crawler := CrawlerService{}

    // Call the Crawl method to fetch the webpage content
    htmlContent, err := crawler.Crawl(context.Background(), url)
    if err != nil {
        fmt.Printf("Error: %s
", err)
        return
    }

    // Print the raw HTML content
    fmt.Println("Raw HTML Content:")
    fmt.Println(htmlContent)
}
