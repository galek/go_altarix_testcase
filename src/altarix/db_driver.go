package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

/*
https://github.com/go-pg/pg
https://github.com/lib/pq
https://github.com/lib/pq/blob/master/listen_example/doc.go
*/

const DB_CONNECT_STRING = "host=localhost port=5432 user=postgres password=postgres dbname=altarix sslmode=disable"

var DB *sql.DB

func OpenConnectionToDB() {
	DB, err = sql.Open("postgres", DB_CONNECT_STRING)
	if err != nil {
		log.Println("Database opening error -->%v\n", err)
		panic("Database error")
	}

	printError(file_line())
}

func CloseConnectionToDB() {
	defer DB.Close()
}

func connectionToDB() {
	OpenConnectionToDB()
	// init_database(&db)
	GetUIDFromAccessTokensByToken("'38b2cfb8-eb40-fc3d-9a81-49304b21cdb6'", &DB)
	//CloseConnectionToDB()
}

func GetUIDFromAccessTokensByToken(_token string, pdb **sql.DB) {
	db := *pdb
	var req string = "SELECT uid FROM access_tokens WHERE token = ?"
	var stntMessageBody *sql.Stmt
	stntMessageBody, err = db.Prepare(req)

	printError(file_line())

	println("token: %s", _token)

	fmt.Println("finished")

	stntMessageBody.Query(_token)

	//Читаем все значения
	// var rows *sql.Rows
	// rows, err = stntMessageBody.Query(_token)

	// var UID int

	// for rows.Next() {
	// 	rows.Scan(&UID)
	// 	log.Println("[DEBUG ONLY] %s %i", _token, UID)
	// }
	// printError(file_line())

	// defer rows.Close()
	// defer stntMessageBody.Close()
	// return 0;
}

/*-----------------------------------------------------------------------------*/
func init_database(pdb **sql.DB) {

	// db := *pdb

	// init_db_strings := []string{
	// 	"DROP SCHEMA IF EXISTS sb CASCADE;",
	// 	"CREATE SCHEMA sb;",
	// 	//be careful - next multiline string is quoted by backquote symbol
	// 	`CREATE TABLE sb.test_data(
	// 		 id serial,
	// 		 device_id integer not null,
	// 		 parameter_id integer not null,
	// 		 value varchar(100),
	// 		 event_ctime timestamp default current_timestamp,
	// 		 constraint id_pk primary key (id));`}

	// for _, qstr := range init_db_strings {
	// 	_, err := db.Exec(qstr)

	// 	if err != nil {
	// 		fmt.Printf("Database init error -->%v\n", err)
	// 		panic("Query error")
	// 	}
	// }
	// fmt.Println("Database rebuilded successfully")
}
