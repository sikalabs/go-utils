package mail

import (
	"fmt"
	"testing"
	"time"

	"github.com/spf13/viper"
)

func TestSendEmailWithAttachment_with_body(t *testing.T) {
	smtpHost, smtpPort, smtpUsername, smtpPassword, emailFrom, emailTo := getSmtpConfigFromEnv(t)

	fileName := fmt.Sprintf("hello-%d.txt", time.Now().Unix())
	fileData := []byte(fmt.Sprintf("Hello World! %d\n", time.Now().Unix()))

	err := SendEmailWithAttachment(
		smtpHost,
		smtpPort,
		smtpUsername,
		smtpPassword,
		emailFrom,
		emailTo,
		"go test: TestSendEmailWithAttachment_with_body",
		"text body\nO.\n\n",
		fileName,
		fileData,
	)

	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func TestSendEmailWithAttachment_without_body(t *testing.T) {
	smtpHost, smtpPort, smtpUsername, smtpPassword, emailFrom, emailTo := getSmtpConfigFromEnv(t)

	fileName := fmt.Sprintf("hello-%d.txt", time.Now().Unix())
	fileData := []byte(fmt.Sprintf("Hello World! %d\n", time.Now().Unix()))

	err := SendEmailWithAttachment(
		smtpHost,
		smtpPort,
		smtpUsername,
		smtpPassword,
		emailFrom,
		emailTo,
		"go test: TestSendEmailWithAttachment_without_body",
		"",
		fileName,
		fileData,
	)

	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func getSmtpConfigFromEnv(t *testing.T) (string, int, string, string, string, string) {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			t.Fatalf("Error reading .env file: %v\n", err)
		}
	}

	viper.AutomaticEnv()

	viper.BindEnv("SMTP_HOST")
	viper.BindEnv("SMTP_PORT")
	viper.BindEnv("SMTP_USERNAME")
	viper.BindEnv("SMTP_PASSWORD")

	smtpHost := viper.GetString("SMTP_HOST")
	if smtpHost == "" {
		t.Fatal("SMTP_HOST not set")
	}

	smtpPort := viper.GetInt("SMTP_PORT")
	if smtpPort == 0 {
		t.Fatal("SMTP_PORT not set")
	}

	smtpUsername := viper.GetString("SMTP_USERNAME")
	if smtpUsername == "" {
		t.Fatal("SMTP_USERNAME not set")
	}

	smtpPassword := viper.GetString("SMTP_PASSWORD")
	if smtpPassword == "" {
		t.Fatal("SMTP_PASSWORD not set")
	}

	emailFrom := viper.GetString("EMAIL_FROM")
	if emailFrom == "" {
		t.Fatal("EMAIL_FROM not set")
	}

	emailTo := viper.GetString("EMAIL_TO")
	if emailTo == "" {
		t.Fatal("EMAIL_TO not set")
	}
	return smtpHost, smtpPort, smtpUsername, smtpPassword, emailFrom, emailTo
}
