import shutil
from pathlib import Path

PROJECT_ROOT = Path.cwd()

folders = [
    "cmd/gostick",
    "internal/controller",
    "internal/input",
    "internal/config",
    "internal/runtime",
    "internal/state",
    "profiles",
    "ui",
]

files = {
    "cmd/gostick/main.go": '''package main

import "gostick/internal/runtime"

func main() {
	runtime.Start()
}
''',

    "internal/runtime/loop.go": '''package runtime

import "fmt"

func Start() {

	fmt.Println("GoStick runtime started")
}
''',

    "internal/state/state.go": '''package state

type AppState struct {
	MouseX int32
	MouseY int32

	ScrollX int32
	ScrollY int32

	DpadX int32
	DpadY int32

	Dragging bool
	PrecisionMode bool
}
''',

    "internal/config/defaults.go": '''package config

const (
	DefaultDeadzone = 2500
	DefaultSensitivity = 15.0
	DefaultPollRateMs = 8
	DefaultPrecisionSensitivity = 6.0
	DefaultScrollThreshold = 4000
)
''',

    "internal/config/config.go": '''package config

type Config struct {
	Deadzone int
	Sensitivity float64
	PrecisionSensitivity float64
	ScrollThreshold int
	PollRateMs int
}
''',

    "internal/controller/device.go": '''package controller

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
''',

    "internal/controller/events.go": '''package controller

type Event struct {
	Type string
	Value int32
}
''',

    "internal/controller/mappings.go": '''package controller

const (
	RightStickX = 3
	RightStickY = 4

	LeftStickX = 0
	LeftStickY = 1

	DpadX = 16
	DpadY = 17

	LT = 2
	RT = 5

	RB = 311
	LB = 310
)
''',

    "internal/input/mouse.go": '''package input

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
''',

    "internal/input/scroll.go": '''package input

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
''',

    "internal/input/keyboard.go": '''package input

import "github.com/bendahl/uinput"

func HandleDpadKeyboard(
	keyboard uinput.Keyboard,
	x int32,
	y int32,
) {

	keyboard.KeyUp(uinput.KeyLeft)
	keyboard.KeyUp(uinput.KeyRight)
	keyboard.KeyUp(uinput.KeyUp)
	keyboard.KeyUp(uinput.KeyDown)

	if x == -1 {
		keyboard.KeyDown(uinput.KeyLeft)
	}

	if x == 1 {
		keyboard.KeyDown(uinput.KeyRight)
	}

	if y == -1 {
		keyboard.KeyDown(uinput.KeyUp)
	}

	if y == 1 {
		keyboard.KeyDown(uinput.KeyDown)
	}
}
''',

    "profiles/default.json": '''{
  "deadzone": 2500,
  "sensitivity": 15,
  "precisionSensitivity": 6,
  "scrollThreshold": 4000,
  "pollRateMs": 8
}
''',

    "README.md": '''# GoStick

Native Linux desktop control with a game controller.

## Features

- Analog cursor movement
- Native wheel scrolling
- Drag support
- D-pad keyboard navigation
- Precision mode
- Xbox controller support
- Wayland compatible
- Low latency uinput backend
'''
}


def create_folders():
    print("\\nCreating folders...")

    for folder in folders:
        path = PROJECT_ROOT / folder

        if not path.exists():
            path.mkdir(parents=True, exist_ok=True)
            print(f"[CREATED] {folder}")
        else:
            print(f"[EXISTS ] {folder}")


def write_files():

    print("\\nProcessing files...")

    for relative_path, content in files.items():

        path = PROJECT_ROOT / relative_path

        if path.exists():

            existing = path.read_text()

            if existing == content:
                print(f"[UNCHANGED] {relative_path}")
                continue

            backup_path = path.with_suffix(path.suffix + ".bak")

            shutil.copy(path, backup_path)

            print(f"[BACKUP ] {backup_path.name}")

        else:
            print(f"[CREATED] {relative_path}")

        path.write_text(content)


def cleanup():

    print("\\nCleanup checks...")

    scratch_go_mod = PROJECT_ROOT / "scratch/go.mod"
    scratch_go_sum = PROJECT_ROOT / "scratch/go.sum"

    if scratch_go_mod.exists():
        print("[WARNING] scratch/go.mod exists")
        print("          Consider removing nested Go modules")

    if scratch_go_sum.exists():
        print("[WARNING] scratch/go.sum exists")


def main():

    print("\\n=== GoStick Project Updater ===")

    create_folders()

    write_files()

    cleanup()

    print("\\nDone.")
    print("\\nRun:")
    print("  go run ./cmd/gostick")


if __name__ == "__main__":
    main()