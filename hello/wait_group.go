package main

import (
    "sync"
    "log"
    "time"
    "math/rand"
)

var wgRecveivers = sync.WaitGroup{}

func ReceiverExit() {
    println("receiver has exited")
    wgRecveivers.Done()
}

func senderExit() {

}

func main() {
    rand.Seed(time.Now().UnixNano())
    log.SetFlags(0)

    const MaxRandomNumber = 100000
	const NumReceivers = 100


    wgRecveivers.Add(NumReceivers)

    dataCh := make(chan int)
    // sender
    go func() {
        for {
            if value := rand.Intn(MaxRandomNumber); value == 0 {
                close(dataCh)
                log.Println("data channel closed by sender.")
                time.Sleep(time.Second * 3)
                return
            } else {
                dataCh <- value
            }
        }
    }()

    // receiver
    for i := 0; i < NumReceivers; i++ {
        go func(){
            defer ReceiverExit()

            // for data := range dataCh {
            //     log.Println(data)
            // }
            for {
                data, ok := <- dataCh
                if ok {
                    log.Println(data)
                } else {
                    return
                }
            }
        }()
    }
    // Wait blocks until the WaitGroup counter is zero.
    wgRecveivers.Wait()
}
