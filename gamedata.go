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

    AnimationClocks       []AnimationClock `json:"animationClocks"`
    SpriteSheetDimensions SSDimensions     `json:"SSDimensions"`
}

// Animation Clock structure
type AnimationClock struct {
    ChangingTicks []int   `json:"changingTicks"` // Ticks which update the animations subscribed to this clock
    SubClocks     [][]int `json:"subClocks"`     // Which indexes animations subscribed to this clock should use
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
