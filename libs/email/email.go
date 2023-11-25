package email

import (
	"fmt"
	"net/smtp"
)

var from string
var password string
var smtpHost string
var smtpPort string

func init() {
	// TODO: 读取配置文件
	from = "122385930@qq.com"
	password = "otojckznxjrlbigf"
	smtpHost = "smtp.qq.com"
	smtpPort = ":587"
}

// sendEmail 函数用于发送邮件通知
func Send(to, subject, body, contentType string) error {
	auth := smtp.PlainAuth("", from, password, smtpHost)

	msg := []byte("To: " + to + "\r\n" +
		"From: " + from + "<" + from + ">" + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: " + contentType + "; charset=UTF-8" + "\r\n" + "\r\n" +
		body)

	if err := smtp.SendMail(smtpHost+smtpPort, auth, from, []string{to}, msg); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
