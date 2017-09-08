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

	"fmt"
	//"io"
	"log"
	// "strings"
	"unicode/utf8"
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

2) Сделать функцию которая будет составлять новое сообщение  __DONE__
3) Добавляем контролирование входных данных __DONE__
4) Генерация JSON'a - debug only __DONE__

// TODO On today:
1) Research on message broker and implement support of postgres+research sql structure 

Do:
-) Добавить валидацию в сообщении. Поле access_token по паттерну. Пока не понятно, может ли оно отличаться. Уточнить
4) Инит очередей
5) Заполняем очередь пустышками
6) Инит Postgres
7) Пишем структуру БД
8) Заполняем очередь из БД
9) Делаем запись в БД(Лучше всего тут так же использовать очередь)
10) Ресерчим Dockerfile
11) Реализация
11) Юнит-тесты
11) Сдача
*/

// JSON
// https://golang.org/pkg/encoding/json/#pkg-examples

var err error

func printError() {
	if err != nil {
		log.Fatalln(err.Error)
	}
}

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

func ValidateJSONMessage(in *MessageIn) bool {
	log.Println("[Debug] ValidateJSON")

	var EC_lenght int = utf8.RuneCountInString(in.Event_code)

	if EC_lenght == 0 || EC_lenght > 255 {
		log.Fatalln("Invalid JSON SIZE for event_code")
		return false
	}

	if in.Stream_type != "email" && in.Stream_type != "sms" && in.Stream_type != "push" {
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
func GenerateJSON(in MessageOut) {
	// Write the buffer
	boolVar, _ := ffjson.Marshal(&in)
	fmt.Println(string(boolVar))

	// TODO: Add write to bd
}

func MessageInToMessageToConverter(in *MessageIn, to *MessageOut, jsonStream string) {
	ffjson.Unmarshal([]byte(jsonStream), &in)

	// debug
	// {
	// 	fmt.Println(in)
	// 	fmt.Println(in.Data.Person_Name)
	// }

	if ValidateJSONMessage(in) == false {
		log.Println("Failed validation")
		return
	}

	if in.Stream_type == "email" {
		log.Println("[Debug] stream_type is email")
		// fmt.Println(in.Data.Person_email)
	} else if in.Stream_type == "sms" {
		log.Println("[Debug] stream_type is sms")
		// fmt.Println(in.Data.PersonSMS)
	} else if in.Stream_type == "push" {
		log.Println("[Debug] stream_type is push")
		// fmt.Println(in.Data.PersonPush)
	}

	to.Access_token = in.Access_token
	to.Event_code = in.Event_code
	to.Stream_type = in.Stream_type

	// TO Value
	if to.Stream_type == "email" {
		log.Println("[Debug] MessageOut stream_type is email")
		to.To = in.Data.Person_email
	} else if to.Stream_type == "sms" {
		log.Println("[Debug] MessageOut stream_type is sms")
		to.To = in.Data.PersonSMS
	} else if to.Stream_type == "push" {
		log.Println("[Debug] MessageOut stream_type is push")
		to.To = in.Data.PersonPush
	}

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

func FromJSONToObj() {

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

	GenerateJSON(out)
}

func main() {

	//{
	//	if _, err = os.Stat("./bulletin.db"); os.IsNotExist(err) {
	//		println("database ./bulletin.db doesn't exist")
	//		return
	//	}
	//}

	println("Run")
	FromJSONToObj()
	println("Finished")

	// defer stmtCateg.Close() // Close the statement when we leave main() / the program terminates
}

// var DB *sql.DB
// var stmtCateg *sql.Stmt //List of categories
// var stntAdds *sql.Stmt  // list of all adds by categoryID
//var stntMessageBody *sql.Stmt // list of all adds by categoryID
