package awossgen

// Data for all tiles, attached to game data
type TilesData []TileData

// Data for a single tile type
// TODO rename -> TileTypeData
type TileData struct {
    Variations map[string][]Frame `json:"vars"`
    ClockData  *TileClockData     `json:"clockData,omitempty"`
    AutoVars   []AutoVarData      `json:"autoVars"`
}

// Data for a tile's clock
type TileClockData struct {
    Clock           int            `json:"clock"`           // Which sub clock to subscribe to
    DefaultSubClock int            `json:"defaultSubClock"` // Default sub clocks used by this tile's variations
    VarSubClocks    map[string]int `json:"varSubClocks"`    // Sub clocks used by this tile's variations
}


// Tile Auto-vars data
type AutoVarData struct {
    TileVar string       `json:"tileVar"`   // The tile variation's short key
    AdjacentTiles [4]uint `json:"adjTiles"` // Numbers describing the adjacent tiles that correspond to the tile
    // variation. Every number is a bit field where the nth bit corresponds to
    // the nth tile type. If bit n is set, tile type n is acceptable in the
    // adjacent tile.
}

// TODO move in autovar pkg
type RawAutoVarsData map[string][]RawAutoVarData

// TODO move in autovar pkg
type RawAutoVarData struct {
    TileVar string                         // The tile variation
    AdjacentTiles [4]string `json:"Tiles"` // Strings describing the adjacent tiles that correspond to the tile
    // variation. In order, the adjacent tiles go: up, right, down, left
}

// Auto var data's adjacent tile indexes
// TODO move to tilesautovar pkg
const AUTOVAR_ADJACENT_TILE_UP    = 0
const AUTOVAR_ADJACENT_TILE_RIGHT = 1
const AUTOVAR_ADJACENT_TILE_DOWN  = 2
const AUTOVAR_ADJACENT_TILE_LEFT  = 3
const ADJACENT_TILE_COUNT         = 4

// Tile type
type TileType uint8

const(
    // Basic neutral tiles, represented visually
    Plain TileType = iota
    Forest
    Mountain
    Road
    Bridge
    River
    Sea
    Reef
    Shore
    Pipe
    PipeFragile
    Silo

    // Additional neutral tiles, represented visually (standard size)
    BaseSmoke
    Empty

    // Additional neutral tiles, represented visually (non-standard size)
    LandPiece

    // Property tiles, represented visually in their own module (properties module)
    Property_HQ
    Property_City
    Property_Base
    Property_Airport
    Property_Port

    // Meta tiles, not represented visually
    OOB
)

const FirstNeutralTileType = Plain
const LastNeutralTileType  = LandPiece
const NeutralTileTypeCount = LastNeutralTileType + 1

func (t TileType) String() string {
    return [...]string{
        "Plain",
        "Forest",
        "Mountain",
        "Road",
        "Bridge",
        "River",
        "Sea",
        "Reef",
        "Shore",
        "Pipe",
        "PipeFragile",
        "Silo",
        "BaseSmoke",
        "Empty",
        "LandPiece",
        "Property_HQ",
        "Property_City",
        "Property_Base",
        "Property_Airport",
        "Property_Port",
    }[t]
}

// Tile Variations
type TileVariation uint8

const(
    Default TileVariation = iota

    Horizontal
    Vertical
    VerticalEnd

    Top
    Bottom
    DirLeft
    DirRight
    TopLeft
    TopRight
    BottomLeft
    BottomRight
    Middle

    // Shadowed variations
    ShadowedDefault
    ShadowedTopLeft
    ShadowedBottomLeft
    ShadowedLeft
    ShadowedHorizontal
    ShadowedVertical
    ShadowedVerticalEnd
    ShadowedTLeft

    // T-shape variations (river/road)
    TTop
    TBottom
    TLeft
    TRight

    // Mountain only
    Small

    // River only
    WaterfallUp
    WaterfallDown
    WaterfallLeft
    WaterfallRight

    // Sea/shore only
    Hole
    HoleHorizontal
    HoleVertical
    HoleLeft
    HoleRight
    HoleTop
    HoleBottom

    // Shore only
    TopConnectedLeft
    TopConnectedRight
    TopConnectedFull

    BottomConnectedLeft
    BottomConnectedRight
    BottomConnectedFull

    LeftConnectedTop
    LeftConnectedBottom
    LeftConnectedFull

    RightConnectedTop
    RightConnectedBottom
    RightConnectedFull

    TopLeftConnectedVertical
    TopLeftConnectedHorizontal
    TopLeftConnectedFull

    TopRightConnectedVertical
    TopRightConnectedHorizontal
    TopRightConnectedFull

    BottomLeftConnectedVertical
    BottomLeftConnectedHorizontal
    BottomLeftConnectedFull

    BottomRightConnectedVertical
    BottomRightConnectedHorizontal
    BottomRightConnectedFull

    // Fragile pipe only
    HorizontalClosed
    HorizontalOpen
    VerticalClosed
    VerticalOpen

    // Silo only
    Used
)

func (v TileVariation) String() string {
    return [...]string{
        "A",
        "B",
        "C",
        "D",
        "E",
        "F",
        "G",
        "H",
        "I",
        "J",
        "K",
        "L",
        "M",
        "N",
        "O",
        "P",
        "Q",
        "R",
        "S",
        "T",
        "U",
        "V",
        "W",
        "X",
        "Y",
        "Z",
        "a",
        "b",
        "c",
        "d",
        "e",
        "f",
        "g",
        "h",
        "i",
        "j",
        "k",
        "l",
        "m",
        "n",
        "o",
        "p",
        "q",
        "r",
        "s",
        "t",
        "u",
        "v",
        "w",
        "x",
        "y",
        "z",
        "0",
        "1",
        "2",
        "3",
        "4",
        "5",
        "6",
        "7",
        "8",
        "9",
        ":",
        ";",
        "-",
        "=",
    }[v]
}

