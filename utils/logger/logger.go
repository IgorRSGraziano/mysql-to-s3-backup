package logger

import (
	"fmt"
	"time"

	"github.com/IgorRSGraziano/mysql-to-s3-backup/services"
)

const (
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorWhite  = "\033[37m"
)

func prepareMessage(message string, tag string, color string) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintf("%s[%s] %s ~ %s%s\033[0m", color, timestamp, tag, message, color)
}

func Info(message string) {
	fmt.Println(prepareMessage(message, "INFO", ColorGreen))
}

func Error(message string) {
	fmt.Println(prepareMessage(message, "ERROR", ColorRed))
}

func Warning(message string) {
	fmt.Println(prepareMessage(message, "WARN", ColorYellow))
}

func Success(message string) {
	fmt.Println(prepareMessage(message, "SUCCESS", ColorBlue))
}

func Fatal(message string) {
	fmt.Println(prepareMessage(message, "FATAL", ColorRed))
	err := services.SendWarnEmail("Backup fatal error", message)
	if err != nil {
		Error("Error sending email:" + err.Error())
	}
	panic(message)
}
