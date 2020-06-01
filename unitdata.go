package awodatagen

// Data for all units, attached to game data
type UnitData [UnitTypeCount]UnitTypeData

// Data for a single unit type
type UnitTypeData struct {
	MovementType    MovementType `json:"movementType"`
	Movement        uint8        `json:"movement"`
	Vision          uint8        `json:"vision"`
    Fuel            uint8        `json:"fuel"`
    Ammo            uint8        `json:"ammo"`
	WeaponPrimary   WeaponType   `json:"weaponPrimary"`
    WeaponSecondary WeaponType   `json:"weaponSecondary"`

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

// Movement type enum
type MovementType uint8

const (
    Foot MovementType = iota
    HeavyFoot
)

// Map for looking up a Weapon Type using its corresponding string
var MovementTypeReverseStrings = map[string]MovementType {
    "Foot": Foot,
    "HeavyFoot": HeavyFoot,
}

// Weapon type enum
type WeaponType uint8

const (
    MachineGun WeaponType = iota
)

// Map for looking up a Weapon Type using its corresponding string
var WeaponTypeReverseStrings = map[string]WeaponType {
    "MachineGun": MachineGun,
}
