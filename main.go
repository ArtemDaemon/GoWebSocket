package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Message struct {
	Action string `json:"action"`
}

type Response struct {
	Status   string `json:"status"`
	Progress int    `json:"progress,omitempty"`
	Data     []User `json:"data,omitempty"`
	Message  string `json:"message,omitempty"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func loadHeavyData(size int) []User {
	users := make([]User, size)
	for i := 0; i < size; i++ {
		users[i] = User{
			ID:    i + 1,
			Name:  randomName(),
			Email: randomEmail(),
		}
		time.Sleep(10 * time.Millisecond) // Simulate heavy data loading
	}
	return users
}
func randomName() string {
	names := []string{"Alice", "Bob", "Charlie", "Diana", "Eve"}
	return names[rand.Intn(len(names))]
}
func randomEmail() string {
	emails := []string{"example.com", "test.com", "demo.com"}
	return fmt.Sprintf("%s@%s", randomName(), emails[rand.Intn(len(emails))])
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		var req Message
		if err := json.Unmarshal(msg, &req); err != nil {
			conn.WriteJSON(Response{Status: "error", Message: "Invalid request"})
			continue
		}

		switch req.Action {
		case "load_sync":
			// Synchronous: blocks until all data is loaded
			users := loadHeavyData(50)
			conn.WriteJSON(Response{Status: "done", Data: users, Message: "Loaded synchronously"})
		case "load_async":
			// Asynchronous: send progress updates
			go func(c *websocket.Conn) {
				size := 50
				users := make([]User, 0, size)
				for i := 0; i < size; i++ {
					users = append(users, User{
						ID:    i + 1,
						Name:  randomName(),
						Email: randomEmail(),
					})
					time.Sleep(10 * time.Millisecond)
					if i%10 == 0 || i == size-1 {
						c.WriteJSON(Response{
							Status:   "progress",
							Progress: (i + 1) * 100 / size,
							Message:  "Loading asynchronously...",
						})
					}
				}
				c.WriteJSON(Response{Status: "done", Data: users, Message: "Loaded asynchronously"})
			}(conn)
		default:
			conn.WriteJSON(Response{Status: "error", Message: "Unknown action"})
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/ws", wsHandler)
	log.Println("WebSocket server started at ws://localhost:8080/ws")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
