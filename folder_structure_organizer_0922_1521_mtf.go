// 代码生成时间: 2025-09-22 15:21:53
package main

import (
    "fmt"
    "io/fs"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "sort"
)

// FolderStructure represents the structure of a folder.
type FolderStructure struct {
    Folders []string `json:"folders"`
    Files   []string `json:"files"`
}

// OrganizeFolderStructure generates a FolderStructure for the given directory path.
func OrganizeFolderStructure(path string) (*FolderStructure, error) {
    var folders []string
    var files []string
    var err error

    // List all items in the directory.
    items, err := ioutil.ReadDir(path)
    if err != nil {
        return nil, err
    }

    // Sort items to ensure consistency in the output.
    sort.Slice(items, func(i, j int) bool {
        return items[i].Name() < items[j].Name()
    })

    for _, item := range items {
        switch {
        case item.IsDir():
            folders = append(folders, item.Name())
        default:
            files = append(files, item.Name())
        }
    }

    return &FolderStructure{Folders: folders, Files: files}, nil
}

// PrintFolderStructure prints the folder structure to the console.
func PrintFolderStructure(fs *FolderStructure) {
    fmt.Println("Folders:")
    for _, folder := range fs.Folders {
        fmt.Println(folder)
    }
    fmt.Println("Files:")
    for _, file := range fs.Files {
        fmt.Println(file)
    }
}

func main() {
    path := "./" // Set the path to the directory you want to organize.
    fs, err := OrganizeFolderStructure(path)
    if err != nil {
        log.Fatalf("Error organizing folder structure: %v", err)
    }

    PrintFolderStructure(fs)
}
