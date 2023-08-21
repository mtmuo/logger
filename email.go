package logger

import (
	"gopkg.in/gomail.v2"
)

type Email struct {
	Host      string   `json:"host" yaml:"host"`
	Port      int      `json:"port" yaml:"port"`
	Username  string   `json:"username" yaml:"username"`
	Password  string   `json:"password" yaml:"password"`
	Recipient []string `json:"recipient" yaml:"recipient"`
	dialer    *gomail.Dialer
}

func (m *Email) Dialer() {
	dialer := gomail.NewDialer(
		m.Host,
		m.Port,
		m.Username,
		m.Password,
	)
	m.dialer = dialer
}

func (m *Email) Send(subject, context string) error {
	message := gomail.NewMessage()
	message.SetHeader("From", m.dialer.Username)
	message.SetHeader("To", m.Recipient...)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", context)
	return m.dialer.DialAndSend(message)
}
