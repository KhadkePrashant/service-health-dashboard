package alert

import (
	"fmt"
	"net/smtp"
	"service-health-dashboard/monitor"
)

type EmailConfig struct {
	SMTPHost string
	SMTPPort string
	SMTPUser string
	SMTPPass string
	EmailTo  string
}

func SendEmailAlert(cfg EmailConfig, status monitor.ServiceStatus) error {
	auth := smtp.PlainAuth("", cfg.SMTPUser, cfg.SMTPPass, cfg.SMTPHost)
	msg := []byte("Subject: Service Down Alert\r\n" +
		"\r\n" +
		fmt.Sprintf("Service %s is down (status: %s)", status.Name, status.Status))
	addr := fmt.Sprintf("%s:%s", cfg.SMTPHost, cfg.SMTPPort)
	return smtp.SendMail(addr, auth, cfg.SMTPUser, []string{cfg.EmailTo}, msg)
}
