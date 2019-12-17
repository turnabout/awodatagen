package main

import (
    "github.com/turnabout/awodatagen"
    "github.com/turnabout/awodatagen/pkg/cogen"
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

// Callback used to get frame images from a section of the game data (tiles, units, etc)
type getFrameImagesCB func(*[]packer.FrameImage)

func main() {
    log.SetFlags(log.LstdFlags | log.Lshortfile)

    // Gather all packed frame images & the sprite sheet image
    var packedTileFrameImages, packedUnitFrameImages, packedCOFrameImages, packedUIFrameImages []packer.FrameImage
    var ssImg *image.RGBA

    ssImg = gatherFrameImages(
        &packedTileFrameImages,
        &packedUnitFrameImages,
        &packedCOFrameImages,
        &packedUIFrameImages,
    )

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

// Gathers frame images from every category of entities making up the sprite sheet
func gatherFrameImages(
    packedTileFrameImagesOut *[]packer.FrameImage,
    packedUnitFrameImagesOut *[]packer.FrameImage,
    packedCOFrameImagesOut   *[]packer.FrameImage,
    packedUIFrameImagesOut   *[]packer.FrameImage,
) *image.RGBA {

    // Gather tiles/properties frame images
    accumImg := gatherStepFrameImages(
        packedTileFrameImagesOut,
        nil,
        tilegen.GetTileFrameImgs,
        propertygen.GetPropertyFrameImgs,
    )

    // Gather units frame images
    accumImg = gatherStepFrameImages(packedUnitFrameImagesOut, accumImg, unitgen.GetUnitFrameImgs)

    // Gather CO frame images
    accumImg = gatherStepFrameImages(packedCOFrameImagesOut, accumImg, cogen.GetCOFrameImgs)

    // Gather UI frame images
    accumImg = gatherStepFrameImages(packedUIFrameImagesOut, accumImg, uigen.GetUIFrameImgs)

    return accumImg
}

// Gathers the frame images for a single step.
// Can use one or many frame image callbacks to process frame images, pack them, use them to draw an image and return
// both the packed frame images and the drawn image.
func gatherStepFrameImages(
    packedFrameImagesOut *[]packer.FrameImage,
    accumImg *image.RGBA,
    frameImagesCBs ...getFrameImagesCB,
) *image.RGBA {

    var frameImages []packer.FrameImage

    // If an accumulated image is given, use it as a base for this step's frame images
    if accumImg != nil {
        frameImages = append(frameImages, packer.FrameImage{
            Image: accumImg,
            Width: accumImg.Bounds().Max.X,
            Height: accumImg.Bounds().Max.Y,
            MetaData: packer.FrameImageMetaData{
                FrameImageDataType: uint8(awodatagen.OtherDataType),
            },
        })
    }

    // Get the frame images for this step using the given callbacks
    for _, cb := range frameImagesCBs {
        cb(&frameImages)
    }

    // Pack the frame images and send as output
    packedFrameImages, sectionWidth, sectionHeight := packer.Pack(&frameImages)
    *packedFrameImagesOut = *packedFrameImages

    // Output the new accumulated image
    return packer.DrawPackedFrames(packedFrameImages, sectionWidth, sectionHeight)
}

// Gather additional visual data and attach to the main visual data object
func attachAdditionalVData(gameData *awodatagen.GameData) {

    // Adds default stages data
    genio.AttachJSONData(
        awodatagen.GetInputPath(awodatagen.AdditionalDir, awodatagen.StagesFileName),
        &gameData.Stages,
    )

    // Adds animation clocks data
    genio.AttachJSONData(
        awodatagen.GetInputPath(awodatagen.AdditionalDir, awodatagen.AnimClocksFileName),
        &gameData.AnimationClocks,
    )

    palettegen.AttachPaletteData(gameData)
}
