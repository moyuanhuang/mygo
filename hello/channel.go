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
  // IMPORTANT: 如果不把c <- 0这种写操作放在goroutine里,会有fatal error: deadlock
  // 因为如果不是线程,main函数会一直等待其他程序从管道里read,然而并不会,所以死循环了
  go func(){
    for i := 0; i < 10; i++ {
      fmt.Println(<-c)  // queue 10 jobs to fetch from c
    }
    quit<-0
  }()
  fibonacci(c, quit);  // output 0 1 1 2 3 5 8 13 21 34 quit
}
