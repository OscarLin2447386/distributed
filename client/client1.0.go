// package main

// import (
// 	"fmt"
// 	"log"
// 	"net"
// 	"bufio"
// 	"os"
// )

// func main() {
// 	conn, err := net.Dial("tcp", "127.0.0.1:8080")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for {
// 		fmt.Println("Enter the message to send: ")

// 		message, err := bufio.NewReader(os.Stdin).ReadString('\n')
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		_, err = conn.Write([]byte(message))
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		messageServer, _ := bufio.NewReader(conn).ReadString('\n')
// 		fmt.Println("Message from the server:", messageServer)

// 	}

// }