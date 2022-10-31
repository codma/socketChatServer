package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

// server
func main() {
	port := ":8080"
	ln, err := net.Listen("tcp", port)
	if err != nil {
		log.Println(err)
		return
	}
	log.Panicln("Start!!!!!!!!!!!")

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			return
		}
		ch := make(chan string)
		go handleReciveMsg(conn, ch)
		handleSendMsg(conn)

		defer conn.Close()
	}
}

func handleSendMsg(conn net.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		myTime := MyTime()
		message := "(" + myTime + ")Boomba: " + text
		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func handleReciveMsg(conn net.Conn, ch chan string) {
	for {
		select {
		case message := <-ch:
			myTime := MyTime()
			if strings.TrimSpace(string(message)) == "STOP" {
				fmt.Println("Exiting TCP server!")
				conn.Close()
				return
			}
			fmt.Print("("+myTime+")", string(message))

		default:
			go readMsg(conn, ch)
			time.Sleep(1000 * time.Millisecond)
		}
	}
}

func readMsg(conn net.Conn, ch chan string) {
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	ch <- message
}

func MyTime() string {
	timeFormat := "2006-01-02 15:04:05"
	t := time.Now()
	myTime := t.Format(timeFormat)
	return myTime
}
