package autovargen

import (
	"github.com/turnabout/awodatagen/pkg/tilegen/tiledata"
)

// Outputs bit field from multiple tile types, used to populate autoVarCompoundVals
func makeCompoundVal(values []tiledata.TileType) uint {
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
	"shadowing": makeCompoundVal([]tiledata.TileType{
		tiledata.Forest,
		tiledata.Mountain,
		tiledata.Silo,
		tiledata.PropertyHQ,
		tiledata.PropertyCity,
		tiledata.PropertyBase,
		tiledata.PropertyAirport,
		tiledata.PropertyPort,
	}),

	// Out of bounds tile
	"oob": makeCompoundVal([]tiledata.TileType{
		tiledata.OOB,
	}),

	// Tiles that make up the land
	"land": makeCompoundVal([]tiledata.TileType{
		tiledata.Plain,
		tiledata.Forest,
		tiledata.Mountain,
		tiledata.Road,
		tiledata.River,
		tiledata.Pipe,
		tiledata.PipeFragile,
		tiledata.Silo,
		tiledata.PropertyHQ,
		tiledata.PropertyCity,
		tiledata.PropertyBase,
		tiledata.PropertyAirport,
		tiledata.PropertyPort,
	}),

	// Tiles that make up the sea
	"sea": makeCompoundVal([]tiledata.TileType{
		tiledata.Sea,
		tiledata.Shore,
		tiledata.Reef,
		tiledata.OOB,
	}),
}
