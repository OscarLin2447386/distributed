package main

import (
	"flag"
	"net"
	"bufio"
	"fmt"
	"os"
)

func read(conn net.Conn) {
	//TODO In a continuous loop, read a message from the server and display it.
	for {
		message_from_server, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println(message_from_server)
	}

}

func write(conn net.Conn) {
	//TODO Continually get input from the user and send messages to the server.
	for {
		fmt.Println("Enter the message to the server:")
		message_to_server, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		conn.Write([]byte(message_to_server))
	}
}

func main() {
	// Get the server address and port from the commandline arguments.
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()
	//TODO Try to connect to the server
	//TODO Start asynchronously reading and displaying messages
	//TODO Start getting and sending user messages.
	conn, _ := net.Dial("tcp", *addrPtr)
	go read(conn)
	write(conn)
}
