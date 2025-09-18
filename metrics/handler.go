package metrics

import (
	"encoding/json"
	"net/http"
	"service-health-dashboard/monitor"
)

func MetricsHandler(statuses *map[string]monitor.ServiceStatus) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(statuses)
	}
}
