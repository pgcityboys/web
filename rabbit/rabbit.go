package rabbit

import (
	"context"
	"fmt"
	"log"
	"time"
	"web/messages"
	"web/utils"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
)

var connection *amqp.Connection;
var channel *amqp.Channel;

const TIMEOUT time.Duration = 5;
const RETRIES int = 5;
const SEND_TIMEOUT time.Duration = 10;

var ChatChannel chan messages.ChatOut = make(chan messages.ChatOut);

func IntializeRMQClient() {
	rabbitAddress := utils.EnvWithDefaults("RABBITMQ_ADDRESS", "localhost")
	rabbitUser := utils.EnvWithDefaults("RABBITMQ_USER", "guest")
	rabbitPassword := utils.EnvWithDefaults("RABBITMQ_PASSWORD", "guest")
	connectionString := fmt.Sprintf("amqp://%s:%s@%s:5672/", rabbitUser, rabbitPassword,rabbitAddress)
	for range RETRIES {
		conn, err := amqp.Dial(connectionString)
		if err != nil {
			log.Println("Could not connect to rmq, retrying")
			time.Sleep(TIMEOUT * time.Second)
		} else { // Initialize client
			connection = conn;
			defer connection.Close();
			var forever chan struct {};
			err := establishChannel();
			if err != nil {
				continue;
			}
			log.Println("Connected to rmq instance")
			go handleChatMessages();
			<-forever;
		}
	}
}

func establishChannel() error {
	c, err := connection.Channel();
	if err != nil {
		log.Println("ERROR: Could not create channel")
		return amqp.Error{};
	}
	channel = c;
	// Receive messages from web app
	channel.ExchangeDeclare("web", "direct", false, false, false, false, nil);
	channel.QueueDeclare("match_req", false, false, false, false, nil);
	channel.QueueDeclare("chat_req", false, false, false, false, nil)
	channel.QueueDeclare("rooms_req", false, false, false, false, nil)
	channel.QueueDeclare("rooms_new", false, false, false, false, nil)
	channel.QueueDeclare("leave_room", false, false, false, false, nil)
	channel.QueueBind("match_req", "match_req", "web", false, nil);
	channel.QueueBind("chat_req", "chat_req", "web", false, nil);
	channel.QueueBind("rooms_req", "rooms_req", "web", false, nil);
	channel.QueueBind("rooms_new", "rooms_new", "web", false, nil);
	channel.QueueBind("leave_room", "leave_room", "web", false, nil);	
	// Send out messages to topic exchange
	channel.ExchangeDeclare("matchmaking", "topic", false, false, false, false, nil);
	channel.QueueDeclare("match_res", false, false, false, false, nil);
	channel.QueueDeclare("chat_notify", false, false, false, false, nil)
	channel.QueueDeclare("room_info", false, false, false, false, nil)
	channel.QueueBind("match_res", "match_res", "matchmaking", false, nil);
	channel.QueueBind("chat_notify", "chat_notify", "matchmaking", false, nil);
	channel.QueueBind("room_info", "room_info", "matchmaking", false, nil);
	return nil;
}

func ensureChannelHealth() error {
	if channel == nil || channel.IsClosed() {
		return establishChannel();
	}
	return nil;
}

func SendMatchmakingRequest(request *messages.MatchRequest) {
	binaryData, _ := proto.Marshal(request)
	channel.Publish("web", "match_req", false, false, amqp.Publishing{Body: binaryData})
}

func SendNewRoomRequest(request *messages.NewRoomRequest) {
	binaryData, _ := proto.Marshal(request)
	channel.Publish("web", "rooms_new", false, false, amqp.Publishing{Body: binaryData})
}

func SendCategoryInfoRequest(request *messages.CategoryInfoRequest) {
	binaryData, _ := proto.Marshal(request)
	channel.Publish("web", "rooms_req", false, false, amqp.Publishing{Body: binaryData})
}

func SendLeaveRoomRequest(request *messages.MatchRequest) {
	binaryData, _ := proto.Marshal(request)
	log.Println("Sending room leave request")
	channel.Publish("web", "leave_room", false, false, amqp.Publishing{Body: binaryData})
}

func SendChatMessage(request *messages.ChatIn) {
	binaryData, _ := proto.Marshal(request)
	channel.Publish("web", "chat_req", false, false, amqp.Publishing{Body: binaryData})
}

func handleChatMessages() {
	chatChan, _ := connection.Channel();
	msgs, _ := chatChan.ConsumeWithContext(context.Background(), "chat_notify", "webapp", true, false, false, false, nil)
	for msg := range msgs {
		var chat messages.ChatOut;
		proto.Unmarshal(msg.Body, &chat);
		ChatChannel<-chat;
	}
}
