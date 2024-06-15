package wsserver

import (
	"encoding/json"
	"log"
	"web/messages"
	"web/rabbit"
)

type WebsocketMessage struct {
	User  string `json:"user"`
	Event string `json:"event"`
	Data  map[string]interface{} `json:"data"`
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
	case "join":
		handleJoin(&message)
	case "category":
		handleCategoryInfo(&message)
	case "leave":
		handleLeaveRoom(&message)
	case "chat":
		handleChatSend(&message)
	default:
		log.Println("AAAAAA zly json")
	}
	return nil;
}

// Message handlers + parsers
func handleCreate(message *WebsocketMessage) {
	log.Println("Sending room create message")
	payload := messages.NewRoomRequest{UserId: message.User, Category: message.Data["category"].(string), MaxUsers: int32(message.Data["maxusers"].(float64)), Description: message.Data["description"].(string)}
	rabbit.SendNewRoomRequest(&payload);
}

func handleJoin(message *WebsocketMessage) {
	log.Println("Sending room join message")
	payload := messages.MatchRequest{UserId: message.User, RoomId: message.Data["id"].(string)}
	rabbit.SendMatchmakingRequest(&payload)
}

func handleCategoryInfo(message *WebsocketMessage) {
	log.Println("Fetching category info")
	payload := messages.CategoryInfoRequest{Category: message.Data["category"].(string), UserId: message.User}
	rabbit.SendCategoryInfoRequest(&payload)
}

func handleLeaveRoom(message *WebsocketMessage) {
	log.Println("Leaving room")
	payload := messages.MatchRequest{UserId: message.User, RoomId: message.Data["id"].(string)}
	rabbit.SendLeaveRoomRequest(&payload)
}

func handleChatSend(message *WebsocketMessage) {
	log.Println("Sending message")
	payload := messages.ChatIn{UserId: message.User, RoomId: message.Data["room"].(string), Content: message.Data["content"].(string)}
	rabbit.SendChatMessage(&payload)
}
