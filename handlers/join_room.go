package handlers

import (
	"chat-room-app/models"
	"chat-room-app/services"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type JoinRoomRequest struct {
	RoomID   string `json:"roomId"`
	Username string `json:"username"`
}

type JoinRoomResponse struct {
	Valid bool         `json:"valid"`
	Room  *models.Room `json:"room"`
	User  *models.User `json:"user"`
}

func JoinRoom(w http.ResponseWriter, r *http.Request) {
	var req JoinRoomRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	_, exists := models.GetRoom(req.RoomID)
	if !exists {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	thisUserId := uuid.New().String()

	models.AddUser(req.RoomID, thisUserId, req.Username)

	room, exists := models.GetRoom(req.RoomID)
	if !exists {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	err := services.BroadcastUserJoined(req.RoomID, map[string]interface{}{
		"type": "user-joined",
		"user": models.User{ID: thisUserId, Username: req.Username},
	})
	if err != nil {
		http.Error(w, "Failed to broadcast user joined", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(JoinRoomResponse{
		Valid: true,
		Room: &models.Room{
			ID:        room.ID,
			Name:      room.Name,
			CreatedAt: room.CreatedAt, 
			//Intentionally Left Out the Room PIN
			Users:     room.Users,
		},
		User: &models.User{
			ID:       thisUserId,
			Username: req.Username,
		},
	})
}
