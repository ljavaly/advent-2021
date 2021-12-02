package main

import (
  "bytes"
  "fmt"
  "io/ioutil"
  "strconv"
)

func main() {
  dat, err := ioutil.ReadFile("input.txt")
  if err != nil {
    panic(err)
  }
  // fmt.Print(string(dat))

  depths := bytes.Split(dat, []byte("\n"))
  fmt.Println(fmt.Sprintf("Got %d depths as input", len(depths)))

  var depths_as_int []int

  for _, depth := range depths {
    if len(depth) == 0 {
      continue
    }
    
    as_int, err := strconv.Atoi(string(depth))
    if err != nil {
      panic(err)
    }
    depths_as_int = append(depths_as_int, as_int)
  }

  increases := 0

  for i, d := range depths_as_int {
    // Skip first measurement
    if i == 0 {
      continue;
    }

    // If this depth is greater than the one before it
    if d > depths_as_int[i - 1] {
      increases += 1
    }
  }

  fmt.Println(fmt.Sprintf("Got %d increases", increases))
}
