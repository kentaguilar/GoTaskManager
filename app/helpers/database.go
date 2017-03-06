package helper

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func OpenDb() *sql.DB{
	db, err = sql.Open("mysql", "root:@/gotask")
    if err != nil{
        panic(err.Error())
    }

    return db
}

func CloseDb(db *sql.DB) {
	defer db.Close()
}