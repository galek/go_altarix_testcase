package main

import (
	// "database/sql"
	// "fmt"
	// "net/http"
	// "strconv"

	//_ "github.com/mattn/go-sqlite3"

	// _ "github.com/go-sql-driver/mysql"
)

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

	println("test");

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