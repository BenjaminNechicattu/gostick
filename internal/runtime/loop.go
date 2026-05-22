package runtime

import (
	"fmt"
	"log"
	"time"

	"gostick/internal/config"
	"gostick/internal/controller"
	"gostick/internal/input"
	"gostick/internal/state"

	"github.com/bendahl/uinput"
	evdev "github.com/holoplot/go-evdev"
)

var (
	mouse    uinput.Mouse
	keyboard uinput.Keyboard
	appState state.AppState
)

func Start() {

	controllerDevice := controller.FindController()

	if controllerDevice == nil {
		fmt.Println("Controller not found")
		return
	}

	var err error

	mouse, err = uinput.CreateMouse(
		"/dev/uinput",
		[]byte("GoStick Mouse"),
	)

	if err != nil {
		log.Fatal(err)
	}

	defer mouse.Close()

	keyboard, err = uinput.CreateKeyboard(
		"/dev/uinput",
		[]byte("GoStick Keyboard"),
	)

	if err != nil {
		log.Fatal(err)
	}

	defer keyboard.Close()

	name, _ := controllerDevice.Name()

	fmt.Println("Connected:", name)

	go eventLoop(controllerDevice)

	ticker := time.NewTicker(
		time.Duration(config.DefaultPollRateMs) *
			time.Millisecond,
	)

	defer ticker.Stop()

	for range ticker.C {

		input.MoveMouse(
			mouse,
			appState.MouseX,
			appState.MouseY,
			config.DefaultDeadzone,
			getSensitivity(),
		)

		input.HandleScroll(
			mouse,
			appState.ScrollX,
			appState.ScrollY,
			config.DefaultScrollThreshold,
		)
	}
}

func getSensitivity() float64 {

	if appState.PrecisionMode {
		return config.DefaultPrecisionSensitivity
	}

	return config.DefaultSensitivity
}

func eventLoop(
	controllerDevice *evdev.InputDevice,
) {

	for {

		ev, err := controllerDevice.ReadOne()
		if err != nil {
			log.Fatal(err)
		}

		switch {

		// RIGHT STICK X
		case ev.Type == 3 &&
			ev.Code == controller.RightStickX:

			appState.MouseX = ev.Value

		// RIGHT STICK Y
		case ev.Type == 3 &&
			ev.Code == controller.RightStickY:

			appState.MouseY = ev.Value

		// LEFT STICK X
		case ev.Type == 3 &&
			ev.Code == controller.LeftStickX:

			appState.ScrollX = ev.Value

		// LEFT STICK Y
		case ev.Type == 3 &&
			ev.Code == controller.LeftStickY:

			appState.ScrollY = ev.Value

		// DPAD X
		case ev.Type == 3 &&
			ev.Code == controller.DpadX:

			appState.DpadX = ev.Value

			input.HandleDpadX(
				keyboard,
				ev.Value,
			)

		// DPAD Y
		case ev.Type == 3 &&
			ev.Code == controller.DpadY:

			appState.DpadY = ev.Value

			input.HandleDpadY(
				keyboard,
				ev.Value,
			)

		// LT -> PRECISION MODE
		case ev.Type == 3 &&
			ev.Code == controller.LT:

			appState.PrecisionMode =
				ev.Value > 120

		// RB -> LEFT CLICK
		case ev.Type == 1 &&
			ev.Code == controller.RB &&
			ev.Value == 1:

			mouse.LeftClick()

		// LB -> RIGHT CLICK
		case ev.Type == 1 &&
			ev.Code == controller.LB &&
			ev.Value == 1:

			mouse.RightClick()

		// RT -> DRAG
		case ev.Type == 3 &&
			ev.Code == controller.RT:

			handleDrag(ev.Value)
		}
	}
}

func handleDrag(v int32) {

	const triggerThreshold = 120

	pressed := v > triggerThreshold

	// START HOLD
	if pressed && !appState.Dragging {

		appState.Dragging = true

		mouse.LeftPress()

		return
	}

	// RELEASE HOLD
	if !pressed && appState.Dragging {

		appState.Dragging = false

		mouse.LeftRelease()
	}
}
