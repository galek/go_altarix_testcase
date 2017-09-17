package main

import (
	"log"

	"github.com/pquerna/ffjson/ffjson"
)

// JSON
// https://golang.org/pkg/encoding/json/#pkg-examples

type MessageDataIn struct {
	Person_Name  string `json:"person_Name"`
	Date         string `json:"date"`
	Person_email string `json:"person_email"`
	PersonSMS    string `json:"person_sms"`
	PersonPush   string `json:"person_push"`
}
type MessageDataOut struct {
	Person_Name string `json:"person_Name"`
	Date        string `json:"date"`
}

type MessageIn struct {
	Access_token string `json:"access_token"`
	Event_code   string `json:"event_code"`
	Stream_type  string `json:"stream_type"`
	Data         MessageDataIn
}

type MessageOut struct {
	Access_token string `json:"access_token"`
	Event_code   string `json:"event_code"`
	Stream_type  string `json:"stream_type"`
	To           string `json:"to"`
	Data         MessageDataOut
}

func MessageInToMessageToConverter(in *MessageIn, to *MessageOut, jsonStream string) {
	if ISDebug {
		log.Printf("JSON %s", jsonStream)
	}

	ffjson.Unmarshal([]byte(jsonStream), &in)

	// debug
	//if ISDebug {
	// 	fmt.Println(in)
	// 	fmt.Println(in.Data.Person_Name)
	// }

	if ValidateJSONMessage(in) == false {
		log.Println("Failed validation")
		return
	}

	if in.Stream_type == "email" || in.Stream_type == "EMAIL" {
		if ISDebug {
			log.Println("[Debug] stream_type is email. ", in.Data.Person_email)
		}
		to.To = in.Data.Person_email
	} else if in.Stream_type == "sms" || in.Stream_type == "SMS" {
		if ISDebug {
			log.Println("[Debug] stream_type is sms. ", in.Data.PersonSMS)
		}
		to.To = in.Data.PersonSMS
	} else if in.Stream_type == "push" || in.Stream_type == "PUSH" {
		if ISDebug {
			log.Println("[Debug] stream_type is push. ", in.Data.PersonPush)
		}
		to.To = in.Data.PersonPush
	}

	to.Access_token = in.Access_token
	to.Event_code = in.Event_code
	to.Stream_type = in.Stream_type

	// TODO: Clean me, deprecated
	// TO Value
	// if to.Stream_type == "email" {
	// 	log.Println("[Debug] MessageOut stream_type is email")
	// 	to.To = in.Data.Person_email
	// } else if to.Stream_type == "sms" {
	// 	log.Println("[Debug] MessageOut stream_type is sms")
	// 	to.To = in.Data.PersonSMS
	// } else if to.Stream_type == "push" {
	// 	log.Println("[Debug] MessageOut stream_type is push")
	// 	to.To = in.Data.PersonPush
	// }

	// Data
	to.Data.Person_Name = in.Data.Person_Name
	to.Data.Date = in.Data.Date

	// dec := json.NewDecoder(strings.NewReader(jsonStream))
	// for {
	// 	var m Message

	// 	if err := dec.Decode(&m); err == io.EOF {
	// 		break
	// 	} else if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("Added Message: \n access_token: %s \n event_code: %s \n stream_type: %s \n END_OF_MESSAGE", m.Access_token, m.Event_code, m.Stream_type)
	// }
}

func FromJSONToObjTest() {

	// TODO: Добавить контролирование входных данных(смотри ТЗ)
	const jsonStream = `
	{
		"access_token": "0d10566b-7e7f-4c17-b2ea-f0e42a4df3c0",
		"event_code": "ispp",
		"stream_type": "email",
		"data": {
		  "person_name": "Иван",
		  "date": "2016-03-03",
		  "person_email": "ivanivanov@gmail.com"
		}
	  }
`

	// Выводим DATA
	in := MessageIn{}
	out := MessageOut{}

	MessageInToMessageToConverter(&in, &out, jsonStream)

	GenerateJSONOut(out)
}

func TestConnectionToDB() {
	CreateQueueFromDB()
}

func TestObjConvertion() {
	println("Run")
	FromJSONToObjTest()
	println("Finished")
}

func main() {
	/*Создаем Send очередь из БД*/
	CreateQueueFromDB()
	/*Получаем очередь*/
	GetQueue()
}
