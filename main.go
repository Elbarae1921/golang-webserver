package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"webserver/util"
)

var dictionary map[string]string = map[string]string{"dog": "type of animal", "cat": "feline animal with fangs"}

var methods []string = []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}

func methodExists(item string) bool {
	for _, i := range methods {
		if i == item {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("Starting socket...")
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error establishing socket connection")
		os.Exit(1)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection")
			return
		}
		log.Printf("Client %s connected\n", c.RemoteAddr().String())
		c.Write([]byte("Welcome\n"))
		go handle(c)
	}
}

func handle(c net.Conn) {
	for {
		request, err := util.Scan(c)
		if err != nil {
			log.Printf("Client %s disconnected\n", c.RemoteAddr().String())
			c.Close()
			return
		}
		request = request[:len(request)-1]
		splitRequest := strings.Split(request, " ")
		method := splitRequest[0]

		if !methodExists(method) {
			c.Write([]byte("ERROR Method not allowed\n"))
			log.Printf("Unknown Method %s", method)
			handle(c)
			return
		}

		resource := splitRequest[1]
		version := splitRequest[2]

		// handleMethod(method, c, args)

		log.Printf("%s:Method: %s Resrouce: %s Version: %s\n", c.RemoteAddr().String(), method, resource, version)
		handle(c)
		return
	}
}

// func handleMethod(method string, c net.Conn, args []string) {
// 	switch method {
// 	case "GET":
// 		get(c, args)
// 		break
// 	case "SET":
// 		set(c, args)
// 		break
// 	case "ALL":
// 		all(c)
// 		break
// 	case "CLEAR":
// 		clear(c)
// 		break
// 	}
// }

// func get(c net.Conn, args []string) {
// 	if len(args) < 1 {
// 		c.Write([]byte("ERROR Argument not provided\n"))
// 		return
// 	}
// 	arg := args[0]

// 	definition, ok := dictionary[arg]

// 	if !ok {
// 		c.Write([]byte("ERROR Definition not set\n"))
// 		return
// 	}

// 	c.Write([]byte(fmt.Sprintf("ANSWER %s\n", definition)))
// }

// func set(c net.Conn, args []string) {
// 	if len(args) < 2 {
// 		c.Write([]byte("ERROR Not enough arguments\n"))
// 		return
// 	}

// 	key := args[0]
// 	args = args[1:len(args)]

// 	definition := strings.Join(args, " ")

// 	dictionary[key] = definition

// 	c.Write([]byte("ANSWER success\n"))
// }

// func all(c net.Conn) {
// 	response := "ANSWER\n"

// 	for key, value := range dictionary {
// 		response += fmt.Sprintf("%s: %s\n", key, value)
// 	}

// 	c.Write([]byte(response))
// }

// func clear(c net.Conn) {
// 	dictionary = map[string]string{}

// 	c.Write([]byte("ANSWER success\n"))
// }
