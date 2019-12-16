package tilegen

func attachTilesClockData(tilesData *TilesData) {

    var tilesClockData map[string]TileClockData

    attachJSONData( getFullProjectPath(tilesDir, tilesClockDataFileName), &tilesClockData )

    for tileStr := range tilesClockData {
        tileType := tileReverseStrings[tileStr]
        data := tilesClockData[tileStr]

        (*tilesData)[tileType].ClockData = &data
    }
}
