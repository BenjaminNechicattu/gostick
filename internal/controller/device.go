package controller

import (
	"log"
	"path/filepath"
	"strings"

	evdev "github.com/holoplot/go-evdev"
)

func FindController() *evdev.InputDevice {

	eventFiles, err := filepath.Glob("/dev/input/event*")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range eventFiles {

		device, err := evdev.Open(file)
		if err != nil {
			continue
		}

		name, err := device.Name()
		if err != nil {
			continue
		}

		lower := strings.ToLower(name)

		if strings.Contains(lower, "xbox") ||
			strings.Contains(lower, "x-box") {

			return device
		}
	}

	return nil
}
