package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

/*
https://github.com/go-pg/pg
https://github.com/lib/pq
https://github.com/lib/pq/blob/master/listen_example/doc.go
*/

/*
See:http://go-database-sql.org/prepared.html
MySQL               PostgreSQL            Oracle
=====               ==========            ======
WHERE col = ?       WHERE col = $1        WHERE col = :col
VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
*/

const INVALID_VALUE int = -1
const INVALID_VALUE_STRING string = ""
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
	GetUIDFromAccessTokensByToken("38b2cfb8-eb40-fc3d-9a81-49304b21cdb6", &DB)
	GetTokenFromAccessTokensByUID(1, &DB)
	CloseConnectionToDB()
}

/*Функция получает уникальный номер из access_tokens по токену*/
func GetUIDFromAccessTokensByToken(_token string, pdb **sql.DB) int {
	db := *pdb;
	/* Note: передача параметра через

	var req string = "SELECT uid FROM access_tokens WHERE token = ? ";
	stntMessageBody, err = db.Prepare(req)
	stntMessageBody.Query(_token)

	НЕ РАБОТАЕТ. ругается на неправильный синтаксис
	*/
	var req string = "SELECT uid FROM access_tokens WHERE token = $1";
	var stntMessageBody *sql.Stmt;
	stntMessageBody, err = db.Prepare(req);

	printError(file_line());

	//Читаем все значения
	var rows *sql.Rows;
	rows, err = stntMessageBody.Query(_token);

	var UID int;
	UID = INVALID_VALUE;

	for rows.Next() {
		rows.Scan(&UID);
		log.Println("[DEBUG ONLY] Requsted with token: ", _token, " result uid ", UID);
	}
	printError(file_line());

	defer rows.Close();
	defer stntMessageBody.Close();

	return UID
}

/*Функция получает access_tokens по номеру*/
func GetTokenFromAccessTokensByUID(uid int, pdb **sql.DB) string {
	db := *pdb;
	var req string = "SELECT token FROM access_tokens WHERE uid = $1";
	var stntMessageBody *sql.Stmt;
	stntMessageBody, err = db.Prepare(req);

	printError(file_line());

	//Читаем все значения
	var rows *sql.Rows;
	rows, err = stntMessageBody.Query(uid);

	var TOKEN string
	TOKEN = INVALID_VALUE_STRING;

	for rows.Next() {
		rows.Scan(&TOKEN);
		log.Println("[DEBUG ONLY] Requsted with uid: ", uid, " result token ", TOKEN);
	}
	printError(file_line());

	defer rows.Close();
	defer stntMessageBody.Close();

	return TOKEN;
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
