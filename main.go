package main

import (
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":25")
	if err != nil {
		log.Fatal("Error listening", err)
	}

	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal("Error accepting connection", err)
			continue
		}

		go handleConn(connection)
	}
}

func handleConn(conn net.Conn) {

	log.Println("::	Conn accepted")

	defer conn.Close()

	response := "220 MAILCHUCK \n"
	_, err := conn.Write([]byte(response))

	if err != nil {
		log.Printf("Server write error: %v", err)
	}
}
