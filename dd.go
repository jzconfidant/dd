package main

import (
	"fmt"
	"net"
	"os"
)

const (
	SERVER = ":7777"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", SERVER)
	if err != nil {
		fmt.Println("Resolve TCP Addr error")
		os.Exit(1)
	}

	l, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("Listen TCP error")
		os.Exit(1)
	}

	fmt.Printf("Server listen at port%s\n", tcpAddr.String())

	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}
		go connHandler(conn)
	}
}

func connHandler(conn net.Conn) {
	defer conn.Close()

	addr := conn.RemoteAddr()
	request := make([]byte, 1024)
	for {
		len, err := conn.Read(request)
		if err != nil {
			break
		}
		if len == 0 {
			fmt.Printf("[%s] Connect has closed by client.\n", addr.String())
			break
		} else {
			fmt.Printf("[%s] %s\n", addr.String(), request)
		}

		request = make([]byte, 1024)
	}
	fmt.Printf("[%s] Closed\n", addr.String())
}
