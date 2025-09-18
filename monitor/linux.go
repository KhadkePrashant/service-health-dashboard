package main

import (
	"fmt"
	"os/exec"
)

func main() {

	diskstats, err := CollectIoStats()
	if err != nil {
		fmt.Println("Error Collecting IO sats:", err)
		return
	}
	fmt.Println(diskstats)
}

func CollectIoStats() (string, error) {

	cmd := exec.Command("iostat", "-xz", "1", "1")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running Io stats %v\n", err)
		return err.Error(), err
	}
	return string(output), nil
}
