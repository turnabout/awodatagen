package propertygen

import (
	"github.com/turnabout/awodatagen/pkg/framedata"
	"github.com/turnabout/awodatagen/pkg/packer"
	"github.com/turnabout/awodatagen/pkg/unitgen"
)

// Generate properties' sprite sheet & visual data
func GetPropertyData(packedFrameImgs *[]packer.FrameImage) *PropertyData {

	// Get the base properties data object containing frame source data
	var propsData *PropertyData = getBasePropertyData(packedFrameImgs)

	// Attach additional data to the properties data
	attachExtraPropData(propsData)

	return propsData
}

// Generate the visual data for Properties' origin
func getBasePropertyData(packedFrameImgs *[]packer.FrameImage) *PropertyData {

	// Weather Variation -> Property Type -> Unit Variation
	propsData := make(PropertyData, PropWeatherVarCount)

	// Initialize Property Type arrays
	for weatherVar := range propsData {
		propsData[weatherVar] = make([][]framedata.Frame, PropTypeCount)

		// Initialize Unit Variation arrays
		for propType := range propsData[weatherVar] {
			var unitVarAmount int

			// HQ Properties have all Unit Variations, while other properties only have one
			if PropertyType(propType) == HQ {
				unitVarAmount = int(unitgen.ArmyTypeCount)
			} else {
				unitVarAmount = 1
			}

			propsData[weatherVar][propType] = make([]framedata.Frame, unitVarAmount)
		}
	}

	// Fill out Src visual data
	for _, frameImg := range *packedFrameImgs {

		// Ignore non-tile frame images
		if frameImg.MetaData.FrameImageDataType != uint8(framedata.PropertyDataType) {
			continue
		}

		propsData[frameImg.MetaData.Variation][frameImg.MetaData.Type][frameImg.MetaData.Animation] = framedata.Frame{
			X:      frameImg.X,
			Y:      frameImg.Y,
			Height: frameImg.Height,
		}
	}

	return &propsData
}

// Attach extra data to property data
func attachExtraPropData(propData *PropertyData) {
	// propsDir := baseDirPath + inputsDirName + propertiesDir
	// attachJSONData(propsDir + palettesFileName, &propData.Palettes)
	// attachJSONData(propsDir +propsLightsRGBFileName, &propData.PropsLightsRGB)
}
