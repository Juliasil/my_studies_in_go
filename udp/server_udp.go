package main

import (
		"fmt"
		"log"
		"net"
		"time"
)

func main() {
	udpServer, err := net.ListenPacket("udp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer udpServer.Close()

	for {
		buf := make([]byte, 1024)
    _, addr, err := udpServer.ReadFrom(buf)
		if err != nil {
			continue
		}
	}
}

