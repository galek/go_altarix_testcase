package main

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/lib/pq"
)

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

func OpenConnectionToDB(pdb **sql.DB) {
	DB, err = sql.Open("postgres", DB_CONNECT_STRING)
	if err != nil {
		log.Println("Database opening error -->%v\n", err)
		panic("Database error")
	}
	// закрытие сделаем в CloseConnectionToDB
	// defer DB.Close()

	printError(file_line())
}

func CloseConnectionToDB(pdb **sql.DB) {
	db := *pdb
	err = db.Close()
	printError(file_line())
}

func IsInValidConnectionDB(pdb **sql.DB) error {
	db := *pdb
	if pdb == nil {
		return errors.New("Pointer to connection is dead")
	}
	errL := db.Ping()
	if errL != nil {
		return errors.New("Connection is lost")
	}

	return errL
}

/*
Создаем очередь из БД
*/
func CreateQueueFromDB() {
	OpenConnectionToDB(&DB)
	GetObjectFromResultTable(&DB)
	CloseConnectionToDB(&DB)
}

/*
Получаем очередь и разбираем ее
*/
func GetQueue() {
	OpenConnectionToDB(&DB)
	RM_Receive("obj")
	CloseConnectionToDB(&DB)
}

/**/
func UTIL_GetUIDByString(_tokenName string, UID_NAME string, _tokenValue string, _dbName string, pdb **sql.DB) int {
	printError(file_line())
	/*Validation layer*/
	if IsInValidConnectionDB(pdb) != nil {
		log.Printf("InValid connection.")
		OpenConnectionToDB(pdb)
	}
	printError(file_line())
	
	db := *pdb
	var req string = "SELECT " + UID_NAME + " FROM " + _dbName + " WHERE " + _tokenName + " = $1"
	if ISDebug {
		log.Println("req: ", req)
	}

	var stntMessageBody *sql.Stmt
	stntMessageBody, err = db.Prepare(req)

	printError(file_line())

	//Читаем все значения
	var rows *sql.Rows
	rows, err = stntMessageBody.Query(_tokenValue)

	var UID int
	UID = INVALID_VALUE

	for rows.Next() {
		rows.Scan(&UID)
		if ISDebug {
			log.Println("[DEBUG ONLY] Requsted with token: ", _tokenValue, " result uid ", UID)
		}
	}
	printError(file_line())

	defer rows.Close()
	defer stntMessageBody.Close()

	return UID
}

func UTIL_GetStringByUID(_tokenName string, UID_NAME string, UID_VALUE int, _dbName string, pdb **sql.DB) string {
	printError(file_line())
	/*Validation layer*/
	if IsInValidConnectionDB(pdb) != nil {
		log.Printf("InValid connection.")
		OpenConnectionToDB(pdb)
	}
	printError(file_line())

	db := *pdb
	var req string = "SELECT " + _tokenName + " FROM " + _dbName + " WHERE " + UID_NAME + " = $1"
	if ISDebug {
		log.Println("req: ", req)
	}

	var stntMessageBody *sql.Stmt
	stntMessageBody, err = db.Prepare(req)

	printError(file_line())

	//Читаем все значения
	var rows *sql.Rows
	rows, err = stntMessageBody.Query(UID_VALUE)

	var UID string
	UID = INVALID_VALUE_STRING

	for rows.Next() {
		rows.Scan(&UID)
		if ISDebug {
			log.Println("[DEBUG ONLY] Requsted with token: ", _tokenName, " result uid ", UID)
		}
	}
	printError(file_line())

	defer rows.Close()
	defer stntMessageBody.Close()

	return UID
}

/*Функция получает уникальный номер из access_tokens по токену*/
func GetUIDFromAccessTokensByToken(_token string, pdb **sql.DB) int {
	return UTIL_GetUIDByString("token", "uid", _token, "access_tokens", pdb)
}

/*Функция получает уникальный номер из access_tokens по токену*/
func GetUIDFromEventCodesByUID(_token string, pdb **sql.DB) int {
	return UTIL_GetUIDByString("descr", "uid", _token, "Event_codes", pdb)
}
func GetUIDFromStreamTypesByUID(_token string, pdb **sql.DB) int {
	return UTIL_GetUIDByString("descr", "uid", _token, "stream_types", pdb)
}
func GetUIDFromNamesByUID(_token string, pdb **sql.DB) int {
	return UTIL_GetUIDByString("name", "uid", _token, "id_names", pdb)
}

/*Функция получает access_tokens по номеру*/
func GetTokenFromAccessTokensByUID(uid int, pdb **sql.DB) string {
	return UTIL_GetStringByUID("token", "uid", uid, "access_tokens", pdb)
}

/*Функция получает Event_code по номеру*/
func GetTokenFromEventCodesByUID(uid int, pdb **sql.DB) string {
	return UTIL_GetStringByUID("descr", "uid", uid, "Event_codes", pdb)
}

/*Функция получает Event_code по номеру*/
func GetTokenFromStreamTypesByUID(uid int, pdb **sql.DB) string {
	return UTIL_GetStringByUID("descr", "uid", uid, "stream_types", pdb)
}

/*Функция получает Имя по UID*/
func GetTokenFromPersonNameByUID(uid int, pdb **sql.DB) string {
	return UTIL_GetStringByUID("name", "uid", uid, "id_names", pdb)
}

