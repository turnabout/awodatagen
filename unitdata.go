package awodatagen

// Data for all units, attached to game data
type UnitData [UnitTypeCount]UnitTypeData

// Data for a single unit type
type UnitTypeData struct {
	MovementType    MovementType `json:"movementType"`
	Movement        uint8        `json:"movement"`
	Vision          uint8        `json:"vision"`
    Fuel            uint8        `json:"fuel"`
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
    Treads
    Tires
    Air
    Naval
)

// Map for looking up a Weapon Type using its corresponding string
var MovementTypeReverseStrings = map[string]MovementType {
    "Foot": Foot,
    "HeavyFoot": HeavyFoot,
    "Treads": Treads,
    "Tires": Tires,
    "Air": Air,
    "Naval": Naval,
}

// Weapon type enum
type WeaponType uint8

const (
    WeaponMachineGunMk1 WeaponType = iota // Infantry secondary
    WeaponMachineGunMk2                   // Mech/BattleCopter secondary
    WeaponMachineGunMk3                   // Recon/Tank/MdTank secondary
    WeaponMachineGunMk4                   // NeoTank secondary

    WeaponCannonRangedMk1                 // Artillery primary
    WeaponCannonRangedMk2                 // Battleship primary
    WeaponCannonMk1                       // Tank primary
    WeaponCannonMk2                       // MdTank primary
    WeaponCannonMk3                       // NeoTank primary

    WeaponMissilesRangedMk1               // Missiles primary
    WeaponMissilesMk1                     // BattleCopter primary
    WeaponMissilesMk2                     // Fighter primary
    WeaponMissilesMk3                     // Cruiser primary

    WeaponBombs                           // Bomber primary
    WeaponRockets                         // Rockets primary
    WeaponTorpedoes                       // Sub primary
    WeaponAntiAirGun                      // Cruiser secondary
    WeaponBazooka                         // Mech primary
    WeaponVulcan                          // AntiAir primary
)

// Map for looking up a Weapon Type using its corresponding string
var WeaponTypeReverseStrings = map[string]WeaponType {
    "MachineGunMk1":     WeaponMachineGunMk1,
    "MachineGunMk2":     WeaponMachineGunMk2,
    "MachineGunMk3":     WeaponMachineGunMk3,
    "MachineGunMk4":     WeaponMachineGunMk4,
    "CannonRangedMk1":   WeaponCannonRangedMk1,
    "CannonRangedMk2":   WeaponCannonRangedMk2,
    "CannonMk1":         WeaponCannonMk1,
    "CannonMk2":         WeaponCannonMk2,
    "CannonMk3":         WeaponCannonMk3,
    "MissilesRangedMk1": WeaponMissilesRangedMk1,
    "MissilesMk1":       WeaponMissilesMk1,
    "MissilesMk2":       WeaponMissilesMk2,
    "MissilesMk3":       WeaponMissilesMk3,
    "Bombs":             WeaponBombs,
    "Rockets":           WeaponRockets,
    "Torpedoes":         WeaponTorpedoes,
    "AntiAirGun":        WeaponAntiAirGun,
    "Bazooka":           WeaponBazooka,
    "Vulcan":            WeaponVulcan,
}
