package main

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/stretchr/testify/assert"
)

func TestSystemChecks(t *testing.T) {
	t.Run("CheckCPU", func(t *testing.T) {
		testCPU(t)
	})
	t.Run("CheckMemory", func(t *testing.T) {
		testMemory(t)
	})
	t.Run("CheckOS", func(t *testing.T) {
		testOS(t)
	})
	t.Run("CheckNetwork", func(t *testing.T) {
		testNetwork(t)
	})
}

// Check the number of CPU cores
func testCPU(t *testing.T) {
	logger.Log(t, "Checking CPU cores...")
	command := "cat /proc/cpuinfo | grep 'processor' | wc -l"
	output, err := runCommand(command)
	assert.NoError(t, err, "Failed to get CPU information")
	cpuCores := strings.TrimSpace(output)
	logger.Logf(t, "CPU Cores: %s", cpuCores)
	assert.NotEmpty(t, cpuCores, "CPU cores count should not be empty")
}

// Check total memory (RAM) in MB
func testMemory(t *testing.T) {
	logger.Log(t, "Checking total memory...")
	command := "free -m | grep Mem | awk '{print $2}'"
	output, err := runCommand(command)
	assert.NoError(t, err, "Failed to get memory information")
	totalMemory := strings.TrimSpace(output)
	logger.Logf(t, "Total Memory (MB): %s", totalMemory)
	assert.NotEmpty(t, totalMemory, "Memory size should not be empty")
}

// Check operating system details
func testOS(t *testing.T) {
	logger.Log(t, "Checking operating system details...")
	command := "uname -a"
	output, err := runCommand(command)
	assert.NoError(t, err, "Failed to get OS information")
	osDetails := strings.TrimSpace(output)
	logger.Logf(t, "Operating System: %s", osDetails)
	assert.NotEmpty(t, osDetails, "OS details should not be empty")
}

// Check network information (IP address)
func testNetwork(t *testing.T) {
	logger.Log(t, "Checking network information...")
	command := "hostname -I | awk '{print $1}'"
	output, err := runCommand(command)
	assert.NoError(t, err, "Failed to get network information")
	ipAddress := strings.TrimSpace(output)
	logger.Logf(t, "IP Address: %s", ipAddress)
	assert.NotEmpty(t, ipAddress, "IP Address should not be empty")
}

// Helper function to run shell commands
func runCommand(command string) (string, error) {
	output, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

