package customlogger

import (
	"log"
	"os"
	"sync"
)

type cLogger struct {
	filename string
	*log.Logger
}

var logger *cLogger
var once sync.Once

func GetInstance() *cLogger {
	once.Do(func() {
		logger = createLogger("info.log")
	})
	return logger
}

func createLogger(fname string) *cLogger {
	file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	return &cLogger{
		filename: fname,
		Logger:   log.New(file, "[HeroesApp]", log.Ltime|log.Ldate|log.LstdFlags),
	}
}