/*Функция получает email по UID*/
func GetTokenFromPersonEmailByUID(uid int, pdb **sql.DB) string {
	return UTIL_GetStringByUID("email", "uid", uid, "uuid_email", pdb)
}

/*Функция получает Номер телефона по UID*/
func GetTokenFromPersonSMSByUID(uid int, pdb **sql.DB) string {
	return UTIL_GetStringByUID("tel_number", "uid", uid, "uuid_sms", pdb)
}

/*Получаем готовый к употреблению объект из ResultTable*/
func GetObjectFromResultTable(pdb **sql.DB) {
	printError(file_line())
	/*Validation layer*/
	if IsInValidConnectionDB(pdb) != nil {
		log.Printf("InValid connection.")
		OpenConnectionToDB(pdb)
	}
	printError(file_line())

	db := *pdb
	var req string = "SELECT " + "*" + " FROM " + " resulttable"
	if ISDebug {
		log.Println("req: ", req)
	}

	var stntMessageBody *sql.Stmt
	stntMessageBody, err = db.Prepare(req)

	printError(file_line())

	//Читаем все значения
	var rows *sql.Rows
	rows, err = stntMessageBody.Query()

	// Allocate size on 0 elements
	// bks := make([]*MessageIn, 0)
	// Not allocate
	var bks []MessageIn // == nil

	var AT_UID int = INVALID_VALUE
	var EC_UID int = INVALID_VALUE
	var ST_UID int = INVALID_VALUE
	var PN_UID int = INVALID_VALUE
	var PNE_UID int = INVALID_VALUE
	var PN_PHONE_UID int = INVALID_VALUE
	var ROW_IT int = 1

	for rows.Next() {
		/*
		 */
		if ISDebug {
			log.Println("ROW_IT ", ROW_IT)
		}

		// for Allocate solution
		// bk := new(MessageIn)
		bk := MessageIn{}

		err = rows.Scan(&AT_UID, &EC_UID, &ST_UID, &PN_UID, &PNE_UID, &PN_PHONE_UID, &bk.Data.Date)

		if ISDebug {
			log.Println("[DEBUG ONLY] ",
				AT_UID,
				EC_UID,
				ST_UID,
				PN_UID,
				PNE_UID,
				PN_PHONE_UID,
				bk.Data.Date)
		}

		printError(file_line())

		// Вот тут на самом деле можно все упростить и поместить в 1 табличку. Но я хотел достичь нормализации, поэтому все так усложнено.
		/*Добавим еще одну таблицу в последствии*/

		/*Convertation part*/
		bk.Access_token = GetTokenFromAccessTokensByUID(AT_UID, pdb)
		bk.Event_code = GetTokenFromEventCodesByUID(EC_UID, pdb)
		bk.Stream_type = GetTokenFromStreamTypesByUID(ST_UID, pdb)
		bk.Data.Person_Name = GetTokenFromPersonNameByUID(PN_UID, pdb)
		bk.Data.Person_email = GetTokenFromPersonEmailByUID(PN_UID, pdb)
		bk.Data.PersonSMS = GetTokenFromPersonSMSByUID(PN_PHONE_UID, pdb)

		if ISDebug {
			log.Printf("[DEBUG ONLY] %s, %s, %s, %s, %s, %s, %s \n", bk.Access_token, bk.Event_code, bk.Stream_type, bk.Data.Person_Name, bk.Data.Person_email, bk.Data.PersonSMS, bk.Data.Date)
		}

		bks = append(bks, bk)

		if ISDebug {
			ROW_IT++
		}
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, bk := range bks {

		/*Send to queue*/
		RM_Send("obj", GenerateJSONIn(bk))

		if ISDebug {
			log.Printf("[DEBUG ONLY] %s, %s, %s, %s, %s, %s, %s \n", bk.Access_token, bk.Event_code, bk.Stream_type, bk.Data.Person_Name, bk.Data.Person_email, bk.Data.PersonSMS, bk.Data.Date)
		}
	}

}

/*-----------------------------------------------------------------------------*/
/*
https://www.compose.com/articles/going-from-postgresql-rows-to-rabbitmq-messages/
*/
func WriteMessageToBD(_mess *MessageOut, pdb **sql.DB) {
	printError(file_line())
	/*Validation layer*/
	if IsInValidConnectionDB(pdb) != nil {
		log.Printf("InValid connection.")
		OpenConnectionToDB(pdb)
	}
	printError(file_line())

	db := *pdb
	var req string = "INSERT INTO public.totable (access_token, event_token, stream_type, person_name, person_to, person_date) VALUES($1, $2, $3, $4, $5, $6)"

	if ISDebug {
		log.Println("req: ", req)
	}

	var stmt *sql.Stmt
	stmt, err = db.Prepare(req)

	if ISDebug {
		log.Println("OpenConnection to DB -PRepare")
	}

	printError(file_line())

	stmt.Query(
		GetUIDFromAccessTokensByToken(_mess.Access_token, pdb),
		GetUIDFromEventCodesByUID(_mess.Event_code, pdb),
		GetUIDFromStreamTypesByUID(_mess.Stream_type, pdb),
		GetUIDFromNamesByUID(_mess.Data.Person_Name, pdb),
		_mess.To,
		_mess.Data.Date)

	printError(file_line())

	stmt.Close()
}
