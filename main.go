package main

import (
	"encoding/gob"
	"flag"
	"go-chat-cli/client"
	"go-chat-cli/login"
	"go-chat-cli/message"
	"go-chat-cli/server"

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
		server.Server()
		return
	}
	// if there is no address or username, open the login prompt
	if dialAddress == "" && username == "" {
		username, dialAddress := login.FetchLoginData()
		//append the port to the address
		dialAddress = dialAddress + ":8080"
		log.Debug("formatted IP", "ip", dialAddress)
		client.Client(username, dialAddress)
		return
	}
	// if there is no address but there is a username, throw an error
	if dialAddress == "" {
		log.Error("Please provide an address")
		return
	}
	// if there is no username but there is an address, throw an error
	if username == "" {
		log.Error("Please provide a username")
		return
	}

	// if there is an address and a username, run the client
	client.Client(username, dialAddress)

}
