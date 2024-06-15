package api

import (
	"log"
	"net/http"
	"web/messages"
	"web/rabbit"

	"github.com/google/uuid"
)

// Hit up this api endpoint to send matchmaking request
func HandleMatchRequest(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	log.Println("Sending mock match request")
	userId := uuid.New().String()
	payload := messages.MatchRequest{UserId: userId, RoomId: id}
	rabbit.SendMatchmakingRequest(&payload)
}

// Hit up this api endpoint to send new room request
func HandleNewRoom(w http.ResponseWriter, r *http.Request) {
	log.Println("Sending new room request")
	payload := messages.NewRoomRequest {UserId: "mock user", Category: "Hot milfs", Description: "Cool room, asf", MaxUsers: 4}
	rabbit.SendNewRoomRequest(&payload)
}
