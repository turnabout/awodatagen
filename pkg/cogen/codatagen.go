package cogen

import (
    "github.com/turnabout/awodatagen"
    "github.com/turnabout/awodatagen/pkg/packer"
)

// Generates the CO-related game data
func GetCOData(packedFrameImgs *[]packer.FrameImage) *awodatagen.COData {
    var COData *awodatagen.COData = getCOBaseData(packedFrameImgs)

    return COData
}

func getCOBaseData(packedFrameImgs *[]packer.FrameImage) *awodatagen.COData {

    // CO -> CO Type Data
    data := make(awodatagen.COData, awodatagen.COCount)

    // Process frame images
    for _, frameImg := range *packedFrameImgs {

        // Ignore non-CO frame images
        if frameImg.MetaData.FrameImageDataType != uint8(awodatagen.CODataType) {
            continue
        }

        // Get metadata on the CO this frame image represents
        CO := awodatagen.CO(frameImg.MetaData.Type)
        army := awodatagen.ArmyType(frameImg.MetaData.Variation)
        frameType := awodatagen.COFrameType(frameImg.MetaData.Index)

        // Set data on this CO
        data[CO].Name = CO.String()
        data[CO].Army = army

        data[CO].Frames[frameType] = awodatagen.Frame{
            X: frameImg.X,
            Y: frameImg.Y,
            Width: frameImg.Width,
            Height: frameImg.Height,
        }
    }

    return &data
}
