package main

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

// Unit types type/enumeration
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

// Unit variations type/enumeration
type UnitVariation uint8

const (
    OS UnitVariation = iota
    BM
    GE
    YC
    BH
)

// Unit animations type/enumeration
type UnitAnimation uint8

const (
    Idle UnitAnimation = iota
    Right
    Up
    Down
    Left
    Done
)
