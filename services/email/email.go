package cmd

import (
	"bytes"
	"cliTest/config"
	"encoding/base64"
	"fmt"
	"net/smtp"
	"strings"
	"time"
)

func address(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}

func Send(receiver, msg string) error {

	server := config.GetInstance().Email
	from := server.User
	password := server.Password
	to := []string{
		receiver + "21",
	}

	buf := bytes.NewBuffer(nil)
	buf.WriteString("From: " + from + "\r\n")
	t := time.Now()
	buf.WriteString("Date: " + t.Format(time.RFC1123Z) + "\r\n")
	buf.WriteString("To: " + strings.Join(to, ",") + "\r\n")
	var coder = base64.StdEncoding
	var subject = "=?UTF-8?B?" + coder.EncodeToString([]byte("Notification")) + "?="
	buf.WriteString("Subject: " + subject + "\r\n")
	buf.WriteString("MIME-Version: 1.0\r\n")
	buf.WriteString(fmt.Sprintf("Content-Type: %s; charset=utf-8\r\n\r\n", "text/plain"))
	buf.WriteString(msg)
	buf.WriteString("\r\n")

	auth := smtp.PlainAuth("", from, password, server.Host)

	err := smtp.SendMail(
		address(server.Host, server.Port),
		auth, from, to, buf.Bytes())
	if err != nil {
		return err
	}
	return nil
}
