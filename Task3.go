package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	cancel := make(chan int)
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		cancel <- 1
	}()

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func handleRead(c net.Conn) {
	defer c.Close()
	for {
		msg, _ := bufio.NewReader(c).ReadString('\n')

		msg = strings.ToLower(msg)
		msg = strings.TrimRight(msg, "\n")
		msg = strings.TrimRight(msg, "\r") // for windows terminal

		if msg == "exit" {
			_, err := c.Write([]byte("\nThe End!\n"))
			if err != nil {
				log.Println(err)
			}
			c.Close()
		}
	}
}
