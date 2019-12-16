package autovargen

import (
    "github.com/turnabout/awossgen"
    "github.com/turnabout/awossgen/pkg/genio"
    "sort"
)

// Attach auto-var data to tile data object
func AttachTilesAutoVarData(tilesData *awossgen.TileData) {
    var rawData rawAutoVarsData

    // Load raw auto var data file into structure
    genio.AttachJSONData( awossgen.GetInputPath(awossgen.TilesDir, awossgen.TilesAutoVarFileName), &rawData )

    // Loop every tile type in the raw data
    for tileTypeStr, tileTypeAutoVars := range rawData {

        // Get the actual tile type for this raw data
        var tileType awossgen.TileType = awossgen.TileReverseStrings[tileTypeStr]

        // TODO: remove temporary debug condition
        if tileType != awossgen.Forest && tileType != awossgen.Plain && tileType != awossgen.Bridge && tileType != awossgen.River && tileType != awossgen.Sea {
            continue
        }

        // Create initial slice containing this tile type's auto-var data
        var tileAutoVar []awossgen.AutoVarData

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
func getAutoVarBitsAmount(autoVarData awossgen.AutoVarData) uint {
    var totalBits uint = 0

    for i := 0; i < adjacentTileCount; i++ {
        totalBits += countBits(uint(autoVarData.AdjacentTiles[i]))
    }

    return totalBits
}

// Processes raw auto-var data into a finished auto-var data struct
func processRawAutoVar(rawAutoVarData rawAutoVarData) awossgen.AutoVarData {

    // Get the tile variation corresponding to the raw auto-var data full string
    var tileVar awossgen.TileVariation = awossgen.TileVarsReverseStrings[rawAutoVarData.TileVar]

    // Create initial result to be filled out
    result := awossgen.AutoVarData{
        TileVar: tileVar.String(),
        AdjacentTiles: [4]uint{0, 0, 0, 0},
    }

    // Process every adjacent tile string into a bit field number representing acceptable tile types
    for i := 0; i < adjacentTileCount; i++ {
        result.AdjacentTiles[i] = processAdjTileStr(rawAutoVarData.AdjacentTiles[i])
    }

    return result
}
