package services

import (
	"testing"
	"time"
)

func Test_Send(t *testing.T) {
	loadEnv()
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	subject := "Unit Test - " + currentTime
	body := "Unit Test - " + currentTime

	err := SendWarnEmail(subject, body)
	if err != nil {
		t.Error("Error sending email:", err)
	}
}
