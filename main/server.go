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

// criando o servidor baseado no protocolo tcp

func main() {
	s, err := net.Listen(s_type, s_host+":"+s_port)
	fmt.Println("Running...")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for {
		connection, err := s.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Println("client connected")
		go connect(connection)
	}
}

// escutando na porta e retornando todos os dados lidos de volta para o cliente
func connect(connection net.Conn) {
	defer connection.Close()
	b := make([]byte, 1024)
	mLen, err := connection.Read(b)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		panic(err)
	}
	fmt.Println("Received: ", string(b[:mLen]))
	// _, err = connection.Write([]byte("Thanks! Got your message:" + string(b[:mLen])))
	_, err = connection.Write(b[0:mLen])
	if err != nil {
		panic(err)
	}
}
