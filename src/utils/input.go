package utils

import (
  "bufio"
  "os"
)

func LoadInput(path string) []string {
  // Get file handle
  file, err := os.Open(path)
  if err != nil {
    panic(err)
  }

  // Get scanner object with attached file
  // Configure scanner to split by newline
  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)

  // Scan through file and add each line as new slice item
  var inputs []string
  for scanner.Scan() {
    inputs = append(inputs, scanner.Text())
  }

  // Close the file handle
  file.Close()

  return inputs
}
