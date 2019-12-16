package autovargen

import "github.com/turnabout/awodatagen"

// Outputs bit field from multiple tile types, used to populate autoVarCompoundVals
func makeCompoundVal(values []awodatagen.TileType) uint {
    var result uint = 0

    for _, val := range values {
        result |= (1 << val)
    }

    return result
}

// Values corresponding to auto var compound symbols
var autoVarCompoundVals = map[string]uint{
    "any":       0xFFFFFFFF,

    "shadowing": makeCompoundVal([]awodatagen.TileType{
        awodatagen.Forest,
        awodatagen.Mountain,
        awodatagen.Silo,
    }),

    "oob": makeCompoundVal([]awodatagen.TileType{
        awodatagen.OOB,
    }),

    "land": makeCompoundVal([]awodatagen.TileType{
        awodatagen.Plain,
        awodatagen.Forest,
        awodatagen.Mountain,
        awodatagen.Road,
        awodatagen.Bridge,
        awodatagen.Pipe,
        awodatagen.PipeFragile,
        awodatagen.Silo,
    }),
}
