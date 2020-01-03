package placementrulesgen

import (
    "fmt"
    "github.com/turnabout/awodatagen"
    "github.com/turnabout/awodatagen/pkg/genio"
)

func AttachTilesPlacementRulesData(tilesData *awodatagen.TileData) {

    // Load raw auto var data file into structure
    var rawData rawTilePlacementRules

    genio.AttachJSONData(
        awodatagen.GetInputPath(awodatagen.OtherDir, awodatagen.TilesPlacementRulesFileName),
        &rawData,
    )

    // Loop every tile type in the raw data
    for tileTypeStr, rawPlacementRules := range rawData {

        // Get the actual tile type for this raw data
        var tileType awodatagen.TileType = awodatagen.TileReverseStrings[tileTypeStr]
        fmt.Printf("%s\n", tileType.String())

        // Create initial slice for this tile type's placement rules
        var placementRules [][]awodatagen.TilePlacementRule


        // Loop raw rules values, appending every one of them to this tile type's AutoVars field
        for _, rawTilePlacementRule := range rawPlacementRules {

            fmt.Printf("%s\n", rawTilePlacementRule)

            placementRules = append(placementRules, processRawPlacementRuleBatch(rawPlacementRules))
        }
    }

}

func processRawPlacementRuleBatch(rawRules [][]rawTilePlacementRule) []awodatagen.TilePlacementRule {

    var result []awodatagen.TilePlacementRule

    /*
        // Create initial result to be filled out
        result := awodatagen.AutoVarData{
            TileVar: tileVar.String(),
            AdjacentTiles: [4]uint{0, 0, 0, 0},
        }

        // Process every adjacent tile string into a bit field number representing acceptable tile types
        for i := 0; i < adjacentTileCount; i++ {
            result.AdjacentTiles[i] = ProcessAdjTileStr(rawAutoVarData.AdjacentTiles[i])
        }

        return result
    */

    return result
}
