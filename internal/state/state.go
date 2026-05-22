package state

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
