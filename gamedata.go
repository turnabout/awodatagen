package awossgen

// Game data JSON structure
type GameData struct {
    Units      UnitsData      `json:"units"`
    Tiles      TilesData      `json:"tiles"`
    Properties PropertiesData `json:"properties"`
    UI         UIData         `json:"ui"`
    Palettes   PaletteData    `json:"palettes"`
    Stages     StageData      `json:"stages"`

    AnimationClocks       []AnimationClock `json:"animationClocks"`
    SpriteSheetDimensions ssDimensions     `json:"ssDimensions"`
}

// Animation Clock structure
type AnimationClock struct {
    ChangingTicks []int   `json:"changingTicks"` // Ticks which update the animations subscribed to this clock
    SubClocks     [][]int `json:"subClocks"`     // Which indexes animations subscribed to this clock should use
}

// Sprite sheet dimensions
type ssDimensions struct {
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
