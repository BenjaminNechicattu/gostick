package input

import (
	"time"

	"github.com/bendahl/uinput"
)

func HandleScroll(
	mouse uinput.Mouse,
	x int32,
	y int32,
	threshold int,
) {

	if y > int32(threshold) {

		mouse.Wheel(false, -1)
		time.Sleep(20 * time.Millisecond)

	} else if y < -int32(threshold) {

		mouse.Wheel(false, 1)
		time.Sleep(20 * time.Millisecond)
	}

	if x > int32(threshold) {

		mouse.Wheel(true, 1)
		time.Sleep(20 * time.Millisecond)

	} else if x < -int32(threshold) {

		mouse.Wheel(true, -1)
		time.Sleep(20 * time.Millisecond)
	}
}
