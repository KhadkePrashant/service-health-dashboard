package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"service-health-dashboard/alert"
	"service-health-dashboard/config"
	"service-health-dashboard/metrics"
	"service-health-dashboard/monitor"
	"syscall"
	"time"
)

func main() {
	cfg := config.Config{
		Services:     []string{"ssh", "cron"}, // change based on OS
		PollInterval: 10,
		SMTPHost:     "smtp.gmail.com",
		SMTPPort:     "587",
		SMTPUser:     "wildlifewondersunveiled@gmail.com",
		SMTPPass:     "ialjcgryxzvcecgs",
		EmailTo:      "khushisahu4326@gmail.com",
	}

	emailCfg := alert.EmailConfig{
		SMTPHost: cfg.SMTPHost,
		SMTPPort: cfg.SMTPPort,
		SMTPUser: cfg.SMTPUser,
		SMTPPass: cfg.SMTPPass,
		EmailTo:  cfg.EmailTo,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	m := &monitor.Monitor{
		Services: cfg.Services,
		Alerts: func(status monitor.ServiceStatus) {
			log.Printf("Service alert: %v", status)
			if err := alert.SendEmailAlert(emailCfg, status); err != nil {
				log.Printf("Failed to send alert: %v", err)
			}
		},
	}

	go m.Start(ctx, time.Duration(cfg.PollInterval)*time.Second)

	http.HandleFunc("/metrics", metrics.MetricsHandler(&m.Statuses))
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
