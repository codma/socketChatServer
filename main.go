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
	//client 접속 대기
	port := ":8080"
	ln, err := net.Listen("tcp", port)
	if err != nil {
		log.Println(err)
		return
	}

	defer ln.Close()

	//client와 연결될 경우 제네릭 변수 리턴
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			return
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {

	for {
		netData, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		//client로부터 받은 데이터 출력
		myTime := MyTime()
		fmt.Print("("+myTime+")", string(netData))

		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		myTime = MyTime()
		message := "(" + myTime + ")Boomba: " + text
		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func MyTime() string {
	timeFormat := "2006-01-02 15:04:05"
	t := time.Now()
	myTime := t.Format(timeFormat)
	return myTime
}
