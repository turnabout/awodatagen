package main

import (
    "os"
)

// Generate properties' sprite sheet & visual data
func generatePropertiesData() *PropertiesData {

    // Get source frame images
    srcFrameImgs := getPropsSrcFrameImgs()
    packedSrcFrameImgs, srcWidth, srcHeight := pack(srcFrameImgs)

    // Get destination frame images
    dstFrameImgs := getPropsDstFrameImgs(srcFrameImgs)
    packedDstFrameImgs, dstWidth, dstHeight := pack(dstFrameImgs)

    // Get fog destination frame images
    fogFrameImgs := getPropsFogDstFrameImgs(srcFrameImgs)
    packedFogFrameImgs, fogWidth, fogHeight := pack(fogFrameImgs)

    vData := PropertiesData{
        Src:    *getPropsSrcVData(packedSrcFrameImgs),
        Dst:    *getPropsDstVData(packedDstFrameImgs),
        FogDst: *getPropsFogDstVData(packedFogFrameImgs),

        SrcWidth:     srcWidth,
        SrcHeight:    srcHeight,
        DstWidth:     dstWidth,
        DstHeight:    dstHeight,
        FogDstWidth:  fogWidth,
        FogDstHeight: fogHeight,

        frameImg: FrameImage{
            Image:    drawPackedFrames(packedSrcFrameImgs, srcWidth, srcHeight),
            Width:    srcWidth,
            Height:   srcHeight,
            MetaData: FrameImageMetaData{Type: uint8(VisualDataProperties)},
        },
    }

    attachExtraPropsVData(&vData)

    return &vData
}

// Gather Frame Images for Properties' source
func getPropsSrcFrameImgs() *[]FrameImage {
    var frameImgs []FrameImage

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
func getPropsDstFrameImgs(originFrameImgs *[]FrameImage) *[]FrameImage {
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

// Gather Frame Images for Properties' fog destination
func getPropsFogDstFrameImgs(originFrameImgs *[]FrameImage) *[]FrameImage {
    var frameImgs[]FrameImage

    // Only keep the first weather variation
    for _, originFrameImg := range *originFrameImgs {
        if PropertyWeatherVariation(originFrameImg.MetaData.Variation) == FirstPropertyWeatherVariation {
            frameImgs = append(frameImgs, originFrameImg)
            frameImgs = append(frameImgs, originFrameImg)
        }
    }

    return &frameImgs
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
        originVData[frameImg.MetaData.Variation][frameImg.MetaData.Type][frameImg.MetaData.Animation] = Frame{
            X: frameImg.X,
            Y: frameImg.Y,
            Height: frameImg.Height,
        }
    }

    return &originVData
}

// Generate the visual data for Properties' destination
func getPropsDstVData(packedFrameImgs *[]FrameImage) *[][]Frame {

    // Property Type -> Animation Frames
    destVData := make([][]Frame, PropertyTypeAmount)

    // Fill out Destination visual data
    for _, frameImg := range *packedFrameImgs {
        destVData[frameImg.MetaData.Type] = append(
            destVData[frameImg.MetaData.Type],
            Frame{
                X: frameImg.X,
                Y: frameImg.Y,
                Height: frameImg.Height,
            },
        )
    }

    return &destVData
}

// Generate the visual data for fog Properties' destination
func getPropsFogDstVData(packedFrameImgs *[]FrameImage) *[][]Frame {

    // Property Type -> Unit Variation
    destVData := make([][]Frame, PropertyTypeAmount)

    // Fill out Destination visual data
    for _, frameImg := range *packedFrameImgs {
        propType := PropertyType(frameImg.MetaData.Type)
        unitVar := UnitVariation(frameImg.MetaData.Animation)

        // Get amount of missing Frames up until the Frame we're processing
        missingFrames := (int(unitVar) + 1) - len(destVData[propType])

        // Add missing Frame(s)
        if missingFrames > 0 {
            for i := 0; i < missingFrames; i++ {
                destVData[propType] = append(destVData[propType], Frame{})
            }
        }

        // Record this frame
        destVData[propType][unitVar] = Frame{
            X: frameImg.X,
            Y: frameImg.Y,
            Height: frameImg.Height,
        }
    }

    return &destVData
}

// Attach extra visual data stored away in JSON files
func attachExtraPropsVData(vData *PropertiesData) {
    basePropsDir := baseDirPath + inputsDirName + propertiesDirName

    attachJSONData(basePropsDir + palettesFileName, &vData.Palettes)
    attachJSONData(basePropsDir + propsLightsOnColor, &vData.PropsLightsOnColor)
}
