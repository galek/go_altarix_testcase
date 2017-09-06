package main

import (
	// "database/sql"
	// "fmt"
	// "net/http"
	// "strconv"

	//_ "github.com/mattn/go-sqlite3"

	// _ "github.com/go-sql-driver/mysql"

	// JSON, for object parsing
	"encoding/json"

	"fmt"
	// "io"
	// "log"
	// "strings"
)

// План
/*
1) Разбор сообщения (тут лучше всего сделать через template)

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


2) Сделать функцию которая будет составлять новое сообщение
3) Добавляем контролирование входных данных
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

	type MessageData struct {
		Person_Name, Date, Person_email string
	}

	type Message struct {
		Access_token, Event_code, Stream_type string
		Data                                  MessageData
	}

	// Выводим DATA
	res := Message{}
	json.Unmarshal([]byte(jsonStream), &res)
	fmt.Println(res)
	fmt.Println(res.Data.Person_Name)

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

//========================================
// func CategoriesHandler(w http.ResponseWriter, r *http.Request) {
// 	MakeCookiesGreatAgain(w, r)
// 	Header(w)
// 	CategoriesShow(w)
// 	println("CategoriesHandler: with DB ", r.FormValue("id"))
// 	Footer(w)
// }

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
var err error

func printError() {
	if err != nil {
		println("Error: with DB ", err.Error())
	}
}
