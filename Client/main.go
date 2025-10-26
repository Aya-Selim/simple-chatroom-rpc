package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
)

// MessageData struct (must match server)
type MessageData struct {
	Name    string
	Message string
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}
	defer client.Close()

	reader := bufio.NewReader(os.Stdin)

	// Ask for user name first
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Printf("Welcome %s! You can start chatting now (type 'exit' to quit)\n\n", name)

	for {
		fmt.Print("Enter message: ")
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)

		if message == "exit" {
			fmt.Printf("Exiting chatroom... Bye  %s ", name)
			break
		}

		data := MessageData{Name: name, Message: message}
		var chatHistory []string

		err = client.Call("ChatServer.SendMessage", data, &chatHistory)
		if err != nil {
			fmt.Println("Error sending message:", err)
			break
		}

		fmt.Println("\nChat history:")
		for _, msg := range chatHistory {
			fmt.Printf("- %s\n", msg)
		}
		fmt.Println()
	}
}
