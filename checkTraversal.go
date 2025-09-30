// This file just contains the function for travelsal blocking
// It is used by other functions like Create, Delete, etc.
// It is separated to avoid circular dependencies.
// Returns nicely lookin fmt errors
// Returns true if it is a traversal, false otherwise

package fileutil

import (
	"fmt"
	"os"
	"path/filepath"
)

// isPathTraversal checks if the given path is a traversal outside the base directory.
func checkTraversal(baseDir, targetPath string) (bool, error) {
	// Get the absolute paths
	absBaseDir, err := filepath.Abs(baseDir)
	if err != nil {
		return false, fmt.Errorf("invalid base directory: %w", err)
	}

	absTargetPath, err := filepath.Abs(targetPath)
	if err != nil {
		return false, fmt.Errorf("invalid target path: %w", err)
	}

	// Check if the target path is within the base directory
	relPath, err := filepath.Rel(absBaseDir, absTargetPath)
	if err != nil {
		return false, fmt.Errorf("error checking path relation: %w", err)
	}

	// If the relative path starts with "..", it's a traversal
	if relPath == ".." || relPath == "." || (len(relPath) > 3 && relPath[:3] == ".."+string(os.PathSeparator)) {
		return true, nil
	}

	return false, nil
}
