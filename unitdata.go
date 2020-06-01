package awodatagen

// Data for all units, attached to game data
type UnitData [UnitTypeCount]UnitTypeData

// Data for a single unit type
type UnitTypeData struct {

    // Variation -> Animations -> Animation
    Variations [][][]Frame `json:"vars"` // TODO: Rename to "frames"
}

// Unit type enum
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

const UnitTypeFirst = Infantry
const UnitTypeLast  = TransportCopter
const UnitTypeCount = UnitTypeLast + 1

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

// Unit animation enum
type UnitAnimation uint8

const (
    Idle UnitAnimation = iota
    Right
    Up
    Down
)

const UnitAnimFirst = Idle
const UnitAnimLast  = Down
const UnitAnimCount = UnitAnimLast + 1

func (a UnitAnimation) String() string {
    return [...]string{"Idle", "Right", "Up", "Down"}[a]
}

// Army type enum (unit variation/property variation)
type ArmyType uint8

const (
    OS ArmyType = iota
    BM
    GE
    YC
    BH
)

const ArmyTypeFirst = OS
const ArmyTypeLast  = BH
const ArmyTypeCount = ArmyTypeLast + 1

func (v ArmyType) String() string {
    return [...]string{"OS", "BM", "GE", "YC", "BH"}[v]
}
