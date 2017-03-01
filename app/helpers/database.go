package helper

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	// "golang.org/x/crypto/bcrypt"
)

var db *sql.DB
var err error

func OpenDb() *sql.DB{
	db, err = sql.Open("mysql", "root:@/godb")
    if err != nil{
        panic(err.Error())
    }
    // defer db.Close()

    // err = db.Ping()
    // if err != nil{
    //     panic(err.Error())
    // }

    return db
}

func CloseDb(db *sql.DB) {
	defer db.Close()
}