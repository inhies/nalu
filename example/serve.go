package main

import (
	"fmt"
	"github.com/inhies/nalu"
	"github.com/jjeffery/stomp"
	"net/http"
	"time"
)

func subscribeLocal(api *nalu.API) {
	// Subscribe to the default chat topic
	sub, err := api.Stomp.Subscribe("/topic/chat.general", stomp.AckClientIndividual)
	if err != nil {
		fmt.Println("Error subscribing to chat:", err)
		return
	}

	// Print out chat messages
	for {
		msg := <-sub.C
		if msg.Err != nil {
			fmt.Println("Error reading message from subscription:", msg.Err)
		}
		fmt.Println("Chat message:", string(msg.Body))
	}
}

func main() {
	// Start the STOMP server. We are willing to ping no faster than once every
	// 10 seconds
	api, err := nalu.NewAPI(10*time.Second, stomp.Options{})
	if err != nil {
		fmt.Println("Error creating API:", err)
		return
	}

	// Runs a loop to print out messages on the chat topic
	go subscribeLocal(api)

	// Websocket location
	http.HandleFunc("/ws", api.Upgrade)

	// Chat files are in /chat/ but there's no need to have to use that URL
	// so we will serve them from /
	http.Handle("/", http.StripPrefix("/",
		http.FileServer(http.Dir("chat/"))))

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
