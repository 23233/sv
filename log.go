package sv

import (
	"io"
	"log"
	"os"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	errFile, err := os.OpenFile("simple_valid.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("open log file fail：", err)
	}
	Info = log.New(os.Stdout, "[simple_valid] Info:", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "[simple_valid] Warning:", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stderr, errFile), "[simple_valid] Error:", log.Ldate|log.Ltime|log.Lshortfile)

}
