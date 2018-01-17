package main

import (
    "fmt"
)

type MyStruct struct {
    A int
    Str string
    StrSlice []string
}

func main() {
    s := MyStruct{}
    fmt.Printf("%d\n %s\n", s.A, s.Str)
    for str := range s.StrSlice {
        fmt.Println(str)
    }
}
