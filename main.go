package main

import {
	"crypto/ed25519"
	"crypto/subtle"
	"encoding/json"
	
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
}

//  Interaction represents the structure of an interaction received from Discord.
type Interaction struct { 
	Type int `json:"type"`
}

// InteractionResponse represents the structure of the response sent to Discord.
type InteractionResponse struct {
	Type int `json:"type"`
}

const {
	InteractionTypePing = 1  // Ping type interaction.
	InteractionTypeCommand = 2 // Command Type Interaction.
	InteractionResponsePong = 1 // Pong respong for a Ping interaction.
}

// main starts the http server to and listen for incoming request.

func main () {
	// set up the http server to route all request to the handlerquest..
	http.HandleFunc("/". handlerquest)
	// read the server port fromt the envirorment variable  or port default to 8080.
	port := os.Getenv("Port")
	if port == "" {
		port = "8080"
	}
}

fmt.Printf("Starting server on port s%.....\n", port)

// start the http server 
if err := http.ListenAndServe(":"
)
