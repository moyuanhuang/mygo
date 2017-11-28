package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main(){
  // if not setting the random seed, the result would be the same every time
  rand.Seed(time.Now().UnixNano())
  var num int
  for i := 0; i < 10; i++{
    num = rand.Intn(10) + 1
    fmt.Println(num)
  }
}
