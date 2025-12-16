package handlers

import (
	"chat-room-app/models"
	"encoding/json"
	"net/http"
)

type CreateRoomRequest struct {
	RoomName string `json:"roomName"`
	Username string `json:"username"`
}

type CreateRoomResponse struct {
	RoomID string        `json:"roomId"`
	PIN    string        `json:"pin"`
	Users  []models.User `json:"users"`
}

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	var req CreateRoomRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	room := models.CreateRoom(req.RoomName, req.Username)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CreateRoomResponse{
		RoomID: room.ID,
		PIN:    room.PIN,
		Users:  room.Users,
	})
}
