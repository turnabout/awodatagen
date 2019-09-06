package main

// Visual data JSON structure
type VisualData struct {
    Units *UnitsData `json:"units"`
    Tiles *TilesData `json:"tiles"`
    Properties *PropertiesData `json:"properties"`
    SSMetaData ssMetaData `json:"ssMetaData"`

    AnimationSubClocks int `json:"animationSubClocks"` // TODO
    Stages int `json:"stages"` // TODO
}

type UnitsData struct {
    Src [][][][]Frame `json:"src"`
    Dst [][][]Frame   `json:"dst"`

    SrcX      int `json:"srcX"`
    SrcY      int `json:"srcY"`
    SrcWidth  int `json:"srcWidth"`
    SrcHeight int `json:"srcHeight"`
    DstWidth  int `json:"dstWidth"`
    DstHeight int `json:"dstHeight"`

    BasePalette Palette       `json:"basePalette"`
    Palettes    []UnitPalette `json:"palettes"`
    BaseDoneOps []CanvasOp    `json:"baseDoneOps"` // Operations used to generate "Done" animation frames

    frameImg FrameImage
}

type TilesData struct {
    Src       []TileData `json:"src"`
    ClockData int        `json:"cData"` // TODO

    SrcX      int `json:"srcX"`
    SrcY      int `json:"srcY"`
    SrcWidth  int `json:"srcWidth"`
    SrcHeight int `json:"srcHeight"`

    BasePalette Palette   `json:"basePalette"`
    Palettes    []Palette `json:"palettes"`
    FogOps      CanvasOp  `json:"fogOps"` // Operations used to apply a "fog" effect to Tiles

    frameImg FrameImage
}

type PropertiesData struct {
    Src    [][][]Frame `json:"src"`
    Dst    [][]Frame   `json:"dst"`
    FogDst [][]Frame   `json:"fogDst"`

    SrcX         int `json:"srcX"`
    SrcY         int `json:"srcY"`
    SrcWidth     int `json:"srcWidth"`
    SrcHeight    int `json:"srcHeight"`
    DstWidth     int `json:"dstWidth"`
    DstHeight    int `json:"dstHeight"`
    FogDstWidth  int `json:"fogDstWidth"`
    FogDstHeight int `json:"fogDstHeight"`

    Palettes       []Palette `json:"palettes"`
    PropsLightsRGB RGB       `json:"propLightsRGB"` // RGB used for Properties' lights

    frameImg FrameImage
}

type TileData struct {
    Variations map[string][]Frame `json:"vars"`
    ClockData int `json:"cData"` // TODO
}

type ssMetaData struct {
    Width int `json:"width"`
    Height int `json:"height"`
}

// Used to store a frame's visual data within the game's sprite sheet
type Frame struct {
    X int      `json:"x"`
    Y int      `json:"y"`
    Width int  `json:"w,omitempty"`
    Height int `json:"h,omitempty"`
}

// Unit Types
type UnitType uint8

const (
    Infantry UnitType = iota
    Mech
    Recon
    Tank
    MdTank
    NeoTank
    APC
    Artillery
    Rockets
    Missiles
    AntiAir
    Battleship
    Cruiser
    Lander
    Sub
    Fighter
    Bomber
    BattleCopter
    TransportCopter
)

const FirstUnitType = Infantry
const LastUnitType = TransportCopter
const UnitTypeAmount = LastUnitType + 1

func (u UnitType) String() string {
    return [...]string{
        "Infantry",
        "Mech",
        "Recon",
        "Tank",
        "MdTank",
        "NeoTank",
        "APC",
        "Artillery",
        "Rockets",
        "Missiles",
        "AntiAir",
        "Battleship",
        "Cruiser",
        "Lander",
        "Sub",
        "Fighter",
        "Bomber",
        "BattleCopter",
        "TransportCopter",
    }[u]
}

// Unit Variations
type UnitVariation uint8

const (
    OS UnitVariation = iota
    BM
    GE
    YC
    BH
)

const FirstUnitVariation = OS
const LastUnitVariation = BH
const UnitVariationAmount = LastUnitVariation + 1

func (v UnitVariation) String() string {
    return [...]string{"OS", "BM", "GE", "YC", "BH"}[v]
}

// Unit Animations
type UnitAnimation uint8

const (
    Idle UnitAnimation = iota
    Right
    Up
    Down
    Done
    Left
)

const FirstUnitAnimation = Idle
const LastUnitAnimation = Down // "Left" and "Done" don't count as base animations as they're generated in-game
const UnitAnimationAmount = LastUnitAnimation + 1
const UnitAnimationFullAmount = Left + 1 // Full amount of unit animations, including in-game generated ones
const UnitExtraAnimationConvert = Done // Amount to add to an Animation index to get its corresponding extra Animation

func (a UnitAnimation) String() string {
    return [...]string{"Idle", "Right", "Up", "Down", "Left", "Done"}[a]
}

// Tile Types
type TileType uint8

const(
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
    BaseSmoke
    LandPiece
    // HQ
    // City
    // Base
    // Airport
    // Port
)

const FirstBasicTileType = Plain
const LastBasicTileType = LandPiece
const BasicTileAmount = LastBasicTileType + 1

// const FirstPropertyTileType = HQ
// const LastPropertyTileType = Port
// const PropertyTileAmount = (LastPropertyTileType + 1) - BasicTileAmount

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
        "LandPiece",
        "HQ",
        "City",
        "Base",
        "Airport",
        "Port",
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
        "SrcX",
        "SrcY",
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
        "<",
        "=",
    }[v]
}

// Visual Data IDs for keeping an order on generated sprite sheets
type VisualDataID uint8

const(
    VisualDataUnits VisualDataID = iota
    VisualDataTiles
    VisualDataProperties
)

const FirstVisualDataID = VisualDataUnits
const LastVisualDataID = VisualDataTiles

// Property Types
type PropertyType uint8

const(
    HQ PropertyType = iota
    City
    Base
    Airport
    Port
)

func (p PropertyType) String() string {
    return [...]string{
        "HQ",
        "City",
        "Base",
        "Airport",
        "Port",
    }[p]
}

const FirstPropertyType = HQ
const LastPropertyType = Port
const PropertyTypeAmount = LastPropertyType + 1

// Property Weather Variations
type PropertyWeatherVariation uint8

const(
    Clear PropertyWeatherVariation = iota
    Snow
)

func (p PropertyWeatherVariation) String() string {
    return [...]string{
        "Clear",
        "Snow",
    }[p]
}

const FirstPropertyWeatherVariation = Clear
const LastPropertyWeatherVariation = Snow
const PropertyWeatherVariationAmount = Snow + 1

// Array with two strings representing a graphical operation on the game's canvas
type CanvasOp [2]string

// Array representing an RGB pixel value
type RGB [3]uint8

// Generic palette
type Palette map[string]RGB

// Unit palette structure
type UnitPalette struct {
    Flip bool `json:"flip"`
    DoneOps []CanvasOp `json:"doneOps"`
    Palette Palette `json:"palette"`
}
