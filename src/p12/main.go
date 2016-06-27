package main

import (
  "fmt"
)

func main() {
  var triangularNumber int = 1
  var nextNumber int = 2
  var requiredCountDivisors = 100

  for countDivisors(triangularNumber) < requiredCountDivisors {
    triangularNumber = triangularNumber + nextNumber
    nextNumber++
  }

  fmt.Println(triangularNumber)
}

func countDivisors(num int) int {
  var count int = 0

  for i := 1; i <= num; i++ {
    if num % i == 0 {
      count++
    }
  }

  return count
}
