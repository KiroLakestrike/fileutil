# fileutil

A simple, platform-independent Go library for safe file creation and utility functions.

## Features

- Safely create new files in a specified directory
- Prevents directory traversal and invalid file names
- Checks for existing files before creation
- Returns clear error messages
- Platform independent (works on Linux, macOS, and Windows)

## Installation

Add to your Go project using:

```sh
go get github.com/KiroLakestrike/fileutil
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/KiroLakestrike/fileutil"
)

func main() {
    ok, err := fileutil.Create("/tmp", "config.yaml")
    if ok {
        fmt.Println("File created successfully!")
    } else {
        fmt.Printf("Failed to create file: %v\n", err)
    }
}
```

### func Create(path, fileName string) (bool, error)

Creates a new empty file with the given name under the given path.

- Returns `(true, nil)` if the file was created successfully.
- Returns `(false, error)` if the file already exists, the name is invalid, or another error occurs.

## Platform Notes

- File permissions (`0644`) are set on Unix-like systems. On Windows, permissions are ignored and default to the current user.
- Directory traversal and invalid file names are blocked for security.

## License

MIT

## Author

KiroLakestrike