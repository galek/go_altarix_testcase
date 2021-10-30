package main

import (
	"fmt"
	"github.com/takama/daemon"
	"log"
)

/*
type MessageData struct {
	PersonName  string `json:"person_Name" validate:"required"`
	Date        string `json:"date" validate:"required"`
	PersonEmail string `json:"person_email"`
	PersonSMS   string `json:"person_sms"`
	PersonPush  string `json:"person_push"`
}

// TODO: тут можно сразу определить что есть что
type MessageIn struct {
	AccessToken string `json:"access_token validate:"required"`
	EventCode   string `json:"event_code" validate:"required"`
	StreamType  string `json:"stream_type" validate:"required"`
	Data        MessageData
}
type MessageOut struct {
	AccessToken string `json:"access_token"`
	EventCode   string `json:"event_code"`
	StreamType  string `json:"stream_type"`
	To          string `json:"to"`
	Data        MessageData
}*/

func main() {
	log.Println("[main] Started")

	srv, err := daemon.New(name, description, daemon.GlobalAgent)
	if err != nil {
		// TODO: Посмотреть как в MO я делал
		log.Fatal("Can't run demon. Started with error:", err)
		return
	}

	service := &Service{srv}

	status, err := service.Manage()
	if err != nil {
		// TODO: По идее должна быть паника
		log.Fatal("Can't run service Manage. Started with error:", err)
		return
	}

	log.Println(fmt.Sprintf("[main] Finished, with status: %s", status))
}
