package models

import (
	"fmt"
	"math/rand"
	"time"
)

type Room struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	PIN       string    `json:"pin"`
	CreatedAt time.Time `json:"createdAt"`
}

var rooms = make(map[string]*Room)

func CreateRoom(name string) *Room {
	room := &Room{
		ID:        generateID(),
		Name:      name,
		PIN:       generatePIN(),
		CreatedAt: time.Now(),
	}

	rooms[room.ID] = room
	return room
}

func GetRoom(id string) (*Room, bool) {
	room, exists := rooms[id]
	return room, exists
}

func generateID() string {
	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, 8)
	for i := range b {
		b[i] = chars[r.Intn(len(chars))]
	}
	return string(b)
}

func generatePIN() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	pin := r.Intn(900000) + 100000
	return fmt.Sprintf("%d", pin)

}
