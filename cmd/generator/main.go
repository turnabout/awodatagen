package main

import (
    "fmt"
    "github.com/turnabout/awossgen"
    "github.com/turnabout/awossgen/pkg/packer"
    "image"
    "log"
)

func main() {
    log.SetFlags(log.LstdFlags | log.Lshortfile)

    // Gather packed frame images for all sections of the sprite sheet
    var packedTileFrameImages []packer.FrameImage
    var packedUnitFrameImages []packer.FrameImage
    var packedUIFrameImages []packer.FrameImage
    var ssImg *image.RGBA

    ssImg = gatherFrameImages(&packedTileFrameImages, &packedUnitFrameImages, &packedUIFrameImages)

    fmt.Printf("%#v\n", ssImg)

    /*
        // Create game data object using the frame images
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

        // Output game data file and sprite sheet image
        outputJSON(&gameData)
        outputSpriteSheet(accumImg)
    */
}

// Gather additional visual data and attach to the main visual data object
func attachAdditionalVData(gameData *awossgen.GameData) {
    // attachJSONData( awossgen.GetInputPath(awossgen.AdditionalDir, awossgen.StagesFileName), &gameData.Stages )
    // attachJSONData( awossgen.GetInputPath(awossgen.AdditionalDir, awossgen.AnimClocksFileName), &gameData.AnimationClocks)
    // palettegen.AttachPaletteData(gameData)
}

func gatherFrameImages(
    packedTileFrameImages *[]packer.FrameImage,
    packedUnitFrameImages *[]packer.FrameImage,
    packedUIFrameImages *[]packer.FrameImage,
) *image.RGBA {
    /*
        // 1. Gather tiles/properties frame images
        var tilesFrameImgs []packer.FrameImage

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
                    FrameImageDataType: SpriteSheetSectionFrameImage,
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
                    FrameImageDataType: SpriteSheetSectionFrameImage,
                },
            },
        }

        getUISrcFrameImgs(&uiFrameImgs)

        packedUiFrameImgs, uiSectionWidth, uiSectionHeight := pack(&uiFrameImgs)
        accumImg = drawPackedFrames(packedUiFrameImgs, uiSectionWidth, uiSectionHeight)
    */

    var test image.RGBA
    return &test
}
