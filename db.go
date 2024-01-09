package main

import (
	"fmt"

	bolt "go.etcd.io/bbolt"
)

func createDevicesBucket(db *bolt.DB) {
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Devices"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
}

func updateDevice(db *bolt.DB, name string, deviceId string) error {
	err := db.Update(func(txn *bolt.Tx) error {
		b := txn.Bucket([]byte("Devices"))
		err := b.Put([]byte(name), []byte(deviceId))
		return err
	})

	return err
}

func getDevices(db *bolt.DB) []AudioDevice {
	var storedDevices []AudioDevice
	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("Devices"))

		b.ForEach(func(k, v []byte) error {
			storedDevices = append(storedDevices, AudioDevice{Name: string(k), ID: string(v), Type: "Playback"})
			return nil
		})
		return nil
	})

	return storedDevices
}
