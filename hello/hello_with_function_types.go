package main

import "fmt"

type Greeting func(name string) string

func (g Greeting) say(n string){
  fmt.Println(g(n))
}

func english(n string) string{
  return "Hello, " + n
}

func french(n string) string{
  return "Bonjour, " + n
}

func main(){
  g := Greeting(english)
  g.say("Quandan")
  g = Greeting(french)
  g.say("Jaculine")
}
