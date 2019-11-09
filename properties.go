// Generates Properties' visual data
package main

import (
    "os"
)

// Generate properties' sprite sheet & visual data
func getPropertiesData(packedFrameImgs *[]FrameImage) *PropertiesData {

    vData := PropertiesData{
        Src: *getPropsSrcVData(packedFrameImgs),
    }

    attachExtraPropsVData(&vData)
    return &vData
}

// Gather Frame Images for Properties' source
func getPropsSrcFrameImgs(frameImgs *[]FrameImage) {
    propsDir := baseDirPath + inputsDirName + propertiesDirName + "/"

    // Loop Weather Variations
    for weatherVar := FirstPropertyWeatherVariation; weatherVar <= LastPropertyWeatherVariation; weatherVar++ {
        weatherDir := propsDir + weatherVar.String() + "/"

        // Loop Property Types
        for propType := FirstPropertyType; propType <= LastPropertyType; propType++ {
            propDir := weatherDir + propType.String() + "/"

            // Loop Unit Variations
            for unitVar := FirstUnitVariation; unitVar <= LastUnitVariation; unitVar++ {
                fullPath := propDir + unitVar.String() + ".png"

                // Ignore this variation if it does not exist on this Property Type
                if _, err := os.Stat(fullPath); os.IsNotExist(err) {
                    continue
                }

                imageObj := getImage(fullPath)

                *frameImgs = append(*frameImgs, FrameImage{
                    Image: imageObj,
                    Width:  imageObj.Bounds().Max.X,
                    Height: imageObj.Bounds().Max.Y,
                    MetaData: FrameImageMetaData{
                        Type:           uint8(propType),
                        Variation:      uint8(weatherVar),
                        Animation:      uint8(unitVar),
                        FrameImageType: PropertyFrameImage,
                    },
                })
            }
        }
    }
}

// Generate the visual data for Properties' origin
func getPropsSrcVData(packedFrameImgs *[]FrameImage) *[][][]Frame {

    // Weather Variation -> Property Type -> Unit Variation
    originVData := make([][][]Frame, PropertyWeatherVariationAmount)

    // Initialize Property Type arrays
    for weatherVar := range originVData {
        originVData[weatherVar] = make([][]Frame, PropertyTypeAmount)

        // Initialize Unit Variation arrays
        for propType := range originVData[weatherVar] {
            var unitVarAmount int

            // HQ Properties have all Unit Variations, while other properties only have one
            if PropertyType(propType) == HQ {
                unitVarAmount = int(UnitVariationAmount)
            } else {
                unitVarAmount = 1
            }

            originVData[weatherVar][propType] = make([]Frame, unitVarAmount)
        }
    }

    // Fill out Src visual data
    for _, frameImg := range *packedFrameImgs {

        // Ignore non-tile frame images
        if frameImg.MetaData.FrameImageType != PropertyFrameImage {
            continue
        }

        originVData[frameImg.MetaData.Variation][frameImg.MetaData.Type][frameImg.MetaData.Animation] = Frame{
            X: frameImg.X,
            Y: frameImg.Y,
            Height: frameImg.Height,
        }
    }

    return &originVData
}

// Attach extra visual data stored away in JSON files
func attachExtraPropsVData(vData *PropertiesData) {
    // propsDir := baseDirPath + inputsDirName + propertiesDirName
    // attachJSONData(propsDir + palettesFileName, &vData.Palettes)
    // attachJSONData(propsDir +propsLightsRGBFileName, &vData.PropsLightsRGB)
}
