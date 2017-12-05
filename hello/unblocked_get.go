package main

import (
  "fmt"
  // "github.com/astaxie/beego/httplib"
)

func UnblockedGet(c chan string, msg string) {
  go func() {
    // request := httplib.Get(url)
    // content, err := request.String()
    // if err != nil {
    //     content = "" + err.Error()
    // }
    c <- msg
  } ()
}

func main() {
  ch1 := make(chan string)
  ch2 := make(chan string)
  UnblockedGet(ch1, "i am calling UnblockedGet on ch1")
  UnblockedGet(ch2, "i am calling UnblockedGet on ch2")

  fmt.Println(<-ch1)
  fmt.Println(<-ch2)
}
