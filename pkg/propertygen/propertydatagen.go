package propertygen

import (
    "github.com/turnabout/awossgen"
    "github.com/turnabout/awossgen/pkg/packer"
)

// Generate properties' sprite sheet & visual data
func GetPropertyData(packedFrameImgs *[]packer.FrameImage) *awossgen.PropertiesData {

    // Get the base properties data object containing frame source data
    var propsData *awossgen.PropertiesData = getBasePropertyData(packedFrameImgs)

    // Attach additional data to the properties data
    attachExtraPropData(propsData)

    return propsData
}

// Generate the visual data for Properties' origin
func getBasePropertyData(packedFrameImgs *[]packer.FrameImage) *awossgen.PropertiesData {

    // Weather Variation -> Property Type -> Unit Variation
    propsData := make(awossgen.PropertiesData, awossgen.PropWeatherVarCount)

    // Initialize Property Type arrays
    for weatherVar := range propsData {
        propsData[weatherVar] = make([][]awossgen.Frame, awossgen.PropertyTypeAmount)

        // Initialize Unit Variation arrays
        for propType := range propsData[weatherVar] {
            var unitVarAmount int

            // HQ Properties have all Unit Variations, while other properties only have one
            if awossgen.PropertyType(propType) == awossgen.HQ {
                unitVarAmount = int(awossgen.UnitVariationAmount)
            } else {
                unitVarAmount = 1
            }

            propsData[weatherVar][propType] = make([]awossgen.Frame, unitVarAmount)
        }
    }

    // Fill out Src visual data
    for _, frameImg := range *packedFrameImgs {

        // Ignore non-tile frame images
        if frameImg.MetaData.FrameImageDataType != uint8(awossgen.PropertyDataType) {
            continue
        }

        propsData[frameImg.MetaData.Variation][frameImg.MetaData.Type][frameImg.MetaData.Animation] = awossgen.Frame{
            X: frameImg.X,
            Y: frameImg.Y,
            Height: frameImg.Height,
        }
    }

    return &propsData
}

// Attach extra data to property data
func attachExtraPropData(propData *awossgen.PropertiesData) {
    // propsDir := baseDirPath + inputsDirName + propertiesDir
    // attachJSONData(propsDir + palettesFileName, &propData.Palettes)
    // attachJSONData(propsDir +propsLightsRGBFileName, &propData.PropsLightsRGB)
}
