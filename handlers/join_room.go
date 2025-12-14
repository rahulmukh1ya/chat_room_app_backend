package handlers

import (
	"chat-room-app/models"
	"encoding/json"
	"net/http"
)

type JoinRoomRequest struct {
	RoomID string `json:"roomId"`
	PIN    string `json:"pin"`
}

func JoinRoom(w http.ResponseWriter, r *http.Request) {
	var req JoinRoomRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	room, exists := models.GetRoom(req.RoomID)
	if !exists {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	if room.PIN != req.PIN {
		http.Error(w, "Invalid PIN", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]bool{"valid": true})
}
