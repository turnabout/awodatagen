// Generates sprite sheets & data used by AWO (outputs visuals.json & spritesheet.png)
package main

import (
    _ "image/png"
)

func main() {
    var vData = VisualData{
        Units:      getUnitsData(),
        Tiles:      getTilesData(),
        Properties: getPropertiesData(),
        SSMetaData: ssMetaData{},
    }

    outputSpriteSheet(joinSpriteSheets(&vData))
    outputJSON(&vData)
}
