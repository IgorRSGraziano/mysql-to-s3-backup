package services

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"gopkg.in/gomail.v2"
)

func SendWarnEmail(title, body string) error {
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))

	if err != nil {
		fmt.Println("Error parsing SMTP_PORT:", err)
		smtpPort = 587
	}

	smtpSecure := os.Getenv("SMTP_SECURE") == "true"

	warnEmail := os.Getenv("WARN_EMAIL")

	errChan := make(chan error, 1)

	go func() {
		d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
		d.SSL = smtpSecure

		m := gomail.NewMessage()
		m.SetHeader("From", smtpUser)
		m.SetHeader("To", warnEmail)
		m.SetHeader("Subject", title)
		m.SetBody("text/html", body)

		errChan <- d.DialAndSend(m)
	}()

	timeoutSeconds := 25

	select {
	case err = <-errChan:
		return err
	case <-time.After(time.Duration(timeoutSeconds) * time.Second):
		return fmt.Errorf("timeout after %d seconds", timeoutSeconds)
	}
}
