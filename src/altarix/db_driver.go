package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const DB_CONNECT_STRING = 
"host=localhost port=5432 user=postgres password=postgres dbname=altarix sslmode=disable"

func connectionToDB() {
	db, err := sql.Open("postgres", DB_CONNECT_STRING)
    defer db.Close()

    if err != nil {
        fmt.Printf("Database opening error -->%v\n", err)
        panic("Database error")
	}
	
	init_database(&db)
}

/*-----------------------------------------------------------------------------*/
func init_database(pdb **sql.DB) {
	
		db := *pdb
	
		init_db_strings := []string{
			"DROP SCHEMA IF EXISTS sb CASCADE;",
			"CREATE SCHEMA sb;",
			//be careful - next multiline string is quoted by backquote symbol
			`CREATE TABLE sb.test_data(
			 id serial,
			 device_id integer not null,
			 parameter_id integer not null,
			 value varchar(100),
			 event_ctime timestamp default current_timestamp,
			 constraint id_pk primary key (id));`}
	
		for _, qstr := range init_db_strings {
			_, err := db.Exec(qstr)
	
			if err != nil {
				fmt.Printf("Database init error -->%v\n", err)
				panic("Query error")
			}
		}
		fmt.Println("Database rebuilded successfully")
	}