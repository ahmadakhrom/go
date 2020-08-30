package main

import (
  "fmt"
  "testing"
)

func TestSum(t *testing.T) {
  t.Run("[1,2,3,4,5]", testSumFunc([]int{1, 2, 3, 4, 5}, 15))
  t.Run("[1,2,3,4,-5]", testSumFunc([]int{1, 2, 3, 4, -5}, 5))
}

func testSumFunc(numbers []int, expected int) func(*testing.T) {
  return func(t *testing.T) {
    actual := Sum(numbers)
    if actual != expected {
      t.Error(fmt.Sprintf("Expected the sum of %v to be %d but instead got %d!", numbers, expected, actual))
    }
  }
}

func Sum(numbers []int) int  {
  var sum int
  for i := range numbers {
    sum += i
  }
  return sum
}