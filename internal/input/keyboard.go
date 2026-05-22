package input

import "github.com/bendahl/uinput"

func HandleDpadX(
	keyboard uinput.Keyboard,
	value int32,
) {

	// RELEASE BOTH
	keyboard.KeyUp(uinput.KeyLeft)
	keyboard.KeyUp(uinput.KeyRight)

	// HOLD LEFT
	if value == -1 {
		keyboard.KeyDown(uinput.KeyLeft)
	}

	// HOLD RIGHT
	if value == 1 {
		keyboard.KeyDown(uinput.KeyRight)
	}
}

func HandleDpadY(
	keyboard uinput.Keyboard,
	value int32,
) {

	// RELEASE BOTH
	keyboard.KeyUp(uinput.KeyUp)
	keyboard.KeyUp(uinput.KeyDown)

	// HOLD UP
	if value == -1 {
		keyboard.KeyDown(uinput.KeyUp)
	}

	// HOLD DOWN
	if value == 1 {
		keyboard.KeyDown(uinput.KeyDown)
	}
}

// A -> ENTER
func HandleAButton(
	keyboard uinput.Keyboard,
	value int32,
) {
	if value == 1 {
		keyboard.KeyPress(uinput.KeyEnter)
	}
}

// X -> ESC
func HandleXButton(
	keyboard uinput.Keyboard,
	value int32,
) {
	if value == 1 {
		keyboard.KeyPress(uinput.KeyEsc)
	}
}