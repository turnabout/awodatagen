//
// awossgen generates the sprite sheet & data JSON file used by AWO.
//
package main

import "log"

func main() {

    log.SetFlags(log.LstdFlags | log.Lshortfile)

    // 1. Gather tiles/properties frame images
    var tilesFrameImgs []FrameImage

    getTilesSrcFrameImgs(&tilesFrameImgs)
    getPropsSrcFrameImgs(&tilesFrameImgs)

    packedTilesFrameImgs, tilesSectionWidth, tilesSectionHeight := pack(&tilesFrameImgs)
    accumImg := drawPackedFrames(packedTilesFrameImgs, tilesSectionWidth, tilesSectionHeight)

    // 2. Gather units frame images
    // Start off the frame images with previously accumulated image including tiles
    var unitsFrameImgs []FrameImage = []FrameImage{
        {
            Image: accumImg,
            Width: tilesSectionWidth,
            Height: tilesSectionHeight,
            MetaData: FrameImageMetaData{
                FrameImageType: SpriteSheetSectionFrameImage,
            },
        },
    }

    getUnitsSrcFrameImgs(&unitsFrameImgs)

    packedUnitsFrameImgs, unitsSectionWidth, unitsSectionHeight := pack(&unitsFrameImgs)
    accumImg = drawPackedFrames(packedUnitsFrameImgs, unitsSectionWidth, unitsSectionHeight)

    // 3. Gather UI frame images
    // Start off the frame images with previously accumulated image including tiles & units
    var uiFrameImgs []FrameImage = []FrameImage{
        {
            Image: accumImg,
            Width: unitsSectionWidth,
            Height: unitsSectionHeight,
            MetaData: FrameImageMetaData{
                FrameImageType: SpriteSheetSectionFrameImage,
            },
        },
    }

    getUISrcFrameImgs(&uiFrameImgs)

    packedUiFrameImgs, uiSectionWidth, uiSectionHeight := pack(&uiFrameImgs)
    accumImg = drawPackedFrames(packedUiFrameImgs, uiSectionWidth, uiSectionHeight)

    // Create visual data object using the frame images
    var gameData = GameData{
        Tiles: *getTilesData(packedTilesFrameImgs),
        Properties: *getPropertiesData(packedTilesFrameImgs),
        Units: *getUnitsData(packedUnitsFrameImgs),
        UI: *getUiData(packedUiFrameImgs),
        SpriteSheetDimensions: ssDimensions{
            Width: unitsSectionWidth,
            Height: unitsSectionHeight,
        },
    }

    attachAdditionalVData(&gameData)

    // Output all results
    outputJSON(&gameData)
    outputSpriteSheet(accumImg)
}
