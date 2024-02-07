package main

import (
	"fmt"
	"net"
)

func testConnection(ip string, port string) bool {
    conn, err := net.Dial("tcp", ip+":"+port)
    if err != nil {
        fmt.Printf("Failed to connect to %s:%s: %v\n", ip, port, err)
        return false
    }
    defer conn.Close()
    fmt.Printf("Successfully connected to %s:%s\n", ip, port)
    return true
}

func main() {
    ips := []string{"127.0.0.1", "192.168.11.209", "192.168.11.255"}
    ports := []string{"2222", "8545", "7545"}

    for _, ip := range ips {
        for _, port := range ports {
            testConnection(ip, port)
        }
    }
}
