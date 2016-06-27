package main

func main() {
  var limit int = 1000000
  var length int = 0
  var sequence int
  var longestLength int = 1
  var longestStart int

  for start := limit; start > 1; start -- {
    sequence = start
    length = 0

    for sequence > 1 {
      if sequence % 2 == 0 {
        sequence = sequence / 2
      } else {
        sequence = 3 * sequence + 1
      }
      length ++
    }

    if length > longestLength {
      longestStart = start
      longestLength = length
    }
  }

  println(longestStart)
  println(longestLength)
}
