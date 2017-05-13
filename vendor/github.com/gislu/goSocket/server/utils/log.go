package utils

import (
	"log"
	"os"
)

// Log function

func LogErr(v ...interface{}) {

	logfile := os.Stdout
	log.Println(v...)
	logger := log.New(logfile, "\r\n", log.Llongfile|log.Ldate|log.Ltime)
	logger.SetPrefix("[Error]")
	logger.Println(v...)
	defer logfile.Close()
}

func Log(v ...interface{}) {

	logfile := os.Stdout
	log.Println(v...)
	logger := log.New(logfile, "\r\n", log.Ldate|log.Ltime)
	logger.SetPrefix("[Info]")
	logger.Println(v...)
	defer logfile.Close()
}
func Logf(format string, v ...interface{}) {

	logfile := os.Stdout
	log.Printf(format, v...)
	logger := log.New(logfile, "\r\n", log.Ldate|log.Ltime)
	logger.SetPrefix("[Info]")
	logger.Printf(format, v...)
	defer logfile.Close()
}

func LogDebug(v ...interface{}) {
	logfile := os.Stdout
	log.Println(v...)
	logger := log.New(logfile, "\r\n", log.Ldate|log.Ltime)
	logger.SetPrefix("[Debug]")
	logger.Println(v...)
	defer logfile.Close()
}

func CheckError(err error) {
	if err != nil {
		LogErr(os.Stderr, "Fatal error: %s", err.Error())
	}
}
