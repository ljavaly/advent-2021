package main

import (
  "bytes"
  "fmt"
  "io/ioutil"
  "strconv"
)

func sum(int_list []int) int {
  total := 0
  for _, n := range int_list {
    total += n
  }
  return total
}

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

  var first_window []int
  var second_window []int
  increases := 0
  non_increases := 0
  total := 0

  for i, _ := range depths_as_int {
    // if i < 4 {
    //   continue
    // }

    first_window = depths_as_int[i:i+3]
    second_window = depths_as_int[i+1:i+4]

    if len(first_window) != 3 {
      fmt.Println("Incomplete first window for item # %d", i)
      continue
    }

    if len(second_window) != 3 {
      fmt.Println("Incomplete second window for item # %d", i)
      continue
    }

    // fmt.Println(sum(first_window), sum(second_window))

    // If this depth is greater than the one before it
    if sum(second_window) > sum(first_window) {
      increases += 1
    } else {
      // fmt.Println(
      //   fmt.Sprintf("%d was not greater than %d",
      //     sum(second_window), sum(first_window)))
      non_increases += 1
    }
    total += 1
  }

  fmt.Println(fmt.Sprintf("Got %d increases", increases))
  fmt.Println(fmt.Sprintf("Got %d non increases", non_increases))
  fmt.Println(fmt.Sprintf("Got %d total", total))
  fmt.Println(fmt.Sprintf("All windows accounted for? %b", total == increases + non_increases))
}
