package input

import (
	"math"

	"github.com/bendahl/uinput"
)

var (
	FractionalX float64
	FractionalY float64
)

func MoveMouse(
	mouse uinput.Mouse,
	x int32,
	y int32,
	deadzone int,
	sensitivity float64,
) {

	fx := float64(x)
	fy := float64(y)

	magnitude := math.Sqrt(fx*fx + fy*fy)

	if magnitude < float64(deadzone) {
		return
	}

	nx := fx / magnitude
	ny := fy / magnitude

	strength := magnitude / 32767.0

	if strength > 1 {
		strength = 1
	}

	curved := math.Pow(strength, 1.8)

	speed := curved * sensitivity

	FractionalX += nx * speed
	FractionalY += ny * speed

	dx := int32(FractionalX)
	dy := int32(FractionalY)

	FractionalX -= float64(dx)
	FractionalY -= float64(dy)

	if dx == 0 && dy == 0 {
		return
	}

	mouse.Move(dx, dy)
}
