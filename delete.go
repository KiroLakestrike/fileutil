package fileutil

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Delete will be used to delete files in the future

// Check for folowing errors: 
//  insufficient permissions, path is a symlink

// The main purpose of this file, is to delete a given file at the specified path. 
// Can be used to get rid of temporary files for examples, during atomic copying for example.
// Will return true and nil if everything worked well, or false and err, if there was an error.

func Delete(path, fileName string) (bool, error) {
	// Prevent directory traversal and invalid file names
	if strings.Contains(fileName, "..") || strings.ContainsAny(fileName, `/\`) {
		return false, fmt.Errorf("invalid file name")
	}

	fullPath := filepath.Join(path, fileName)

	// Convert to absolute path
	absPath, err := filepath.Abs(fullPath)
	if err != nil {
		return false, fmt.Errorf("invalid path: %w", err)
	}

	// Check if the Path is a symlink
	fileInfo, err := os.Lstat(absPath)
	if err != nil {
		return false, fmt.Errorf("error checking if file %v is a symlink: %w", fileName, err)
	}
	if fileInfo.Mode()&os.ModeSymlink != 0 {
		return false, fmt.Errorf("path %v is a symlink", absPath)
	}

	// Check if the file exists
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return false, fmt.Errorf("file %v does not exist", fileName)
	} else if err != nil {
		return false, fmt.Errorf("error checking if file %v exists: %w", fileName, err)
	}

	// Check if the path is a directory
	fileInfo, err = os.Stat(absPath)
	if err != nil {
		return false, fmt.Errorf("error getting file info for %v: %w", fileName, err)
	}
	if fileInfo.IsDir() {
		return false, fmt.Errorf("path %v is a directory", absPath)
	}

	// Check for write permissions
	file, err := os.OpenFile(absPath, os.O_WRONLY, 0)
	if err != nil {
		if os.IsPermission(err) {
			return false, fmt.Errorf("insufficient permissions to delete file %v: %w", fileName, err)
		}
		return false, fmt.Errorf("error opening file %v for writing: %w", fileName, err)
	}
	file.Close()
	
	// Attempt to delete the file
	err = os.Remove(absPath)
	if err != nil {
		return false, fmt.Errorf("error deleting file %v: %w", fileName, err)
	}

	return true, nil
}