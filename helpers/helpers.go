package helpers

import (
	"fmt"
	"net/url"
	"strings"
	"sync"

	"github.com/TwiN/go-color"
	"github.com/gorilla/websocket"
)

var (
	CONN     *websocket.Conn
	API_ONLY = "true"
	MU       sync.Mutex
	
)

func InitiliazeWebSocketConnection() error {
	// Define the WebSocket endpoint URL.
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:7778", Path: "/ws"}

	// Establish a WebSocket connection.
	if API_ONLY == "true" {
		return nil
	}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println("Error while connecting to the WebSocket server: ", err)
		return err
	}
	CONN = conn
	return nil
}

func CloseWSConnection() error {
	if CONN != nil {
		err := CONN.Close()
		if err != nil {
			if !(API_ONLY == "true") {
				fmt.Println("WebSocket connection error while closing: ", err)
				return err
			} else {
				return nil
			}
		}
	} else {
		// Handle the case when CONN is nil
		return fmt.Errorf("CONN is nil")
	}
	return nil
}

// Sends message to websobsocket connection
func SendMessageWS(module string, msg string, loglevel string) error {
	MU.Lock()
	defer MU.Unlock()
	loglevel = strings.ToUpper(loglevel)
	if CONN != nil {
		if loglevel == "ERROR" || loglevel == "WARN" || loglevel == "INFO" {
			if loglevel == "ERROR" {
				loglevel = color.InRed("[ERROR]")
			} else if loglevel == "WARN" {
				loglevel = color.InYellow("[WARN]")
			} else {
				loglevel = color.InBlue("[INFO]")
			}

			if module != "" {
				module = color.InYellow(fmt.Sprintf("[%s]", module))
				msg = fmt.Sprintf("%s %s %s", loglevel, module, msg)
			} else {
				msg = fmt.Sprintf("%s %s", loglevel, msg)
			}

			err := CONN.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				// Check API_ONLY setting. If its true, don't return error
				if !(API_ONLY == "true") {
					return err
				}

			}
			return nil
		} else if loglevel == "" {
			err := CONN.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				// Check API_ONLY setting. If its true, don't return error
				if !(API_ONLY == "true") {
					return err
				}

			}
			return nil
		}
	} else {
		// CONN = nil
		if !(API_ONLY == "true") {
			err := fmt.Errorf("Websocket Connection is nil but trying to send a WS message.")
			return err
		} else {
			return nil
		}
	}

	return nil

}
