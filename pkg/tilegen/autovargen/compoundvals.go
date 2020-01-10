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

    // All tiles
    "any": 0xFFFFFFFF,

    // Tiles that cast a shadow onto certain tiles to their right
    "shadowing": makeCompoundVal([]awodatagen.TileType{
        awodatagen.Forest,
        awodatagen.Mountain,
        awodatagen.Silo,
        awodatagen.PropertyHQ,
        awodatagen.PropertyCity,
        awodatagen.PropertyBase,
        awodatagen.PropertyAirport,
        awodatagen.PropertyPort,
    }),

    // Out of bounds tile
    "oob": makeCompoundVal([]awodatagen.TileType{
        awodatagen.OOB,
    }),

    // Tiles that make up the land
    "land": makeCompoundVal([]awodatagen.TileType{
        awodatagen.Plain,
        awodatagen.Forest,
        awodatagen.Mountain,
        awodatagen.Road,
        awodatagen.River,
        awodatagen.Pipe,
        awodatagen.PipeFragile,
        awodatagen.Silo,
        awodatagen.PropertyHQ,
        awodatagen.PropertyCity,
        awodatagen.PropertyBase,
        awodatagen.PropertyAirport,
        awodatagen.PropertyPort,
    }),

    // Tiles that make up the sea
    "sea": makeCompoundVal([]awodatagen.TileType{
        awodatagen.Sea,
        awodatagen.Shore,
        awodatagen.Reef,
        awodatagen.OOB,
    }),
}
