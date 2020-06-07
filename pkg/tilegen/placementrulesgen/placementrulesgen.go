package placementrulesgen

import (
	"github.com/turnabout/awodatagen"
	"github.com/turnabout/awodatagen/internal/genio"
	"github.com/turnabout/awodatagen/internal/utilities"
	"github.com/turnabout/awodatagen/pkg/tilegen/autovargen"
	"github.com/turnabout/awodatagen/pkg/tilegen/tiledata"
)

// Raw placement rule positions and their corresponding X/Y offset values
var rawPlacementRulePositions = map[string][2]int{
	"Middle":      {0, 0},
	"TopLeft":     {-1, -1},
	"TopRight":    {1, -1},
	"BottomLeft":  {-1, 1},
	"BottomRight": {1, 1},
	"Bottom":      {0, 1},
	"Right":       {1, 0},
	"Top":         {0, -1},
	"Left":        {-1, 0},
}

func AttachTilesPlacementRulesData(tilesData *tiledata.TileData) {

	// Load raw auto var data file into structure
	var rawData rawTilePlacementRules

	genio.AttachJSONData(
		utilities.GetInputPath(awodatagen.OtherDir, awodatagen.TilesPlacementRulesFileName),
		&rawData,
	)

	// Loop every tile type in the raw data
	for tileTypeStr, rawPlacementRules := range rawData {

		// Get the actual tile type for this raw data
		var tileType tiledata.TileType = tiledata.TileReverseStrings[tileTypeStr]

		// Create initial slice for this tile type's placement rules
		var placementRules []tiledata.TilePlacementRule

		// Loop raw rules & process
		for _, rawPlacementRule := range rawPlacementRules {
			placementRules = append(placementRules, processPlacementRule(rawPlacementRule))
		}

		// Store final result in tile data object
		(*tilesData)[tileType].PlacementRules = placementRules
	}
}

// Process a raw tile placement rule into an exportable tile placement rule
func processPlacementRule(rawRule rawTilePlacementRule) tiledata.TilePlacementRule {

	var result []tiledata.TilePlacementRuleComponent

	// Process every individual placement rule in this batch
	for _, rawRuleComponent := range rawRule {
		result = append(
			result,
			tiledata.TilePlacementRuleComponent{
				OffsetX: rawPlacementRulePositions[rawRuleComponent.Position][0],
				OffsetY: rawPlacementRulePositions[rawRuleComponent.Position][1],
				Tiles:   autovargen.ProcessAdjTileStr(rawRuleComponent.Tiles),
			},
		)
	}

	return result
}
