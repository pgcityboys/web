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
	payload := messages.MatchRequest{UserId: "mock user", Category: "hot milfs"}
	rabbit.SendMatchmakingRequest(&payload)
}
