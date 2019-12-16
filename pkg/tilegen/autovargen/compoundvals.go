package autovargen

import "github.com/turnabout/awossgen"

// Outputs bit field from multiple tile types, used to populate autoVarCompoundVals
func makeCompoundVal(values []awossgen.TileType) uint {
    var result uint = 0

    for _, val := range values {
        result |= (1 << val)
    }

    return result
}

// Values corresponding to auto var compound symbols
var autoVarCompoundVals = map[string]uint{
    "any":       0xFFFFFFFF,

    "shadowing": makeCompoundVal([]awossgen.TileType{
        awossgen.Forest,
        awossgen.Mountain,
        awossgen.Silo,
    }),

    "oob": makeCompoundVal([]awossgen.TileType{
        awossgen.OOB,
    }),

    "land": makeCompoundVal([]awossgen.TileType{
        awossgen.Plain,
        awossgen.Forest,
        awossgen.Mountain,
        awossgen.Road,
        awossgen.Bridge,
        awossgen.Pipe,
        awossgen.PipeFragile,
        awossgen.Silo,
    }),
}
