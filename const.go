package main

import "image"

// Visual data JSON structure
type VisualData struct {
    Units UnitsData `json:"units"`
}

type UnitsData struct {
    Origin [][][][]Frame `json:"origin"`
    Dest   [][][]Frame   `json:"dest"`
}

// Used to store a frame's visual data within the game's sprite sheet
type Frame struct {
    X int      `json:"x"`
    Y int      `json:"y"`
    Width int  `json:"w"`
    Height int `json:"h"`
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
    UnitTypeAmount
)

const FirstUnitType = Infantry

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
    UnitVariationAmount
)

const FirstUnitVariation = OS

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
    UnitAnimationAmount // "Left" and "Done" don't count as base animations as they're generated in-game
    Left
    Done
)

const FirstUnitAnimation = Idle

func (a UnitAnimation) String() string {
    return [...]string{"Idle", "Right", "Up", "Down", "Left", "Done"}[a]
}

// A frame's image data (image/width/Height)
type FrameImage struct {
    Image  image.Image
    Width  int
    Height int
}

// Data detailing a row of sprite images in a sprite sheet
type RowData struct {
    Height int // Height in pixels
    Amount int // Amount of images in the row
    Y      int // Row's Y coordinate
}
