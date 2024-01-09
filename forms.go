package main

import (
	"errors"
	"log"

	"github.com/charmbracelet/huh"
	bolt "go.etcd.io/bbolt"
)

func runListForm(db *bolt.DB) {
	var (
		deviceId  string
		confirmed bool
	)

	storedDevices := getDevices(db)
	options := createSelectOptions(storedDevices)

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title("Select playback device").Options(options...).Value(&deviceId),
			huh.NewConfirm().Title("Switch?").Value(&confirmed),
		))

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	if confirmed {
		switchToAudioDevice(deviceId)
	} else {
		print("Could not switch to device..")
	}
}

func runAddForm(db *bolt.DB) {
	var (
		deviceId  string
		name      string
		confirmed bool
	)

	devices := parseAudioDeviceList()
	playbackDevices := filterPlaybackDevices(devices)
	options := createSelectOptions(playbackDevices)

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title("Select a playback device").Options(options...).Value(&deviceId),
			huh.NewInput().Title("What do you call this device?").Description("Used to switch output to this device. '>twig headphones'").Value(&name).Validate(func(s string) error {
				if s == "" {
					return errors.New("Name cannot be empty")
				}
				return nil
			}),
			huh.NewConfirm().Title("Add device?").Value(&confirmed),
		))

	err := form.Run()

	if err != nil {
		log.Fatal(err)
	}

	if confirmed {
		updateDevice(db, name, deviceId)
		print("Device added successfully")
	} else {
		print("Device not added")
	}
}
