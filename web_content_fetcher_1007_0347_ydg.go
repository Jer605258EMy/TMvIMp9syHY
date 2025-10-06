// 代码生成时间: 2025-10-07 03:47:21
package main

import (
  "bufio"
  "bytes"
  "context"
  "fmt"
  "io"
  "io/ioutil"
  "net/http"
  "strings"
)

// WebContentFetcher defines the structure for fetching web content
type WebContentFetcher struct {
  Client *http.Client
}

// NewWebContentFetcher initializes a new WebContentFetcher
func NewWebContentFetcher() *WebContentFetcher {
  return &WebContentFetcher{
    Client: &http.Client{},
  }
}

// FetchContent fetches the content of a given URL
func (fetcher *WebContentFetcher) FetchContent(url string) (string, error) {
  // Send an HTTP request to the specified URL
  resp, err := fetcher.Client.Get(url)
  if err != nil {
    return "", err
  }
  defer resp.Body.Close()

  // Check if the response status code is successful
  if resp.StatusCode != http.StatusOK {
    return "", fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
  }

  // Read the response body
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return "", err
  }

  // Convert the body to a string and return it
  return string(body), nil
}

// Main function to demonstrate the usage of WebContentFetcher
func main() {
  fetcher := NewWebContentFetcher()
  url := "http://example.com"
  content, err := fetcher.FetchContent(url)
  if err != nil {
    fmt.Printf("Error fetching content: %v
", err)
  } else {
    fmt.Printf("Fetched content: %s
", content)
  }
}
