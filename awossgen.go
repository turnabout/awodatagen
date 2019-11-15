//
// awossgen generates the sprite sheet & data JSON file used by AWO.
//
package main

func main() {

    // Gather tiles frame images
    var tilesFrameImgs []FrameImage

    getTilesSrcFrameImgs(&tilesFrameImgs)
    getPropsSrcFrameImgs(&tilesFrameImgs)

    packedTilesFrameImgs, tilesSectionWidth, tilesSectionHeight := pack(&tilesFrameImgs)
    tilesImg := drawPackedFrames(packedTilesFrameImgs, tilesSectionWidth, tilesSectionHeight)

    // Gather units frame images
    // Start off the frame images with the previously accumulated image including tiles
    var unitsFrameImgs []FrameImage = []FrameImage{
        {
            Image: tilesImg,
            Width: tilesSectionWidth,
            Height: tilesSectionHeight,
            MetaData: FrameImageMetaData{
                FrameImageType: SpriteSheetSectionFrameImage,
            },
        },
    }

    getUnitsSrcFrameImgs(&unitsFrameImgs)

    packedUnitsFrameImgs, unitsSectionWidth, unitsSectionHeight := pack(&unitsFrameImgs)
    unitsImg := drawPackedFrames(packedUnitsFrameImgs, unitsSectionWidth, unitsSectionHeight)

    // Create visual data object using the frame images
    var gameData = VisualData{
        Tiles: getTilesData(packedTilesFrameImgs),
        Properties: getPropertiesData(packedTilesFrameImgs),
        Units: getUnitsData(packedUnitsFrameImgs),
        SpriteSheetDimensions: ssDimensions{
            Width: unitsSectionWidth,
            Height: unitsSectionHeight,
        },
    }

    attachAdditionalVData(&gameData)

    // Output all results
    outputJSON(&gameData)
    outputSpriteSheet(unitsImg)
}
