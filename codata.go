package awodatagen

// Data for all COs, attached to game data
type COData []COTypeData

// Data for a single CO type
type COTypeData struct {
    Name string    `json:"name"`
    Army ArmyType  `json:"army"`
    Frames []Frame `json:"frames"`
}

// CO enum
type CO uint8

const(
    Andy CO = iota
    Max
    Sami
    Nell
    Hachi
    Olaf
    Grit
    Colin
    Eagle
    Drake
    Jess
    Kanbei
    Sonja
    Sensei
    Flak
    Adder
    Lash
    Hawke
    Sturm
)

// Map for looking up a CO using its corresponding string
var COReverseStrings = map[string]CO {
    "Andy": Andy,
    "Max": Max,
    "Sami": Sami,
    "Nell": Nell,
    "Hachi": Hachi,
    "Olaf": Olaf,
    "Grit": Grit,
    "Colin": Colin,
    "Eagle": Eagle,
    "Drake": Drake,
    "Jess": Jess,
    "Kanbei": Kanbei,
    "Sonja": Sonja,
    "Sensei": Sensei,
    "Flak": Flak,
    "Adder": Adder,
    "Lash": Lash,
    "Hawke": Hawke,
    "Sturm": Sturm,
}

// CO frame type enum (all possible CO images)
type COFrameType uint8

const(
    Body COFrameType = iota
    FaceNeutral
    FaceGood
    FaceBad
)

// Map for looking up a CO face type using its corresponding string
var COFrameTypeReverseStrings = map[string]COFrameType{
    "Body": Body,
    "FaceNeutral": FaceNeutral,
    "FaceGood": FaceGood,
    "FaceBad": FaceBad,
}
