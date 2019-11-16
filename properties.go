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

    // Loop property types
    for propType := FirstPropertyType; propType <= LastPropertyType; propType++ {
        typeDir := propsDir + propType.String() + "/"

        // Loop variations (weather)
        for weatherVar := FirstPropertyWeatherVariation; weatherVar <= LastPropertyWeatherVariation; weatherVar++ {
            weatherDir := typeDir + weatherVar.String() + "/"

            // Loop unit variations (armies)
            for unitVar := FirstUnitVariation; unitVar <= LastUnitVariation; unitVar++ {
                fullPath := weatherDir + unitVar.String() + ".png"

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

    // Property type -> Variation (weather) -> Unit variation (army)
    originVData := make([][][]Frame, PropertyTypeAmount)

    // Initialize property type arrays
    for propertyType := range originVData {

        // Get amount of unit variations (armies) for this property type (HQs have all, others have only one)
        var unitVarAmount int

        if PropertyType(propertyType) == HQ {
            unitVarAmount = int(UnitVariationAmount)
        } else {
            unitVarAmount = 1
        }

        originVData[propertyType] = make([][]Frame, PropertyWeatherVariationAmount)

        // Initialize variation (weather) arrays
        for weatherVar := range originVData[propertyType] {

            originVData[propertyType][weatherVar] = make([]Frame, unitVarAmount)
        }
    }

    // Fill out properties source data
    for _, frameImg := range *packedFrameImgs {

        // Ignore non-tile frame images
        if frameImg.MetaData.FrameImageType != PropertyFrameImage {
            continue
        }

        propertyType := PropertyType(frameImg.MetaData.Type)
        weatherVariation := Weather(frameImg.MetaData.Variation)
        army := UnitType(frameImg.MetaData.Animation)

        originVData[propertyType][weatherVariation][army] = Frame{
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
