package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"sync"
)

// ChatServer holds all messages and manages synchronization
type ChatServer struct {
	messages []string
	mu       sync.Mutex
}

// MessageData struct represents the message and sender name
type MessageData struct {
	Name    string
	Message string
}

// SendMessage adds a formatted message to chat history and returns the full chat
func (cs *ChatServer) SendMessage(data MessageData, reply *[]string) error {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	formatted := fmt.Sprintf("%s: %s", data.Name, data.Message)
	cs.messages = append(cs.messages, formatted)

	*reply = make([]string, len(cs.messages))
	copy(*reply, cs.messages)

	// Show message on server console too
	fmt.Println(formatted)

	return nil
}

func main() {
	server := new(ChatServer)
	rpc.Register(server)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
	defer listener.Close()

	fmt.Println("Chatroom server is running on port 1234...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Connection error:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
