package propertygen

import (
    "github.com/turnabout/awodatagen"
    "github.com/turnabout/awodatagen/pkg/packer"
)

// Generate properties' sprite sheet & visual data
func GetPropertyData(packedFrameImgs *[]packer.FrameImage) *awodatagen.PropertyData {

    // Get the base properties data object containing frame source data
    var propsData *awodatagen.PropertyData = getBasePropertyData(packedFrameImgs)

    // Attach additional data to the properties data
    attachExtraPropData(propsData)

    return propsData
}

// Generate the visual data for Properties' origin
func getBasePropertyData(packedFrameImgs *[]packer.FrameImage) *awodatagen.PropertyData {

    // Weather Variation -> Property Type -> Unit Variation
    propsData := make(awodatagen.PropertyData, awodatagen.PropWeatherVarCount)

    // Initialize Property Type arrays
    for weatherVar := range propsData {
        propsData[weatherVar] = make([][]awodatagen.Frame, awodatagen.PropTypeCount)

        // Initialize Unit Variation arrays
        for propType := range propsData[weatherVar] {
            var unitVarAmount int

            // HQ Properties have all Unit Variations, while other properties only have one
            if awodatagen.PropertyType(propType) == awodatagen.HQ {
                unitVarAmount = int(awodatagen.ArmyTypeCount)
            } else {
                unitVarAmount = 1
            }

            propsData[weatherVar][propType] = make([]awodatagen.Frame, unitVarAmount)
        }
    }

    // Fill out Src visual data
    for _, frameImg := range *packedFrameImgs {

        // Ignore non-tile frame images
        if frameImg.MetaData.FrameImageDataType != uint8(awodatagen.PropertyDataType) {
            continue
        }

        propsData[frameImg.MetaData.Variation][frameImg.MetaData.Type][frameImg.MetaData.Animation] = awodatagen.Frame{
            X: frameImg.X,
            Y: frameImg.Y,
            Height: frameImg.Height,
        }
    }

    return &propsData
}

// Attach extra data to property data
func attachExtraPropData(propData *awodatagen.PropertyData) {
    // propsDir := baseDirPath + inputsDirName + propertiesDir
    // attachJSONData(propsDir + palettesFileName, &propData.Palettes)
    // attachJSONData(propsDir +propsLightsRGBFileName, &propData.PropsLightsRGB)
}
