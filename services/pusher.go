package services

import (
	"log"
	"os"

	"github.com/pusher/pusher-http-go/v5"
)

var PusherClient *pusher.Client

func InitPusher() {
	appID := os.Getenv("PUSHER_APP_ID")
	key := os.Getenv("PUSHER_KEY")
	secret := os.Getenv("PUSHER_SECRET")
	cluster := os.Getenv("PUSHER_CLUSTER")

	if appID == "" || key == "" || secret == "" || cluster == "" {
		log.Fatal("Missing Pusher environment variables")
	}

	PusherClient = &pusher.Client{
		AppID:   appID,
		Key:     key,
		Secret:  secret,
		Cluster: cluster,
		Secure:  true,
	}

	log.Println("Pusher Initialized")
}

func BroadcastMessage(roomID string, data map[string]interface{}) error {
	channelName := "chat-" + roomID
	eventName := "new-message"

	err := PusherClient.Trigger(channelName, eventName, data)
	if err != nil {
		log.Printf("Pusher trigger error: %v", err)
		return err
	}

	log.Printf("Message sent to channel: %s", channelName)
	return nil
}
