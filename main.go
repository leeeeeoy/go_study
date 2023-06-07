package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/leeeeeoy/go_study/model"
)

func main() {
	l, err := net.Listen("tcp", ":8080")

	if nil != err {
		log.Fatalf("fail to bind address to 8080; err: %v", err)
	}

	fmt.Println(l.Addr())

	defer l.Close()

	for {
		conn, err := l.Accept()

		if nil != err {
			log.Printf("fail to accept; err: %v", err)

			continue
		}

		go ConnHandler(conn)
	}
}

var textData []model.Message

func ConnHandler(conn net.Conn) {
	recvBuf := make([]byte, 4096)

	for {
		n, err := conn.Read(recvBuf)

		if nil != err {
			if io.EOF == err {
				log.Printf("connection is closed from client; %v", conn.RemoteAddr().String())

				return
			}

			log.Printf("fail to receive data; err: %v", err)

			return
		}

		if 0 < n {
			data := recvBuf[:n]

			var m model.Message

			err := json.Unmarshal(data[:n], &m)

			if err != nil {
				log.Println(err)
			}

			textData = append(textData, m)

			log.Println(textData)

			reqBodyBytes := new(bytes.Buffer)

			json.NewEncoder(reqBodyBytes).Encode(textData)

			reqBodyBytes.Bytes()

			_, err = conn.Write(reqBodyBytes.Bytes())
			if err != nil {
				log.Println(err)

				return
			}
		}
	}
}
