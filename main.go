package main

import (
	"chatterbox-cli/client"
	"chatterbox-cli/login"
	"chatterbox-cli/message"
	"chatterbox-cli/server"
	"encoding/gob"
	"flag"
	"fmt"

	log "github.com/charmbracelet/log"
)

func main() {
	gob.Register(message.Message{})
	log.SetLevel(log.DebugLevel)
	var username string
	var dialAddress string
	var isServer bool
	flag.StringVar(&dialAddress, "address", "", "Address to dial")
	flag.StringVar(&username, "username", "", "Username to use")
	flag.BoolVar(&isServer, "server", false, "Run as server")
	flag.Parse()

	if isServer {
		server.Server(username)
		return
	}
	// if there is no address or username, open the login prompt
	if dialAddress == "" && username == "" {
		username, dialAddress := login.FetchLoginData()
		// clear the terminal
		fmt.Print("\033[H\033[2J")
		//append the port to the address (default is 5051)
		dialAddress = dialAddress + ":5051"
		client.Client(username, dialAddress)
		return
	}
	// if there is no address but there is a username, throw an error
	if dialAddress == "" {
		return
	}
	// if there is no username but there is an address, throw an error
	if username == "" {
		return
	}

	// if there is an address and a username, run the client
	client.Client(username, dialAddress)

}
