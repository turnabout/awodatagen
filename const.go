package main

// Unit types type/enumeration
type unitType uint8

const (
    Infantry unitType = iota
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
type unitVariation uint8

const (
    OS unitVariation = iota
    BM
    GE
    YC
    BH
)

// Unit animations type/enumeration
type unitAnimation uint8

const (
    Idle unitAnimation = iota
    Right
    Up
    Down
)
