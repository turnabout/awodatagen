package main

func attachTilesClockData(tilesDir string, tilesData *TilesData) {

    var tilesClockData map[string]TileClockData

    attachJSONData(tilesDir + tilesClockDataFileName, &tilesClockData)

    for tileStr := range tilesClockData {
        tileType := tileReverseStrings[tileStr]
        data := tilesClockData[tileStr]

        (*tilesData)[tileType].ClockData = &data
    }
}
