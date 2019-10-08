package main

import "fmt"


var autoVarStrings = map[string]int{
    "any": 0xFF,
    "shadowing": int(Forest) | int(Mountain) | int(Silo),
    "oob": int(OOB),
}

// Attach auto var data to accumulated tiles data
func attachTilesAutoVarData(tilesDir string, vData *TilesData) {
    var rawData RawAutoVarsData

    // Load raw auto var data file into structure
    attachJSONData(tilesDir + tilesAutoVarFileName, &rawData)

    // Loop every tile type
    for tileTypeStr, tileTypeAutoVars := range rawData {
        var tileType TileType = tileReverseStrings[tileTypeStr]

        // Temporary: only process Forest
        if tileType != Forest {
            continue
        }

        fmt.Printf("%s\n", tileType)

        // Add initial slice for the tile type
        vData.Src[tileType].AutoVars = []AutoVarData{}

        // Loop auto var values, appending every one of them to this tile type's AutoVars field
        for _, autoVarData := range tileTypeAutoVars {

            vData.Src[tileType].AutoVars = append(
                vData.Src[tileType].AutoVars,
                processRawAutoVar(autoVarData),
            )
        }
    }
}

// Process the adjacent tiles in a raw autovar data struct and produce a final exported struct, containing a short
// string version of the tile variation and 4 bit field numbers representing the acceptable adjacent tiles for this var.
func processRawAutoVar(rawAutoVarData RawAutoVarData) AutoVarData {

    var tileVar TileVariation = tileVarsReverseStrings[rawAutoVarData.TileVar]

    result := AutoVarData{
        TileVar: tileVar.String(),
        AdjacentTiles: [4]int{0, 0, 0, 0},
    }

    fmt.Printf("%s\n", tileVar.String())

    // Process every adjacent tile string into a bit field number representing acceptable tile types
    for i := 0; i < ADJACENT_TILE_AMOUNT; i++ {
        var adjTileStr string = rawAutoVarData.AdjacentTiles[i]

        fmt.Printf("%s - ", adjTileStr)

    }
    fmt.Printf("\n")
    fmt.Printf("%#v\n\n", result)

    return result
}
