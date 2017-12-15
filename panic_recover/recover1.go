package main

import (
    "fmt"
)

func WillPanic() {
    panic("crash!")
}

func main() {
    defer func() {
        fmt.Println("You say:", recover())
        fmt.Println("I say \"NO!\"")
    } ()

    WillPanic()
}
// output
// You say: crash!
// I say "NO!"
