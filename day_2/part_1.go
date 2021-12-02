package main

import (
  "fmt"
  "strconv"
  "strings"
  "utils"
)

var filepath string = "/Users/ljavaly/Documents/advent_of_code_2021/day_2/input.txt"

// Define Submarine struct with properties and methods
type Submarine struct {
  position int
  depth    int
}

func (s *Submarine) forward(n int) {
  s.position += n
}
func (s *Submarine) up(n int) {
  s.depth -= n
}
func (s *Submarine) down(n int) {
  s.depth += n
}

func (s *Submarine) applyStep(direction string, distance int) {
  switch direction {
  case "forward":
    s.forward(distance)
  case "up":
    s.up(distance)
  case "down":
    s.down(distance)
  default:
    panic(fmt.Sprintf("Got invalid direction! %s", direction))
  }
}


// Define Step struct and its properties
type Step struct {
  direction string
  distance  int
}

func parseStep(input string) Step {
  parts := strings.Split(input, " ")
  direction, distance := parts[0], parts[1]

  distance_int, err := strconv.Atoi(distance)
  if err != nil {
    panic(err)
  }

  return Step{direction: direction, distance: distance_int}
}


func main() {
  inputs := utils.LoadInput(filepath)
  // fmt.Println(inputs[:5])

  var sub Submarine

  for _, input := range(inputs) {
    step := parseStep(input)
    // fmt.Println(step.direction, step.distance)
    sub.applyStep(step.direction, step.distance)
  }

  fmt.Println("Final position:", sub.position)
  fmt.Println("Final depth:", sub.depth)
  fmt.Println("Position * depth =", sub.position * sub.depth)
}
