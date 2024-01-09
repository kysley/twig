package main

import (
	"flag"
	"log"
	"os"

	"github.com/charmbracelet/huh"
	bolt "go.etcd.io/bbolt"
)

func createSelectOptions(devices []AudioDevice) []huh.Option[string] {
	var options []huh.Option[string]
	for _, device := range devices {
		options = append(options, huh.NewOption(device.Name, device.ID))
	}

	return options
}

func main() {
	argsWithoutProg := os.Args[1:]

	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("twig.db", 0600, nil)

	if err != nil {
		log.Fatal(err)
	}

	// Ensures that the Devices bucket exists
	createDevicesBucket(db)

	defer db.Close()

	if (argsWithoutProg == nil) || (len(argsWithoutProg) == 0) {
		runListForm(db)
	} else {
		commandOrDevice := argsWithoutProg[0]

		flag.Parse()

		if commandOrDevice == "list" {
			runListForm(db)
		} else if commandOrDevice == "add" {
			runAddForm(db)
		} else {
			db.View(func(tx *bolt.Tx) error {
				// Assume bucket exists and has keys
				b := tx.Bucket([]byte("Devices"))
				deviceId := b.Get([]byte(commandOrDevice))

				if deviceId == nil {
					log.Fatalf("Device '%s' not found", commandOrDevice)
				} else {
					switchToAudioDevice(string(deviceId))
				}
				return nil
			})

			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
