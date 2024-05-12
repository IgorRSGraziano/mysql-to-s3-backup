package services

import (
	"fmt"
	"time"

	"github.com/IgorRSGraziano/mysql-to-s3-backup/models"
	"gopkg.in/gomail.v2"
)

func SendWarnEmail(title, body string) error {
	config := models.LoadConfig()

	switch {
	case config.SMTP.User == "":
		return fmt.Errorf("SMTP user not set")
	case config.SMTP.Password == "":
		return fmt.Errorf("SMTP password not set")
	case config.SMTP.Host == "":
		return fmt.Errorf("SMTP host not set")
	case config.SMTP.Port == 0:
		return fmt.Errorf("SMTP port not set")
	case config.SMTP.NotificationEmail == "":
		return fmt.Errorf("SMTP notification email not set")
	}

	smtpUser := config.SMTP.User
	smtpPass := config.SMTP.Password
	smtpHost := config.SMTP.Host
	smtpPort := config.SMTP.Port

	smtpSecure := config.SMTP.Secure

	warnEmail := config.SMTP.NotificationEmail

	errChan := make(chan error, 1)

	var err error

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
