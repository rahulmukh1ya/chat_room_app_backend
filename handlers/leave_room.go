package handlers

import (
	"chat-room-app/models"
	"chat-room-app/services"
	"encoding/json"
	"net/http"
)

type LeaveRoomRequest struct {
	RoomID   string `json:"roomId"`
	UserID   string `json:"userId"`
}

type LeaveRoomResponse struct {
	Valid bool `json:"valid"`
}

func LeaveRoom(w http.ResponseWriter, r *http.Request) {
	var req LeaveRoomRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	_, exists := models.GetRoom(req.RoomID)
	if !exists {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	user, exists := models.GetUser(req.RoomID, req.UserID)
	if !exists {
		http.Error(w, "User not found in room", http.StatusNotFound)
		return
	}

	success := models.RemoveUser(req.RoomID, req.UserID)
	if !success {
		http.Error(w, "User not found in room", http.StatusNotFound)
		return
	}

	err := services.BroadcastUserLeft(req.RoomID, map[string]interface{}{
		"type":   "user-left",
		"user": user,
	})
	if err != nil {
		http.Error(w, "Failed to broadcast user left", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]bool{"valid": true})
}
