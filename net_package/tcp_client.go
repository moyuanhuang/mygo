package main

import (
    "fmt"
    "net"
    "os"
    "io/ioutil"
)

func main() {
    if(len(os.Args) != 2){
        fmt.Println("Usage %s host:port", os.Args[0])
        os.Exit(0)
    }
    service := os.Args[1]
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkErr(err)
    conn, err := net.DialTCP("tcp4", nil, tcpAddr)
    checkErr(err)
    _, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
    checkErr(err)
    resp, err := ioutil.ReadAll(conn)
    checkErr(err)
    fmt.Println(string(resp))
}

func checkErr(err error) {
    if err != nil {
        fmt.Println("Fatal error: ", err.Error())
        os.Exit(0)
    }
}
