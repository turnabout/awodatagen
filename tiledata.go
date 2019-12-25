package awodatagen

// Data for all tiles, attached to game data
type TileData []TileTypeData

// Data for a single tile type
type TileTypeData struct {
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

// Tile auto-vars data
type AutoVarData struct {
    TileVar string       `json:"tileVar"`   // The tile variation's short key
    AdjacentTiles [4]uint `json:"adjTiles"` // Numbers describing the adjacent tiles that correspond to the tile
                                            // variation. Every number is a bit field where the nth bit corresponds to
                                            // the nth tile type. If bit n is set, tile type n is acceptable in the
                                            // adjacent tile.
}

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

const NeutralTileTypeFirst = Plain
const NeutralTileTypeLast  = LandPiece
const NeutralTileTypeCount = NeutralTileTypeLast + 1

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
        "A","B","C","D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W",
        "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s",
        "t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ":", ";", "-", "=",
    }[v]
}

// Map for looking up a Tile Type using its corresponding full string
var TileReverseStrings = map[string]TileType {

    // Basic neutral tiles, represented visually
    "Plain": Plain,
    "Forest": Forest,
    "Mountain": Mountain,
    "Road": Road,
    "Bridge": Bridge,
    "River": River,
    "Sea": Sea,
    "Reef": Reef,
    "Shore": Shore,
    "Pipe": Pipe,
    "PipeFragile": PipeFragile,
    "Silo": Silo,

    // Additional neutral tiles, represented visually (standard size)
    "BaseSmoke": BaseSmoke,
    "Empty": Empty,

    // Additional neutral tiles, represented visually (non-standard size)
    "LandPiece": LandPiece,
}

// Map for looking up a Tile Variation using its corresponding full string
var TileVarsReverseStrings = map[string]TileVariation {
    "Default": Default,
    "Horizontal": Horizontal,
    "Vertical": Vertical,
    "VerticalEnd": VerticalEnd,
    "Top": Top,
    "Bottom": Bottom,
    "Left": DirLeft,
    "Right": DirRight,
    "TopLeft": TopLeft,
    "TopRight": TopRight,
    "BottomLeft": BottomLeft,
    "BottomRight": BottomRight,
    "Middle": Middle,
    "ShadowedDefault": ShadowedDefault,
    "ShadowedTopLeft": ShadowedTopLeft,
    "ShadowedBottomLeft": ShadowedBottomLeft,
    "ShadowedLeft": ShadowedLeft,
    "ShadowedHorizontal": ShadowedHorizontal,
    "ShadowedVertical": ShadowedVertical,
    "ShadowedVerticalEnd": ShadowedVerticalEnd,
    "ShadowedTLeft": ShadowedTLeft,
    "TTop": TTop,
    "TBottom": TBottom,
    "TLeft": TLeft,
    "TRight": TRight,
    "Small": Small,
    "WaterfallUp": WaterfallUp,
    "WaterfallDown": WaterfallDown,
    "WaterfallLeft": WaterfallLeft,
    "WaterfallRight": WaterfallRight,
    "Hole": Hole,
    "HoleHorizontal": HoleHorizontal,
    "HoleVertical": HoleVertical,
    "HoleLeft": HoleLeft,
    "HoleRight": HoleRight,
    "HoleTop": HoleTop,
    "HoleBottom": HoleBottom,
    "TopConnectedLeft": TopConnectedLeft,
    "TopConnectedRight": TopConnectedRight,
    "TopConnectedFull": TopConnectedFull,
    "BottomConnectedLeft": BottomConnectedLeft,
    "BottomConnectedRight": BottomConnectedRight,
    "BottomConnectedFull": BottomConnectedFull,
    "LeftConnectedTop": LeftConnectedTop,
    "LeftConnectedBottom": LeftConnectedBottom,
    "LeftConnectedFull": LeftConnectedFull,
    "RightConnectedTop": RightConnectedTop,
    "RightConnectedBottom": RightConnectedBottom,
    "RightConnectedFull": RightConnectedFull,
    "TopLeftConnectedVertical": TopLeftConnectedVertical,
    "TopLeftConnectedHorizontal": TopLeftConnectedHorizontal,
    "TopLeftConnectedFull": TopLeftConnectedFull,
    "TopRightConnectedVertical": TopRightConnectedVertical,
    "TopRightConnectedHorizontal": TopRightConnectedHorizontal,
    "TopRightConnectedFull": TopRightConnectedFull,
    "BottomLeftConnectedVertical": BottomLeftConnectedVertical,
    "BottomLeftConnectedHorizontal": BottomLeftConnectedHorizontal,
    "BottomLeftConnectedFull": BottomLeftConnectedFull,
    "BottomRightConnectedVertical": BottomRightConnectedVertical,
    "BottomRightConnectedHorizontal": BottomRightConnectedHorizontal,
    "BottomRightConnectedFull": BottomRightConnectedFull,
    "HorizontalClosed": HorizontalClosed,
    "HorizontalOpen": HorizontalOpen,
    "VerticalClosed": VerticalClosed,
    "VerticalOpen": VerticalOpen,
    "Used": Used,
}
