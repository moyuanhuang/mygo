package main

import (
    "fmt"
    "net"
    "errors"
    "os"
)

func MyParseIP(addr string) (string, error) {
    // original interface is net.ParseIP(s string) IP
    // Also, it doesn't allow ip-address with port
    // e.g. 127.0.1.1:8080 is invalid

    ip, _, err  := net.SplitHostPort(addr)
    if err == nil {
        return ip, nil
    }

    ip2 := net.ParseIP(addr)
    if ip2 == nil {
        return "", errors.New("Invalid IP!")
    } else {
        return ip2.String(), nil
    }
}

func main() {
    if(len(os.Args) != 2) {
        fmt.Printf("Usage %s ip-addr", os.Args[0])
        os.Exit(0)
    }
    name := os.Args[1]
    addr := net.ParseIP(name) // won't take ip with port
    if addr == nil {
        fmt.Printf("result of net.ParseIP: Invalid ip address\n")
    } else {
        fmt.Printf("result of net.ParseIP: %s\n", addr.String())
    }

    addr2, err := MyParseIP(name)
    if err != nil {
        fmt.Printf("result of MyParseIP: %s\n", err)
    } else {
        fmt.Printf("result of MyParseIP: %s\n", addr2)
    }
    os.Exit(0)
}
