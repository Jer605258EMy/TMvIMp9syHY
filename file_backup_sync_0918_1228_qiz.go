// 代码生成时间: 2025-09-18 12:28:06
package main

import (
    "context"
    "fmt"
    "google.golang.org/grpc"
    "io"
    "io/ioutil"
    "log"
    "net"
    "os"
    "path/filepath"
    "strings"
)

// FileBackupSyncService defines the service that handles file backup and sync operations
type FileBackupSyncService struct {
    // Add any fields if necessary
}

// NewFileBackupSyncService creates a new instance of FileBackupSyncService
func NewFileBackupSyncService() *FileBackupSyncService {
    return &FileBackupSyncService{}
}

// BackupFile is a method to backup a given file
func (s *FileBackupSyncService) BackupFile(ctx context.Context, in *BackupRequest) (*BackupResponse, error) {
    // Implement file backup logic here
    // For simplicity, we are just copying the file to a backup directory
    sourcePath := in.GetFilePath()
    backupPath := filepath.Join("./backup", filepath.Base(sourcePath))

    // Check if the file exists before attempting to backup
    if _, err := os.Stat(sourcePath); os.IsNotExist(err) {
        return nil, fmt.Errorf("file not found: %w", err)
    }

    // Create backup directory if it doesn't exist
    if _, err := os.Stat("./backup"); os.IsNotExist(err) {
        if err := os.MkdirAll("./backup", 0755); err != nil {
            return nil, fmt.Errorf("failed to create backup directory: %w", err)
        }
    }

    // Copy file to backup directory
    if err := copyFile(sourcePath, backupPath); err != nil {
        return nil, fmt.Errorf("failed to backup file: %w", err)
    }

    return &BackupResponse{Success: true}, nil
}

// SyncFiles is a method to sync files between two directories
func (s *FileBackupSyncService) SyncFiles(ctx context.Context, in *SyncRequest) (*SyncResponse, error) {
    // Implement file sync logic here
    sourcePath := in.GetSourcePath()
    targetPath := in.GetTargetPath()

    // Perform file sync operations
    if err := syncFiles(sourcePath, targetPath); err != nil {
        return nil, fmt.Errorf("failed to sync files: %w", err)
    }

    return &SyncResponse{Success: true}, nil
}

// copyFile copies a file from source to destination
func copyFile(src, dst string) error {
    sourceFile, err := os.Open(src)
    if err != nil {
        return err
