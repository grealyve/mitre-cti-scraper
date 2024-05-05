package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/TwiN/go-color"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var (
	URL          string
	ResponseBody []byte
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	fmt.Println("[+] Client connected")
	fmt.Println("\n---------------------")

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.Contains(string(p), "mitre_EXIT_mitre") {
			os.Exit(0)
		}
		fmt.Printf((fmt.Sprintf("%s\n%v", p, color.Reset)))

	}
}

// Request to API with the defined URL
func requester() {
	fmt.Println("[+] API Requester registered.")

	response, err := http.Get(URL)

	if err != nil {
		fmt.Println("HTTP GET error!:", err)
		return
	}
	defer response.Body.Close()

	ResponseBody, err = io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Couldn't read response:", err)
		return
	}
}
