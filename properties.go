package main

import (
    "os"
)

// Generate properties' sprite sheet & visual data
func generatePropertiesData() *PropertiesData {

    // Generate origin data
    frameImgs := gatherPropsFrameImages()
    packedFrameImgs, originWidth, originHeight := pack(frameImgs)

    return &PropertiesData{
        Origin: *generatePropsOriginVData(packedFrameImgs),

        Width: originWidth,
        Height: originHeight,

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