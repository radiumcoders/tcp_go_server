package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// os.Args gives a slice of strings containing the command-line arguments passed to the program.
	// os.Args[0] is the path to the executable file, os.Args[1] is the first argument, and so on.
	if len(os.Args) < 2 {
		fmt.Println("Usage: tcpserver <port>")
		os.Exit(1)
	}
	// here we are extracting the port from the command-line arguments
	port := fmt.Sprintf(":%s", os.Args[1])

	// extracting the listener and error
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// using defer to close the listener when all the code ends and resources are released
	defer listener.Close()
	// listener listening on 
	fmt.Printf("tcp server listening on %n port" , listener.Addr())
}
