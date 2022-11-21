package main

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

// Получатели
var Rcpt = []string{
	"recipient1@tomail.com",
	"recipient2@tomail.com",
}

type Mail struct {
    Sender  string
    To      []string
    Subject string
    Body    string
}

func main() {
    senderUser := "testmailganer1@rambler.ru"

    to := []string{
        "testmailganer1@rambler.ru",
    }

    password := "Password123qwe"

    subject := "Список подписчиков"
    body := `<p>Список <b>подписчиков</b></p>`
	for i := range to {
		body += "<li>" + to[i] + "</li>"
	}

    request := Mail{
        Sender:  senderUser,
        To:      to,
        Subject: subject,
        Body:    body,
    }

    addr := "smtp.rambler.ru:587"
    host := "smtp.rambler.ru"

    msg := BuildMessage(request)
    auth := smtp.PlainAuth("", senderUser, password, host)
    err := smtp.SendMail(addr, auth, senderUser, to, []byte(msg))
	if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Email sent successfully")
}

func BuildMessage(mail Mail) string {
    msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
    msg += fmt.Sprintf("From: %s\r\n", mail.Sender)
    msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
    msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
    msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

    return msg
}