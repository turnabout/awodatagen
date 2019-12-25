package awodatagen

// Game data JSON structure
type GameData struct {
    Units      UnitData     `json:"units"`
    Tiles      TileData     `json:"tiles"`
    Properties PropertyData `json:"properties"`
    UI         UIData       `json:"ui"`
    Palettes   PaletteData  `json:"palettes"`
    Stages     StageData    `json:"stages"`
    COs        COData       `json:"COs"`

    Clocks                []AnimationClock `json:"clocks"`
    SpriteSheetDimensions SSDimensions     `json:"ssDimensions"`
}

// Animation Clock structure
type AnimationClock struct {
    Frames []int `json:"frames"` // Frame counts that make the animation clock tick
    Values []int `json:"values"` // Values emitted by the animation clock when it ticks
}

// Sprite sheet dimensions
type SSDimensions struct {
    Width int `json:"width"`
    Height int `json:"height"`
}

// Frame data
type Frame struct {
    X int      `json:"x"`
    Y int      `json:"y"`
    Width int  `json:"w,omitempty"`
    Height int `json:"h,omitempty"`
}

// Sections of the game data
type GameDataType uint8

const(
    UnitDataType GameDataType = iota
    TileDataType
    PropertyDataType
    CODataType
    UIDataType
    OtherDataType
)
