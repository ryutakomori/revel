package util

import (
	"io"
	l "log"
	"os"
	"time"
)

var logfile *os.File

func Intialize() {
	ymd := time.Now().Format("2006-01-02")

	path := "/go/src/api/log/api_" + ymd + ".log"

	logfile, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic("cannnot open go.log:" + err.Error())
	}
	// io.MultiWriteで標準出力とファイルの両方を束ねてlogの出力先に設定する
	l.SetOutput(io.MultiWriter(logfile, os.Stdout))
	l.SetOutput(logfile)
	l.SetFlags(l.Ldate | l.Ltime | l.Llongfile)
}

func Println(v ...interface{}) {
	l.Println(v)
}

func Close() {
	logfile.Close()
}
