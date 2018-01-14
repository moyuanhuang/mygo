package main

import (
    "fmt"
    "net"
    "os"
)

// func ResolveTCPAddr(net, addr string) (*TCPAddr, os.Error)

// type TCPAddr struct {
//     IP IP
//     Port int
// }
func main() {
    if(len(os.Args) != 3) {
        fmt.Printf("Usage %s [tcp4/tcp6/tcp] ip-addr\n", os.Args[0])
        os.Exit(0)
    }
    ipType := os.Args[1]
    addr := os.Args[2]
    result, err  := net.ResolveTCPAddr(ipType, addr)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("IP address is ", result.IP.String())
        fmt.Println("Port is ", result.Port)
    }
}
