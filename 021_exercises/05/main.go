package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go writeread(conn)
	}
}

func writeread(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		//line := scanner.Text()
		fmt.Println("Code got here")
		io.WriteString(conn, "You are connected")
	}
	defer conn.Close()

	fmt.Println("Selesai")

}
