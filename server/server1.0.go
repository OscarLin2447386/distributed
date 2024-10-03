// package main

// import (
// 	"fmt"
// 	"net"
// 	"bufio"
// 	"log"
// 	"os"
// )

// func handleConnection(conn net.Conn, n int){
// 	// Read message from client
// 	for {
// 		message, err := bufio.NewReader(conn).ReadString('\n')
// 		if err != nil {
// 			log.Fatal(err)
// 		}else{
// 			fmt.Println("Received message from client", n, ":", message)
// 			fmt.Println("Enter the message to send:")
// 			messageServer, err := bufio.NewReader(os.Stdin).ReadString('\n')
// 			if err != nil {
// 				log.Fatal(err)
// 			}else {
// 				conn.Write([]byte(messageServer))
// 			}

// 		}
// 	}
// }

// func main(){
// 	// Listen for incoming connections
// 	ln, err := net.Listen("tcp", ":8080")
// 	if err != nil {
// 		log.Fatal(err)
// 	}else{
// 		fmt.Println("Server is listening on port 8080...")
// 	}

// 	n := 0
// 	for {
// 		// Accept incoming connections
// 		conn, err := ln.Accept()
// 		if err != nil {
// 			log.Fatal(err)
// 		}else{
// 			n++
// 			fmt.Println("Connection established with client ", n, "!")
// 			go handleConnection(conn, n)
// 		}
// 	}

// }
