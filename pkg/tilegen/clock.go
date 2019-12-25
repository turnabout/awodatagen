package tilegen

import (
    "github.com/turnabout/awodatagen"
    "github.com/turnabout/awodatagen/pkg/genio"
)

func attachTilesClockData(tileData *awodatagen.TileData) {
    return
    var tilesClockData map[string]awodatagen.TileClockData

    // Fill out map with keys being tile short strings and values being tile clock data
    genio.AttachJSONData(
        awodatagen.GetInputPath(awodatagen.OtherDir, awodatagen.TilesClockDataFileName),
        &tilesClockData,
    )

    // Attach tile clock data to tile data object using the map
    for tileStr := range tilesClockData {
        tileType := awodatagen.TileReverseStrings[tileStr]
        data := tilesClockData[tileStr]

        (*tileData)[tileType].ClockData = &data
    }
}
