//
// awossgen generates the sprite sheet & data JSON file used by AWO.
//
package main

func main() {
    var vData = VisualData{
        Units:      getUnitsData(),
        Tiles:      getTilesData(),
        Properties: getPropertiesData(),
        SSMetaData: ssMetaData{},
    }

    attachAdditionalVData(&vData)
    outputSpriteSheet(joinSpriteSheets(&vData))
    outputJSON(&vData)
}
