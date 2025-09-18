package monitor

import (
	"context"
	"log"
	"time"
)

type Monitor struct {
	Services []string
	Statuses map[string]ServiceStatus
	Alerts   func(ServiceStatus)
}

func (m *Monitor) Start(ctx context.Context, interval time.Duration) {
	m.Statuses = make(map[string]ServiceStatus)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			for _, service := range m.Services {
				go m.checkWithRetry(service)
			}
		}
	}
}

func (m *Monitor) checkWithRetry(service string) {
	var status ServiceStatus
	var err error

	backoff := time.Second
	retries := 3

	for i := 0; i < retries; i++ {
		status, err = GetServiceStatus(service)
		if err == nil {
			break
		}
		log.Printf("Error checking service %s: %v. Retrying in %v...", service, err, backoff)
		time.Sleep(backoff)
		backoff *= 2
	}

	m.Statuses[service] = status
	if status.Status != "active" {
		m.Alerts(status)
	}
}
