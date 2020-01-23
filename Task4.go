package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	hostPort    = "localhost:8000"
	exitSeconds = 5
)

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

// 1. Сервер пишет приветствие
// 2. Выход с сервера
// 3. Оповещение клиентов
// 4. Закрытие сервера
func main() {
	listener, err := net.Listen("tcp", hostPort)
	if err != nil {
		log.Fatal(err)
	}
	go waitExit()

	log.Println("Hello!: ", hostPort)

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all clients
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)
	who := conn.RemoteAddr().String()
	ch <- "You are " + who + "\rType 'exit' for disconnect"
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		s := input.Text()
		if strings.ToLower(s) == "exit" {
			ch <- "Good bye!"
			time.Sleep(time.Second) // что бы успеть увидить сообщение перед выходом
			break
		}
		messages <- who + ": " + s
	}

	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func waitExit() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if strings.ToLower(scanner.Text()) == "exit" {
			log.Println("Sending to clients messages about", exitSeconds, "seconds...")
			messages <- "Server shutdown in " + strconv.Itoa(exitSeconds) + " seconds, Sorry..."
			time.Sleep(exitSeconds * time.Second)
			log.Println("Exiting... Good bye!")
			os.Exit(0)
		}
	}
}
