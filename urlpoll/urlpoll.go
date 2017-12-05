package main

import (
  "log"
  "net/http"
  "time"
)

const (
  numPollers = 2
  pollInterval = 5 * time.Second
  statusInterval = 2 * time.Second
  errTimeout = 2 * time.Second
)

var urls = []string{
  "http://www.google.com/",
  "http://golang.org/",
  "http://blog.golang.org/",
}

type State struct{
  url string
  status string
}

func StateMonitor(updateInterval time.Duration) chan<- State {
  updates := make(chan State)
  urlStatus := make(map[string]string)
  ticker := time.NewTicker(updateInterval)
  go func(){
    for{  // loop forever, even though returned
      select{
      case <-ticker.C:
        logState(urlStatus)
      case s := <-updates:
        urlStatus[s.url] = s.status
      }
    }
  }()
  return updates
}

func logState(urlStatus map[string]string) {
  for k, v := range urlStatus{
    log.Printf(" %s %s", k, v)
  }
}

type Resource struct {
  url string
  errCount int
}

func (r *Resource) Poll() string{
  resp, err := http.Head(r.url)
  if err != nil {
    log.Println("Error", r.url, err)
    r.errCount++
    return err.Error()
  }
  r.errCount = 0
  return resp.Status
}

func (r *Resource) Sleep(done chan<- *Resource) {
  // must convert int to time.Duration
  time.Sleep(pollInterval + time.Duration(r.errCount) * errTimeout)
  done <- r
}

func Poller(in <-chan *Resource, out chan<- *Resource, status chan<- State){
  for r := range in{
    s :=  r.Poll()
    status <- State{r.url, s}
    out <- r
  }
}

func main(){
  pending, complete := make(chan *Resource), make(chan *Resource)
  status := StateMonitor(statusInterval)

  for i := 0; i < numPollers; i++{
    go Poller(pending, complete, status)
  }

  go func() {
    for _, url := range urls {
      pending <- &Resource{url, 0}
    }
  } ()

  // go func() {
  //   for r := range complete{
  //     r.Sleep(pending)
  //   }
  // } ()

  // why seperate goroutine for each elem in complete??
  for r := range complete {
    go r.Sleep(pending)
  }
}
