package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func read(conn net.Conn, channel chan<- int) {
	//TODO In a continuous loop, read a message from the server and display it.
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err == io.EOF {
			channel <- 1
		}
		fmt.Println(msg)

	}
}

func write(conn net.Conn, channel chan<- int) {
	//TODO Continually get input from the user and send messages to the server.
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter your message: ")
		msg, err := reader.ReadString('\n')
		if err == io.EOF {
			channel <- 1
		}
		fmt.Fprint(conn, msg)
	}
}

func main() {
	retchann := make(chan int)
	// Get the server address and port from the commandline arguments.
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()
	//TODO Try to connect to the server
	//TODO Start asynchronously reading and displaying messages
	//TODO Start getting and sending user messages.
	conn, err := net.Dial("tcp", *addrPtr)
	handleError(err)
	go read(conn, retchann)
	go write(conn, retchann)

	select {
	case <-retchann:
		return
	}

}
