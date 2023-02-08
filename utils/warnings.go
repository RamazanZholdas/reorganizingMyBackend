package utils

import (
	"log"
	"os"
)

func CreateLogFile() {
	file, err := os.OpenFile("./../../logs/ProgramLogs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()
}
