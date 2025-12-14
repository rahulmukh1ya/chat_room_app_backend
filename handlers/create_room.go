package handlers

import (
	"chat-room-app/models"
	"encoding/json"
	"net/http"
)

type CreateRoomRequest struct {
	RoomName string `json:"roomName"`
}

type CreateRoomResponse struct {
	RoomID string `json:"roomId"`
	PIN    string `json:"pin"`
}

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	var req CreateRoomRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	room := models.CreateRoom(req.RoomName)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CreateRoomResponse{
		RoomID: room.ID,
		PIN: room.PIN,
	})
}