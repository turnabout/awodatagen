package tilegen

import (
    "github.com/turnabout/awossgen"
    "github.com/turnabout/awossgen/pkg/genio"
)

func attachTilesClockData(tileData *awossgen.TileData) {

    var tilesClockData map[string]awossgen.TileClockData

    // Fill out map with keys being tile short strings and values being tile clock data
    genio.AttachJSONData(
        awossgen.GetInputPath(awossgen.TilesDir, awossgen.TilesClockDataFileName),
        &tilesClockData,
    )

    // Attach tile clock data to tile data object using the map
    for tileStr := range tilesClockData {
        tileType := awossgen.TileReverseStrings[tileStr]
        data := tilesClockData[tileStr]

        (*tileData)[tileType].ClockData = &data
    }
}
