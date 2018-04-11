package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	connHost := "localhost:1234"
	flag.StringVar(&connHost, "host", "localhost:1234", "hostname:port to bind the server to")
	flag.Parse()

	l, err := net.Listen("tcp", connHost)
	if err != nil {
		log.Fatal("Error listening:", err.Error())
	}
	defer l.Close()
	log.Println("Listening on ", connHost)

	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	file, err := os.Create(fmt.Sprintf("ok-%d.log", time.Now().Unix()))
	if err != nil {
		log.Fatal("unable to create log file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Fprintln(file, line)

		_, err := fmt.Fprintln(conn, "OK")
		if err != nil {
			log.Println("error sending ok", err)
		}
	}

	defer conn.Close()
}
