package api

import (
	"log"
	"net/http"
	"web/messages"
	"web/rabbit"
)

// Hit up this api endpoint to send matchmaking request
func HandleMatchRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("Sending mock match request")
	payload := messages.MatchRequest{UserId: "mock user", RoomId: "test"}
	rabbit.SendMatchmakingRequest(&payload)
}

// Hit up this api endpoint to send new room request
func HandleNewRoom(w http.ResponseWriter, r *http.Request) {
	log.Println("Sending new room request")
	payload := messages.NewRoomRequest {UserId: "mock user", Category: "Hot milfs", Description: "Cool room, asf", MaxUsers: 4}
	rabbit.SendNewRoomRequest(&payload)
}
