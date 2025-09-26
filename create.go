// Create new empty files with this
package fileutil

import(
  "os"
  "path/filepath"
  "fmt"
)

func create(path string, fileName string) (bool, error) {
  // Will create an empty file with the given extension under the given path.
  // If file already Exists, the function will stop with a error message "File already exists"
  // Returns true and nil, if the file ws created successfully, and false and error if the file was not created

  fullPath := filepath.Join(path, fileName)

  // Checks if the file already exists
  if _, err := os.Stat(fullPath); err == nil {
    err := fmt.Errorf("File %v already exists", fileName)
    return false, err
  } else if !os.IsNotExist(err) {
    err := fmt.Errorf("Another error occured when checking if the file %v exists", fileName)
    return false, err
  }

  // Create the new file
  file, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0644)
  if err != nil {
    err := fmt.Errorf("There was an error creating the file %v", fileName)
    return false, err
  }
  defer file.Close()

  // everything went well
  return true, nil
}
