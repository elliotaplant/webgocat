package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("wss URL is required as only arg")
		return
	}
	url := os.Args[1]

	conn, resp, err := websocket.DefaultDialer.Dial(url, http.Header{})
	if err != nil {
		fmt.Printf("Error establishing WebSocket connection: %s\n", err)
		if resp != nil {
			fmt.Printf("Response received\n")
			fmt.Printf("Status Code: %d\n", resp.StatusCode)
			body, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				fmt.Printf("Body: %s\n", body)
			}
			fmt.Printf("Headers:\n")
			for key, values := range resp.Header {
				fmt.Printf("    %s: %s\n", key, strings.Join(values, ", "))
			}
		}
		return
	}
	defer conn.Close()
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
			fmt.Printf("Error reading message from WebSocket: %s\n", err)
		}
		fmt.Printf("< %s\n", msg)
	}
}
