package tilegen

import (
	"github.com/turnabout/awodatagen"
	"github.com/turnabout/awodatagen/pkg/genio"
	"github.com/turnabout/awodatagen/pkg/tilegen/tiledata"
)

func attachTilesClockData(tileData *tiledata.TileData) {
	var tilesClockData map[string]tiledata.TileClockData

	// Fill out map with keys being tile short strings and values being tile clock data
	genio.AttachJSONData(
		awodatagen.GetInputPath(awodatagen.OtherDir, awodatagen.TilesClockDataFileName),
		&tilesClockData,
	)

	// Loop filled map & use to fill tile variations' clock data
	for tileStr, tileTypeClockData := range tilesClockData {
		tileType := tiledata.TileReverseStrings[tileStr]

		// Initially set all variations to the default clock value
		for varStr, varData := range (*tileData)[tileType].Variations {

			// Set variation data clock index to the default clock index
			clockIndex := tileTypeClockData.DefaultClock
			varData.ClockIndex = &clockIndex

			// Set back the variation data in the original tile data
			(*tileData)[tileType].Variations[varStr] = varData
		}

		// Override the clock for variations that have a specific value
		for varStr, varClock := range tileTypeClockData.VarClocks {

			// Get the looped variation & data object for this variation
			loopedVar := tiledata.TileVarsReverseStrings[varStr]
			varData := (*tileData)[tileType].Variations[loopedVar.String()]

			// Set the variation data clock index
			clockIndex := varClock
			varData.ClockIndex = &clockIndex

			// Set back the variation data in the original tile data
			(*tileData)[tileType].Variations[loopedVar.String()] = varData
		}
	}
}
