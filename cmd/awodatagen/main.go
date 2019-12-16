package main

import (
    "github.com/turnabout/awodatagen"
    "github.com/turnabout/awodatagen/pkg/genio"
    "github.com/turnabout/awodatagen/pkg/packer"
    "github.com/turnabout/awodatagen/pkg/palettegen"
    "github.com/turnabout/awodatagen/pkg/propertygen"
    "github.com/turnabout/awodatagen/pkg/tilegen"
    "github.com/turnabout/awodatagen/pkg/uigen"
    "github.com/turnabout/awodatagen/pkg/unitgen"
    "image"
    "log"
)

func main() {
    log.SetFlags(log.LstdFlags | log.Lshortfile)

    // Gather packed frame images for all sections of the sprite sheet
    var packedTileFrameImages []packer.FrameImage
    var packedUnitFrameImages []packer.FrameImage
    var packedUIFrameImages   []packer.FrameImage
    var ssImg *image.RGBA

    ssImg = gatherFrameImages(&packedTileFrameImages, &packedUnitFrameImages, &packedUIFrameImages)

    // Create game data object using the frame images
    var gameData = awodatagen.GameData{
        Tiles:      *tilegen.GetTileData(&packedTileFrameImages),
        Properties: *propertygen.GetPropertyData(&packedTileFrameImages),
        Units:      *unitgen.GetUnitData(&packedUnitFrameImages),
        UI:         *uigen.GetUIData(&packedUIFrameImages),

        SpriteSheetDimensions: awodatagen.SSDimensions{
            Width: ssImg.Bounds().Max.X,
            Height: ssImg.Bounds().Max.Y,
        },
    }

    attachAdditionalVData(&gameData)

    // Output results
    genio.OutputJSON(&gameData)
    genio.OutputSpriteSheet(ssImg)
}

// Gather additional visual data and attach to the main visual data object
func attachAdditionalVData(gameData *awodatagen.GameData) {
    genio.AttachJSONData(
        awodatagen.GetInputPath(awodatagen.AdditionalDir, awodatagen.StagesFileName),
        &gameData.Stages,
    )

    genio.AttachJSONData(
        awodatagen.GetInputPath(awodatagen.AdditionalDir, awodatagen.AnimClocksFileName),
        &gameData.AnimationClocks,
    )

    palettegen.AttachPaletteData(gameData)
}

func gatherFrameImages(
    packedTileFrameImagesOut *[]packer.FrameImage,
    packedUnitFrameImagesOut*[]packer.FrameImage,
    packedUIFrameImagesOut*[]packer.FrameImage,
) *image.RGBA {
    // 1. Gather tiles/properties frame images
    var tileFrameImages []packer.FrameImage

    tilegen.GetTileFrameImgs(&tileFrameImages)
    propertygen.GetPropertyFrameImgs(&tileFrameImages)
    packedTileFrameImages, tilesSectionWidth, tilesSectionHeight := packer.Pack(&tileFrameImages)

    accumImg := packer.DrawPackedFrames(packedTileFrameImages, tilesSectionWidth, tilesSectionHeight)
    *packedTileFrameImagesOut = *packedTileFrameImages

    // 2. Gather units frame images
    // Start off the frame images with previously accumulated image including tiles
    var unitsFrameImages []packer.FrameImage = []packer.FrameImage{
        {
            Image: accumImg,
            Width: tilesSectionWidth,
            Height: tilesSectionHeight,
            MetaData: packer.FrameImageMetaData{
                FrameImageDataType: uint8(awodatagen.OtherDataType),
            },
        },
    }

    unitgen.GetUnitFrameImgs(&unitsFrameImages)
    packedUnitFrameImages, unitsSectionWidth, unitsSectionHeight := packer.Pack(&unitsFrameImages)

    accumImg = packer.DrawPackedFrames(packedUnitFrameImages, unitsSectionWidth, unitsSectionHeight)
    *packedUnitFrameImagesOut = *packedUnitFrameImages

    // 3. Gather UI frame images
    // Start off the frame images with previously accumulated image including tiles & units
    var UIFrameImages []packer.FrameImage = []packer.FrameImage{
        {
            Image: accumImg,
            Width: unitsSectionWidth,
            Height: unitsSectionHeight,
            MetaData: packer.FrameImageMetaData{
                FrameImageDataType: uint8(awodatagen.OtherDataType),
            },
        },
    }

    uigen.GetUIFrameImgs(&UIFrameImages)
    packedUIFrameImages, uiSectionWidth, uiSectionHeight := packer.Pack(&UIFrameImages)

    accumImg = packer.DrawPackedFrames(packedUIFrameImages, uiSectionWidth, uiSectionHeight)
    *packedUIFrameImagesOut = *packedUIFrameImages

    return accumImg
}
