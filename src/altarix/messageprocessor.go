package main

import (
	// "database/sql"
	// "fmt"
	// "net/http"
	// "strconv"

	//_ "github.com/mattn/go-sqlite3"

	// _ "github.com/go-sql-driver/mysql"

	// JSON, for object parsing
	//"encoding/json"
	// see
	/*https://github.com/pquerna/ffjson  - it faster
	 */
	"github.com/pquerna/ffjson/ffjson"

	// "fmt"
	//"io"
	"log"
	// "strings"
	// "unicode/utf8"
)

// План
/*
Finished:
1) Разбор сообщения (тут лучше всего сделать через template) __DONE__

template <class DataT>
struct Message
{
	String access_token, event_code, stream_type
	DataT data;
}

struct DataIn
{
	...
}

struct DataOut
{
	...
}

Use:
Message<DataIn> in;
Message<DataOut> out;

1) Research on message broker and implement support of postgres+research sql structure __DONE__
2) Сделать функцию которая будет составлять новое сообщение  __DONE__
3) Добавляем контролирование входных данных __DONE__
4) Генерация JSON'a - debug only __DONE__
5) Пишем структуру БД __DONE__
6) Инит очередей __DONE__
7) Заполняем очередь пустышками __DONE__
8) Инит Postgres __DONE__
9) Заполняем очередь из БД __DONE__

// TODO On today:

Do:
9) Делаем запись в БД(Лучше всего тут так же использовать очередь) - Очередь отменяется, т.к тут пайалайн работы такой не прокатит. 
9.5) Переписываем как демона.append


10) Ресерчим Dockerfile
11) Реализация Dockerfile
11) Юнит-тесты 
11) Сдача
*/

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
	log.Printf("JSON %s", jsonStream)

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
			log.Println("[Debug] stream_type is email")
			// fmt.Println(in.Data.Person_email)
		}
		to.To = in.Data.Person_email
	} else if in.Stream_type == "sms" || in.Stream_type == "SMS" {
		if ISDebug {
			log.Println("[Debug] stream_type is sms")
			// fmt.Println(in.Data.PersonSMS)
		}
		to.To = in.Data.PersonSMS
	} else if in.Stream_type == "push" || in.Stream_type == "PUSH" {
		if ISDebug {
			log.Println("[Debug] stream_type is push")
			// fmt.Println(in.Data.PersonPush)
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
	// RM_Send();
	// RM_Receive();

	// in := MessageIn{}
	// in.Access_token="0d10566b-7e7f-4c17-b2ea-f0e42a4df3c0";
	// in.Event_code="ispp";
	// in.Stream_type="email";
	// in.Data.Date="2016-03-03";
	// in.Data.Person_Name="Иван";
	// in.Data.Person_email="ivanivanov@gmail.com";

	// RM_Send("hello", GenerateJSONIn(in));

	/*Создаем Send очередь из БД*/
	CreateQueueFromDB()
	/*Получаем очередь*/
	GetQueue()
}
