package db

import (
	define "api/app/util/define"
	"log"
	l "log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() *gorm.DB {
	DBMS := "mysql"
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	PROTOCOL := "tcp(" + os.Getenv("MYSQL_HOST") + ":3306)"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/orange?charset=utf8mb4&parseTime=True&loc=Asia%2FTokyo"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	// Query output destination
	ymd := time.Now().Format("2006-01-02")
	path := "/go/src/api/log/sql_" + ymd + ".log"
	logfile, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	// Set as log output destination
	l.SetOutput(logfile)
	l.SetFlags(l.Ldate | l.Ltime | l.Llongfile)

	db.LogMode(true)
	db.SetLogger(log.New(logfile, "", 0))

	return db
}

// クエリ実行
func Exec(sql string) {
	db := Connect()

	tx := db.Begin()

	if err := db.Exec(sql).Error; err != nil {
		log.Println(sql)
		log.Println(err)
		os.Exit(1)
	}

	tx.Commit()
	defer db.Close()
}

func Handler(errs []error) int {
	// check Not Found
	if len(errs) == 1 {
		if gorm.IsRecordNotFoundError(errs[0]) {
			return define.SUCCESS
		}
	}

	if len(errs) > 0 {
		for _, db_err := range errs {
			log.Println(db_err)
		}
		return define.DB_SQL_FAILURE
	}
	return define.SUCCESS
}
