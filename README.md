# GoStick

Native Linux desktop control using a game controller.

GoStick converts a controller into a low-latency desktop input system with:
- analog mouse movement
- native scrolling
- keyboard navigation
- drag support
- precision mode
- Wayland-compatible uinput injection

Built in Go for Linux.

---

# Features

## Mouse Control
- Right analog stick controls cursor
- Radial acceleration curve
- Subpixel precision accumulation
- Precision mode using LT

## Mouse Buttons
- RB → Left Click
- LB → Right Click
- RT → Click + Hold Drag

## Scrolling
- Left stick vertical → Vertical scroll
- Left stick horizontal → Horizontal scroll
- Native uinput wheel events

## Keyboard Navigation
- D-pad simulates keyboard arrow keys
- Tap → single movement
- Hold → natural OS key repeat

## Linux Native
- Uses `/dev/uinput`
- Works with Wayland
- No X11 dependency
- No external automation tools required

---

# Architecture

GoStick uses a modular architecture designed for:
- GUI support
- live configuration
- rebinding
- profiles
- multiple controller types
- future plugin systems

---

# Project Structure

```text
gostick/
├── cmd/
│   └── gostick/
│       └── main.go
│
├── internal/
│   ├── config/
│   │   ├── config.go
│   │   └── defaults.go
│   │
│   ├── controller/
│   │   ├── device.go
│   │   ├── events.go
│   │   └── mappings.go
│   │
│   ├── input/
│   │   ├── keyboard.go
│   │   ├── mouse.go
│   │   └── scroll.go
│   │
│   ├── runtime/
│   │   └── loop.go
│   │
│   └── state/
│       └── state.go
│
├── profiles/
│   └── default.json
│
├── ui/
│
├── go.mod
└── README.md