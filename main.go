package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("wss URL is required as only arg")
		return
	}
	url := os.Args[1]

	conn, _, err := websocket.DefaultDialer.Dial(url, http.Header{})
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Connected to %s\n", url)

	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			text, _ := reader.ReadBytes('\n')
			text = bytes.Replace(text, []byte{'\n'}, []byte{}, -1)
			conn.WriteMessage(websocket.TextMessage, text)
		}
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			panic(err)
		}
		fmt.Printf("< %s\n", msg)
	}
}
