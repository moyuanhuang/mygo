package main

import (
    "fmt"
    // "time"
    "strconv"
)
func makeCakeAndSend(cs chan string, count int) {
    for i := 1; i <= count; i++ {
        fmt.Println("Sending Cake " + strconv.Itoa(i))
        cakeName := "Strawberry Cake " + strconv.Itoa(i)
        cs <- cakeName //send a strawberry cake
    }
    // elegant way of closing channel: close as soon as finish sending data
    close(cs)
}
func receiveCakeAndPack(cs chan string) {
    for s := range cs {
        fmt.Println("Packing received cake: ", s)
    }
}
func main() {
    cs := make(chan string)
    go makeCakeAndSend(cs, 5)

    // if use goroutine for receiver without time.Sleep(), program will directly exit
    // REASON: goroutines lifecycle can not exceed main()'s. Since main() is not blocked by anything (although the goroutine is). It will exit immediately before goroutine would have time 'send the cakes'
    // go receiveCakeAndPack(cs)
    // time.Sleep(2 * 1e9)

    receiveCakeAndPack(cs)
}
