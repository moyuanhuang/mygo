package main

import (
    "fmt"
    "net"
    "os"
    "io/ioutil"
)

// func DialTCP(net string, laddr, raddr *TCPAddr) (c *TCPConn, err os.Error)
func main() {
    if len(os.Args) != 2 {
        fmt.Printf("Usage %s host:port", os.Args[0])
        os.Exit(0)
    }
    service := os.Args[1]
    addr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)
    conn, err := net.DialTCP("tcp4", nil, addr)
    checkError(err)
    _, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
    checkError(err)
    // func ReadAll(r io.Reader) ([]byte, error)
    // ReadAll reads from r until an error or EOF and returns the data it read.
    resp, err := ioutil.ReadAll(conn)
    checkError(err)
    fmt.Println(string(resp))
    os.Exit(0)
}

func checkError(err error) {
    if err != nil {
        fmt.Println("Fatal Error: ", err.Error())
        // fmt.Fprintf(os.Stderr, "Fatal error: %s", err.String())
        os.Exit(1)
    }
}
