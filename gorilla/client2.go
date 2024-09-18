package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	addr := "localhost:8080"
	uid := "abc"
	u := url.URL{Scheme: "ws", Host: addr, Path: "ws", RawQuery: "uuid=" + uid}

	log.Printf("Connecting to %s", u.String())

	header := http.Header{}
	header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBfa2V5IjoiYWlTZXJhY2giLCJhcHBfc2VjcmV0IjoiOWRoNFVvQ19zaiIsInV1aWQiOiJmMjQ5ZWRiZjQiLCJleHAiOjE3MjYyMTA4MTgsImlzcyI6ImFpU2VyYWNoIn0.CpBNZJyGae6nGO7XnY7dMHTVXxvY039qCBQ9M3rK8Lo")

	var mu sync.Mutex
	var c *websocket.Conn
	var err error

	done := make(chan struct{})

	msgChan := make(chan string)
	go func() {
		for {
			msgChan <- time.Now().Format("2006-01-02 15:04:05")
			time.Sleep(5 * time.Second)
		}

	}()

	connect := func() {
		mu.Lock()
		defer mu.Unlock()
		for {
			c, _, err = websocket.DefaultDialer.Dial(u.String(), header)
			if err == nil {
				log.Println("Connected to WebSocket server")
				break
			}
			log.Printf("Dial error: %v, retrying...", err)
			time.Sleep(5 * time.Second)
		}
		go readMessages(c, done)
		go pingMessages(c, done, msgChan)
	}

	connect()
	defer func() {
		if c != nil {
			c.Close()
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	c.SetPongHandler(func(appData string) error {
		log.Println("Received Pong")
		return c.SetReadDeadline(time.Now().Add(60 * time.Second))
	})

	for {
		select {
		case <-done:
			log.Println("Connection closed, reconnecting...")
			connect()
		case <-interrupt:
			log.Println("Interrupt received, closing connection...")

			mu.Lock()
			if c != nil {
				err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				if err != nil {
					log.Println("Write close error:", err)
					mu.Unlock()
					return
				}
			}
			mu.Unlock()
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

func readMessages(c *websocket.Conn, done chan struct{}) {
	defer func() {
		done <- struct{}{}
	}()
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Unexpected close error: %v", err)
			}
			return
		}
		log.Printf("Received message: %s", message)
	}
}

func pingMessages(c *websocket.Conn, done chan struct{}, msgChan chan string) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := c.WriteMessage(websocket.PingMessage, []byte("ping")); err != nil {
				log.Println("Write ping error:", err)
				c.Close()
				done <- struct{}{}
				return
			}
		case msg := <-msgChan:
			if err := c.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
				log.Println("Write TextMessage error:", err)
			}
		}
	}
}
