package define

func DatabaseTemplate() string {
    return `package define

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

var Db *sql.DB

func init () {
    Db, _ = sql.Open("mysql", Connection)
}
`
}