package main

func main() {
  var limit int = 1000000

  sumFactorsList := generateSumFactors(limit)

  //for i:=2; i <= limit; i++ {
  //}
  var result = 0
  result = findSmallestElement(284, limit, sumFactorsList)

  println(result)
}

func findSmallestElement(num int, limit int, factorsList []int) int {
  var result = 0
  var chainLength int = 0
  var newNumber int = num
  var chain []int = make([]int, 0, 1)
  var broken bool

  for !contains(chain, newNumber) {
    chain = append(chain, newNumber)

    newNumber = factorsList[newNumber]

    if newNumber > limit {
      broken = true
      break
    }
  }

  if !broken {
    var smallest int = limit
    if len(chain) > chainLength {
      for j:=0; j < len(chain); j++ {
        if chain[j] < smallest {
          smallest = chain[j]
        }
      }

      result = smallest
    }
  }

  return result
}

func contains(arr []int, value int) bool {
  for _, i := range(arr) {
    if i == value {
      return true
    }
  }

  return false
}

func generateSumFactors(limit int) []int {
  sumFactorList := make ([]int, limit + 1)

  for i := 1; i <= limit / 2; i++ {
    for j := 2*i; j <= limit; j += i {
      sumFactorList[j] += i
    }
  }

  return sumFactorList
}
