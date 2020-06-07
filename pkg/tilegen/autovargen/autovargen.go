package autovargen

import (
	"github.com/turnabout/awodatagen"
	"github.com/turnabout/awodatagen/pkg/genio"
	"github.com/turnabout/awodatagen/pkg/tilegen/tiledata"
	"github.com/turnabout/awodatagen/pkg/utilities"
	"sort"
)

// Attach auto-var data to tile data object
func AttachTilesAutoVarData(tilesData *tiledata.TileData) {
	var rawData rawAutoVarsData

	// Load raw auto var data file into structure
	genio.AttachJSONData(utilities.GetInputPath(awodatagen.OtherDir, awodatagen.TilesAutoVarFileName), &rawData)

	// Loop every tile type in the raw data
	for tileTypeStr, tileTypeAutoVars := range rawData {

		// Get the actual tile type for this raw data
		var tileType tiledata.TileType = tiledata.TileReverseStrings[tileTypeStr]

		// Create initial slice containing this tile type's auto-var data
		var tileAutoVar []tiledata.AutoVarData

		// Loop auto var values, appending every one of them to this tile type's AutoVars field
		for _, rawAutoVarData := range tileTypeAutoVars {
			tileAutoVar = append(tileAutoVar, processRawAutoVar(rawAutoVarData))
		}

		// Sort auto var data by amount of active bits
		sort.Slice(tileAutoVar, func(i, j int) bool {
			return getAutoVarBitsAmount(tileAutoVar[i]) < getAutoVarBitsAmount(tileAutoVar[j])
		})

		// Store final result in tile data object
		(*tilesData)[tileType].AutoVars = tileAutoVar
	}
}

// Gets the total amount of active bits in an auto var data struct
func getAutoVarBitsAmount(autoVarData tiledata.AutoVarData) uint {
	var totalBits uint = 0

	for i := 0; i < adjacentTileCount; i++ {
		totalBits += utilities.CountBits(uint(autoVarData.AdjacentTiles[i]))
	}

	return totalBits
}

// Processes raw auto-var data into a finished auto-var data struct
func processRawAutoVar(rawAutoVarData rawAutoVarData) tiledata.AutoVarData {

	// Get the tile variation corresponding to the raw auto-var data full string
	var tileVar tiledata.TileVariation = tiledata.TileVarsReverseStrings[rawAutoVarData.TileVar]

	// Create initial result to be filled out
	result := tiledata.AutoVarData{
		TileVar:       tileVar.String(),
		AdjacentTiles: [4]uint{0, 0, 0, 0},
	}

	// Process every adjacent tile string into a bit field number representing acceptable tile types
	for i := 0; i < adjacentTileCount; i++ {
		result.AdjacentTiles[i] = ProcessAdjTileStr(rawAutoVarData.AdjacentTiles[i])
	}

	return result
}
