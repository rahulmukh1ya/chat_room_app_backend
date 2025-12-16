package handlers

import (
	"chat-room-app/models"
	"chat-room-app/services"
	"encoding/json"
	"net/http"
)

type SendMessageRequest struct {
	RoomID           string `json:"roomId"`
	EncryptedMessage string `json:"encryptedMessage"`
	UserID           string `json:"userId"`
	Username         string `json:"username"`
	Timestamp        string `json:"timestamp"`
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	var req SendMessageRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	_, exists := models.GetRoom(req.RoomID)
	if !exists {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	data := map[string]interface{}{
		"encryptedMessage": req.EncryptedMessage,
		"userId":           req.UserID,
		"username":         req.Username,
		"timestamp":        req.Timestamp,
	}

	if err := services.BroadcastMessage(req.RoomID, data); err != nil {
		http.Error(w, "Failed to send message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "sent"})
}
