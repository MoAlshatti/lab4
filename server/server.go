package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
)

type Message struct {
	sender  int
	message string
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func acceptConns(ln net.Listener, conns chan net.Conn) {
	// TODO: all
	// Continuously accept a network connection from the Listener
	// and add it to the channel for handling connections.
	for {
		conn, err := ln.Accept()
		handleError(err)
		conns <- conn
	}
}

func handleClient(client net.Conn, clientid int, msgs chan Message) {
	// TODO: all
	// So long as this connection is alive:
	// Read in new messages as delimited by '\n's
	// Tidy up each message and add it to the messages channel,
	// recording which client it came from.
	reader := bufio.NewReader(client)
	for {
		msgString, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		formattedMsg := fmt.Sprintf("%v: %s", clientid, msgString)
		msg := Message{clientid, formattedMsg}
		msgs <- msg

	}
}

func main() {
	clientId := 0
	// Read in the network port we should listen on, from the commandline argument.
	// Default to port 8030
	portPtr := flag.String("port", ":8030", "port to listen on")
	flag.Parse()

	//TODO Create a Listener for TCP connections on the port given above.
	ln, _ := net.Listen("tcp", *portPtr)

	//Create a channel for connections
	conns := make(chan net.Conn)
	//Create a channel for messages
	msgs := make(chan Message)
	//Create a mapping of IDs to connections
	clients := make(map[int]net.Conn)

	//Start accepting connections
	go acceptConns(ln, conns)
	for {
		select {
		case conn := <-conns:
			//TODO Deal with a new connection
			// - assign a client ID
			// - add the client to the clients map
			// - start to asynchronously handle messages from this client
			clients[clientId] = conn
			go handleClient(conn, clientId, msgs)
			clientId++

		case msg := <-msgs:
			//TODO Deal with a new message
			// Send the message to all clients that aren't the sender
			for i := range clients {
				if i != msg.sender {
					fmt.Fprint(clients[i], msg.message)

				}
			}
			fmt.Println(msg.message)
		}
	}
}
