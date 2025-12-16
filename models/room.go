package models

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type Room struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	PIN       string    `json:"pin"`
	CreatedAt time.Time `json:"createdAt"`
	Users     []User    `json:"users"`
}

var rooms = make(map[string]*Room)

func CreateRoom(name string, creatorUsername string) *Room {
	room := &Room{
		ID:        generateID(),
		Name:      name,
		PIN:       generatePIN(),
		CreatedAt: time.Now(),
		Users: []User{
			{ID: GenerateUserID(), Username: creatorUsername},
		},
	}

	rooms[room.ID] = room
	return room
}

func AddUser(roomID string, userID string, username string) bool {
	room, exists := rooms[roomID]
	if !exists {
		return false
	}

	for _, user := range room.Users {
		if user.ID == userID {
			return true
		}
	}

	room.Users = append(room.Users, User{ID: userID, Username: username})
	return true
}

func RemoveUser(roomID string, userID string) bool {
	room, exists := rooms[roomID]
	if !exists {
		return false
	}

	for i, user := range room.Users {
		if user.ID == userID {
			room.Users = append(room.Users[:i], room.Users[i+1:]...)
			return true
		}
	}

	return false
}

func GetUsers(roomID string) ([]User, bool) {
	room, exists := rooms[roomID]
	if !exists {
		return nil, false
	}
	return room.Users, true
}


func GetUser(roomID string, userID string) (User, bool) {
	room, exists := rooms[roomID]
	if !exists {
		return User{}, false
	}
	for _, user := range room.Users {
		if user.ID == userID {
			return user, true
		}
	}
	return User{}, false
}

func GetRoom(id string) (*Room, bool) {
	room, exists := rooms[id]
	return room, exists
}

func GenerateUserID() string {
	return uuid.New().String()
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
