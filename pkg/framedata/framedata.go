package framedata

// Frame data
type Frame struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"w,omitempty"`
	Height int `json:"h,omitempty"`
}

// Frame image types
// Metadata attached to gathered frame images, used to identify which group of
// game data they belong to later on.
type FrameType uint8

const (
	UnitDataType FrameType = iota
	TileDataType
	PropertyDataType
	CODataType
	UIDataType
	OtherDataType
)
