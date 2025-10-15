package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

// what to achive ?
// task : 1 get the args from the command line if less than 2 exit and print usage
// task : 2 get the port from the command line and start a server on that port
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Use: go run main.go <port> pls :D")
		os.Exit(1)
	}
	port := fmt.Sprintf(":%s", os.Args[1])
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("error occured D: err --> ", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Server started on port", port)
	// handle connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error occured D: err --> ", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	// makes the reading of data more efficient and convenient
	reader := bufio.NewReader(conn)
	for {
		by, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			if err != io.EOF {
				fmt.Println("error occured D: err --> ", err)
			}
			return
		}
		fmt.Printf("request recived: %s", by)
		line := fmt.Sprintf("Echo: %s", by)
		fmt.Printf("response: %s", by)
		
		_, err = conn.Write([]byte(line))
		if err != nil {
			fmt.Println("error occured D: err --> ", err)
			return
		}
	}
}
