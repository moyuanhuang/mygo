////////////////////////////////////////////////////////
//  this is the Goroutine version of server.
// it can powerfully handle multiple client case, which
// is impossisble for the non-Goroutine version as in
// ./tcp_server.go.bak
////////////////////////////////////////////////////////

package main

import (
    "fmt"
    "net"
    "os"
    "time"
)

func handleConn(conn net.Conn) {
    defer conn.Close()
    conn.Write([]byte(time.Now().String()))
}

func main() {
    service := ":7777"
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkErr(err)
    listener, err := net.ListenTCP("tcp4", tcpAddr)
    checkErr(err)
    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        go handleConn(conn)
    }
}

func checkErr(err error) {
    if err != nil {
        fmt.Println("Fatal error: ", err.Error())
        os.Exit(0)
    }
}
