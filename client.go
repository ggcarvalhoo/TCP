package main

// importando pacotes

import (
	"fmt"
	"net"
)

// estruturação dos dados

const (
	s_host = "localhost"
	s_port = "8080"
	s_type = "tcp"
)

// conectando na porta via protocolo tcp
func main() {
	connection, err := net.Dial(s_type, s_host+":"+s_port)
	if err != nil {
		panic(err)
	}

	_, err = connection.Write([]byte("Hello!"))
	b := make([]byte, 1024)
	mLen, err := connection.Read(b)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("Received: ", string(b[:mLen]))
	defer connection.Close()
}
