package main

import (
	"fmt"
	"io"
	"net"
	"os"

	"github.com/ShingoYadomoto/codecrafters-redis-go/app/resp"
)

func serve(conn net.Conn) {
	defer conn.Close()

	for {
		// ToDO: analyze request bytes
		b := make([]byte, 1024)
		_, err := conn.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}

		_, err = conn.Write(resp.SimpleStrings("PONG"))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go serve(conn)
	}
}
