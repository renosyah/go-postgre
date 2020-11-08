package auth

import "database/sql"

var (
	dbPool *sql.DB
	option *Option
)

func Init(db *sql.DB, opt *Option) {
	dbPool = db
	option = opt
}
