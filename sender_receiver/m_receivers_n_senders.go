/*
Refer to this post. https://moyuanhuang.github.io/2017/12/13/how-should-we-close-channel/#more. This is for the N receivers M senders case.
*/

package main

import (
    "fmt"
    "sync"
)

func main() {
    NumberOfValues := 100
    values := make([]int, NumberOfValues)
    for i := range values{
        values[i] = i
    }

    result := ComputeSum(values)
    fmt.Printf("Sum: %d\n", result)
}

func ComputeSum(values []int) int {
    var sum int
    ch := make(chan int, 100)

    NumberOfRoutines := 10
    chunkSize := len(values) / NumberOfRoutines

    wg := sync.WaitGroup{}
    wg.Add(NumberOfRoutines)

    for i := 0; i < 10; i++ {
        offset := i * 10
        chunk := values[offset: offset + chunkSize]
        go func(){
            defer wg.Done()
            s := 0
            for _, v := range chunk {
                s += v
            }
            ch <- s
        }()
    }

    go func(){
        wg.Wait()
        close(ch)
    }()

    for v := range ch{
        sum += v
    }
    return sum
}
