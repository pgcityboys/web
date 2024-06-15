package wsserver

import (
	"encoding/json"
	"log"
)

type WebsocketMessage struct {
	User  string `json:"user"`
	Event string `json:"event"`
	Data  []byte `json:"data"`
}

type ChatMessage struct {
	Room    string `json:"room"`
	Content string `json:"content"`
}

type JoinRoom struct {
	Id string `json:"id"`
}

type LeaveRoom struct {
	Id string `json:"id"`
}

type CreateRoom struct {
	Category    string `json:"category"`
	Description string `json:"description"`
	MaxUsers    int    `json:"maxusers"`
}

func handleInMessage(data []byte) error {
	var message WebsocketMessage;
	json.Unmarshal(data, &message);
	switch message.Event {
	case "create":
		handleCreate(&message)	
	default:
		log.Println("AAAAAA zly json")
	}
	return nil;
}

// Message handlers + parsers
func handleCreate(message *WebsocketMessage) {
	var payload CreateRoom;
	json.Unmarshal(message.Data, &payload)
	log.Println(payload)
}

