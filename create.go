// Create creates a new empty file at the specified path.
// Returns true and nil if successful, false and error otherwise.
package fileutil

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Create creates a new empty file with the given name under the given path.
// If the file already exists, it returns false and an error.
// Returns true and nil if the file was created successfully, false and error otherwise.
// The function is platform independent.

func Create(path, fileName string) (bool, error) {
	// Prevent directory traversal with the check traversal function

	isTraversal, err := checkTraversal(path, filepath.Join(path, fileName))
	if err != nil {
		return false, fmt.Errorf("error checking path traversal: %w", err)
	}
	// Also ensure the fileName does not contain any path separators
	if isTraversal || strings.Contains(fileName, string(os.PathSeparator)) {
		return false, fmt.Errorf("invalid file name: %v", fileName)
	}

	fullPath := filepath.Join(path, fileName)

	// Convert to absolute path
	absPath, err := filepath.Abs(fullPath)
	if err != nil {
		return false, fmt.Errorf("invalid path: %w", err)
	}

	// Check if the file already exists
	if _, err := os.Stat(absPath); err == nil {
		return false, fmt.Errorf("file %v already exists", fileName)
	} else if !os.IsNotExist(err) {
		return false, fmt.Errorf("error checking if file %v exists: %w", fileName, err)
	}

	// Create the new file (0644 is platform independent for regular files)
  // Windows will ignore it though might add support for ACL someday.
  // but that looks like a nightmare
	file, err := os.OpenFile(absPath, os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0644)
	if err != nil {
		return false, fmt.Errorf("error creating file %v: %w", fileName, err)
	}
	defer file.Close()

	return true, nil
}
