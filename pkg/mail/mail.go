package mail

import (
	"bytes"
	_ "embed"
	"encoding/base64"
	"fmt"
	"mime/multipart"
	"net/smtp"
	"net/textproto"
	"path/filepath"
)

func SendEmailWithAttachment(
	smtpHost string,
	smtpPort int,
	smtpUsername string,
	smtpPassword string,
	emailFrom string,
	emailTo string,
	subject string,
	body string,
	filename string,
	fileData []byte,
) error {
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

	bodyBuf := new(bytes.Buffer)
	writer := multipart.NewWriter(bodyBuf)

	// Create the plain text part
	if body != "" {
		textPart, err := writer.CreatePart(textproto.MIMEHeader{
			"Content-Type":              []string{"text/plain; charset=UTF-8"},
			"Content-Transfer-Encoding": []string{"7bit"},
		})
		if err != nil {
			panic(err)
		}
		textPart.Write([]byte(body))
	}

	// Create the attachment part
	attachmentPart, err := writer.CreatePart(textproto.MIMEHeader{
		"Content-Type":              []string{"application/octet-stream"},
		"Content-Transfer-Encoding": []string{"base64"},
		"Content-Disposition":       []string{fmt.Sprintf(`attachment; filename="%s"`, filepath.Base(filename))},
	})
	if err != nil {
		panic(err)
	}
	attachmentPart.Write([]byte(base64.StdEncoding.EncodeToString(fileData)))

	writer.Close()

	message := []byte("To: " + emailTo + "\r\n" +
		"From: " + emailFrom + "\r\n" +
		"Subject: " + subject + "\r\n")
	message = append(message, []byte(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\r\n\r\n", writer.Boundary()))...)
	message = append(message, bodyBuf.Bytes()...)

	return smtp.SendMail(fmt.Sprintf("%s:%d", smtpHost, smtpPort), auth, emailFrom, []string{emailTo}, message)
}
