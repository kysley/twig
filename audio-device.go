package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type AudioDevice struct {
	ID   string
	Name string
	Type string
}

func parseAudioDeviceList() []AudioDevice {
	var devices []AudioDevice

	output, _ := exec.Command("powershell", "Get-AudioDevice", "-List").CombinedOutput()

	scanner := bufio.NewScanner(bytes.NewReader(output))

	var device AudioDevice

	// Step through each line of the output
	for scanner.Scan() {
		line := scanner.Text()
		// Skip empty lines
		if (line == "") || (line == "\n") {
			continue
		}
		// ID is the last line that we are interested in- append to device list and reset device struct
		// This is a bit hacky, but it works. A better approach would be resetting device if a new line is found
		if strings.HasPrefix(line, "ID") {
			device.ID = strings.TrimSpace(strings.Split(line, ":")[1])
			devices = append(devices, device)
			device = AudioDevice{}
		} else if strings.HasPrefix(line, "Name") {
			device.Name = strings.TrimSpace(strings.Split(line, ":")[1])
		} else if strings.HasPrefix(line, "Type") {
			device.Type = strings.TrimSpace(strings.Split(line, ":")[1])
		}
	}

	return devices
}

func filterPlaybackDevices(devices []AudioDevice) []AudioDevice {
	var playbackDevices []AudioDevice

	for _, device := range devices {
		if device.Type == "Playback" {
			playbackDevices = append(playbackDevices, device)
		}
	}

	return playbackDevices
}

func switchToAudioDevice(deviceId string) (string, error) {
	out, err := exec.Command("powershell", "Set-AudioDevice", "-ID", fmt.Sprintf("'%s'", deviceId)).CombinedOutput()

	if err != nil {
		return string(out), err
	} else {
		return string(out), nil
	}
}
