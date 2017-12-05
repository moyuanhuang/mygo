package main

import (
  "fmt"
)

func fibonacci(c, quit chan int){
  x, y := 0, 1
  for{  // recursively input number into c
    select{
    case c <- x:
      x, y = y, x + y
    case <-quit:
      fmt.Println("quit")
      return
    }
  }
}

func main(){
  c := make(chan int)
  quit := make(chan int)
  go func(){
    for i := 0; i < 10; i++ {
      fmt.Println(<-c)  // queue 10 jobs to fetch from c
    }
    quit<-0
  }()
  fibonacci(c, quit);  // output 0 1 1 2 3 5 8 13 21 34 quit
}
