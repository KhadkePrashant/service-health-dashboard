package monitor

import (
	"bytes"
	"os/exec"
	"runtime"
	"strings"
)

type ServiceStatus struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Uptime string `json:"uptime,omitempty"`
}

func GetServiceStatus(name string) (ServiceStatus, error) {
	if runtime.GOOS == "windows" {
		return queryWindowsService(name)
	}
	return querySystemdService(name)
}

func querySystemdService(name string) (ServiceStatus, error) {
	cmd := exec.Command("systemctl", "is-active", name)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	status := strings.TrimSpace(out.String())

	return ServiceStatus{
		Name:   name,
		Status: status,
	}, err
}

func queryWindowsService(name string) (ServiceStatus, error) {
	cmd := exec.Command("sc", "query", name)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	output := out.String()
	status := "unknown"
	if strings.Contains(output, "RUNNING") {
		status = "active"
	} else if strings.Contains(output, "STOPPED") {
		status = "inactive"
	}

	return ServiceStatus{
		Name:   name,
		Status: status,
	}, err
}
