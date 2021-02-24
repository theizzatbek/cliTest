package cmd

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/smtp"
	"strings"
	"time"
)

type server struct {
	Host     string
	Port     string
	User     string
	Password string
}

func (s *server) Address() string {
	return s.Host + ":" + s.Port
}

func Send(receiver, msg string) error {

	file, _ := ioutil.ReadFile("./conf/email.json")

	server := server{}

	_ = json.Unmarshal(file, &server)

	//fmt.Println(server)
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

	err := smtp.SendMail(server.Address(), auth, from, to, buf.Bytes())
	if err != nil {
		return err
	}
	return nil
}
