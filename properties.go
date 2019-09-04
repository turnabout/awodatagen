package main

import (
    "os"
)

// Generate properties' sprite sheet & visual data
func generatePropertiesData() *PropertiesData {

    // Generate origin data
    originFrameImgs := gatherPropsFrameImages()
    packedFrameImgs, originWidth, originHeight := pack(originFrameImgs)

    // Generate destination data
    destFrameImgs := gatherPropsDestFrameImages(originFrameImgs)
    packedDestFrameImgs, destWidth, destHeight := pack(destFrameImgs)

    return &PropertiesData{
        Origin: *generatePropsOriginVData(packedFrameImgs),
        Dest: *generatePropsDestVData(packedDestFrameImgs),
        FogDest: nil,

        Width: originWidth,
        Height: originHeight,
        DestWidth: destWidth,
        DestHeight: destHeight,

        frameImg: FrameImage{
            Image: drawPackedFrames(packedFrameImgs, originWidth, originHeight),
            Width: originWidth,
            Height: originHeight,
            MetaData: FrameImageMetaData{Type: uint8(VisualDataProperties)},
        },
    }
}

// Gather Frame Images for Properties' origin
func gatherPropsFrameImages() *[]FrameImage {
    var frameImgs []FrameImage

    propsDir := baseDirPath + imageInputsDirName + propertiesDirName + "/"

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

                frameImgs = append(frameImgs, FrameImage{
                    Image: imageObj,
                    Width:  imageObj.Bounds().Max.X,
                    Height: imageObj.Bounds().Max.Y,
                    MetaData: FrameImageMetaData{
                        Type: uint8(propType),
                        Variation: uint8(weatherVar),
                        Animation: uint8(unitVar),
                    },
                })
            }
        }
    }

    return &frameImgs
}

// Gather Frame Images for Properties' destination
func gatherPropsDestFrameImages(originFrameImgs *[]FrameImage) *[]FrameImage {
    var frameImgs []FrameImage

    // Only keep the first weather variation & first unit variation
    for _, originFrameImg := range *originFrameImgs {
        if PropertyWeatherVariation(originFrameImg.MetaData.Variation) == FirstPropertyWeatherVariation &&
            UnitVariation(originFrameImg.MetaData.Animation) == FirstUnitVariation {
            frameImgs = append(frameImgs, originFrameImg)
            frameImgs = append(frameImgs, originFrameImg)
        }
    }

    return &frameImgs
}

// Generate the visual data for Properties' origin
func generatePropsOriginVData(packedFrameImgs *[]FrameImage) *[][][]Frame {

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

    // Fill out Origin visual data
    for _, frameImg := range *packedFrameImgs {
        originVData[frameImg.MetaData.Variation][frameImg.MetaData.Type][frameImg.MetaData.Animation] = Frame{
            X: frameImg.X,
            Y: frameImg.Y,
            Height: frameImg.Height,
        }
    }

    return &originVData
}

// Generate the visual data for Properties' destination
func generatePropsDestVData(packedFrameImgs *[]FrameImage) *[][][]Frame {

    // Property Type -> Unit Variation -> Animation Frames
    destVData := make([][][]Frame, PropertyTypeAmount)

    // Initialize property type arrays
    for propType := range destVData {

        var unitVarAmount int

        // HQ Properties have all Unit Variations, while other properties only have one
        if PropertyType(propType) == HQ {
            unitVarAmount = int(UnitVariationAmount)
        } else {
            unitVarAmount = 1
        }

        destVData[propType] = make([][]Frame, unitVarAmount)
    }

    // Fill out Destination visual data
    for _, frameImg := range *packedFrameImgs {
        destVData[frameImg.MetaData.Type][frameImg.MetaData.Animation] = append(
            destVData[frameImg.MetaData.Type][frameImg.MetaData.Animation],
            Frame{
                X: frameImg.X,
                Y: frameImg.Y,
                Height: frameImg.Height,
            })
    }

    return &destVData
}
