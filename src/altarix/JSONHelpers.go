package main

import (
	"fmt"
	"log"
	"unicode/utf8"

	"github.com/pquerna/ffjson/ffjson"
)

func ValidateJSONMessage(in *MessageIn) bool {
	if ISDebug {
		log.Println("[Debug] ValidateJSON")
	}

	var EC_lenght int = utf8.RuneCountInString(in.Event_code)

	if EC_lenght == 0 || EC_lenght > 255 {
		log.Fatalln("Invalid JSON SIZE for event_code")
		return false
	}

	if in.Stream_type != "email" && in.Stream_type != "EMAIL" && in.Stream_type != "sms" && in.Stream_type != "SMS" && in.Stream_type != "push" && in.Stream_type != "PUSH" {
		log.Fatalln("Invalid JSON type for Stream_type")
		return false
	}

	// validate on nill. but must be nill, because not pointer
	if (MessageDataIn{}) == in.Data {
		log.Fatalln("Invalid JSON object for Data")
		return false
	}
	return true
}

// Generate JSON func from MessageOut, used for debug
// see https://ashirobokov.wordpress.com/2016/09/22/json-golang-cheat-sheet/
func GenerateJSONOut(in MessageOut) string {
	// Write the buffer
	boolVar, _ := ffjson.Marshal(&in)

	res := string(boolVar)

	if ISDebug {
		fmt.Println(res)
	}

	return res
}

func GenerateJSONIn(in MessageIn) string {
	// Write the buffer
	boolVar, _ := ffjson.Marshal(&in)
	res := string(boolVar)

	if ISDebug {
		fmt.Println(res)
	}

	return res
}
