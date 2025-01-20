package main

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/stretchr/testify/assert"
)

// TestCPUMemory runs the CPU and memory checks.
func TestCPUMemory(t *testing.T) {
	t.Run("CheckCPU", func(t *testing.T) {
		testCPU(t)
	})
	t.Run("CheckMemory", func(t *testing.T) {
		testMemory(t)
	})
}

// testCPU checks the number of CPU cores on the machine.
func testCPU(t *testing.T) {
	logger.Log(t, "Checking CPU cores...")
	command := "cat /proc/cpuinfo | grep 'processor' | wc -l"
	output, err := runCommand(command)
	assert.NoError(t, err, "Failed to get CPU information")
	cpuCores := strings.TrimSpace(output)
	logger.Logf(t, "CPU Cores: %s", cpuCores)
	assert.NotEmpty(t, cpuCores, "CPU cores count should not be empty")
}

// testMemory checks the total memory on the machine in MB.
func testMemory(t *testing.T) {
	logger.Log(t, "Checking total memory...")
	command := "free -m | grep Mem | awk '{print $2}'"
	output, err := runCommand(command)
	assert.NoError(t, err, "Failed to get memory information")
	totalMemory := strings.TrimSpace(output)
	logger.Logf(t, "Total Memory (MB): %s", totalMemory)
	assert.NotEmpty(t, totalMemory, "Memory size should not be empty")
}

// runCommand executes a shell command and returns the output.
func runCommand(command string) (string, error) {
	output, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

