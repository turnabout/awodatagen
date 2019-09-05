// Generates units' sprite sheet & visual data
package main

import (
    "io/ioutil"
    "log"
    "os"
)

// Generate units' sprite sheet & visuals data.
// SrcX/SrcY specifies the coordinates of the units' sprite sheet within the final raw sprite sheet
func getUnitsData()  *UnitsData {

    // Get source frame images (for the raw sprite sheet)
    srcFrameImgs := getUnitsSrcFrameImgs()
    packedSrcFrameImgs, srcWidth, srcHeight := pack(srcFrameImgs)

    // Get destination frame images (for sprite sheets generated in-game)
    dstFrameImgs := getUnitsDstFrameImgs(srcFrameImgs)
    packedDstFrameImgs, dstWidth, dstHeight := pack(dstFrameImgs)

    return &UnitsData{
        Src:       *getUnitsSrcVData(packedSrcFrameImgs),
        Dst:       *getUnitsDstVData(packedDstFrameImgs),
        SrcWidth:  srcWidth,
        SrcHeight: srcHeight,
        DstWidth:  dstWidth,
        DstHeight: dstHeight,
        frameImg: FrameImage{
            Image:    drawPackedFrames(packedSrcFrameImgs, srcWidth, srcHeight),
            Width:    srcWidth,
            Height:   srcHeight,
            MetaData: FrameImageMetaData{Type: uint8(VisualDataUnits)},
        },
    }
}

// Gets Frame Images from every single Unit image
func getUnitsSrcFrameImgs() *[]FrameImage {
    var frameImgs []FrameImage

    unitsDir := baseDirPath + imageInputsDirName + unitsDirName + "/"

    // Loop Units
    for unitType := FirstUnitType; unitType <= LastUnitType; unitType++ {
        unitDir := unitsDir + unitType.String() + "/"

        // Loop Variations
        for unitVar := FirstUnitVariation; unitVar <= LastUnitVariation; unitVar++ {
            varDir := unitDir + unitVar.String() + "/"

            // Ignore this variation if it does not exist on this unit
            if _, err := os.Stat(varDir); os.IsNotExist(err) {
                break
            }

            // Loop Animations
            for anim := FirstUnitAnimation; anim <= LastUnitAnimation; anim++ {
                getAnimFrameImgs(unitType, unitVar, anim, varDir + anim.String() + "/", &frameImgs)
            }
        }
    }

    return &frameImgs
}

// Get all Frame Images from the given Unit Animation
func getAnimFrameImgs(uType UnitType, uVar UnitVariation, uAnim UnitAnimation, animDir string, frameImgs *[]FrameImage) {
    imgFiles, err := ioutil.ReadDir(animDir)

    if err != nil {
        log.Fatal(err)
    }

    // Loop every image of this Animation
    for index, imgFile := range imgFiles {
        imageObj := getImage(animDir + imgFile.Name())

        *frameImgs = append(*frameImgs, FrameImage{
            Image: imageObj,
            Width:  imageObj.Bounds().Max.X,
            Height: imageObj.Bounds().Max.Y,
            MetaData: FrameImageMetaData{
                Type: uint8(uType),
                Variation: uint8(uVar),
                Animation: uint8(uAnim),
                Index: index,
            },
        })
    }
}

// Generate the origin visual data (units' visual data on the raw sprite sheet) using packed Frame Images
func getUnitsSrcVData(packedFrameImgs *[]FrameImage) *[][][][]Frame {

    // Unit Type -> Variation -> Animation -> Animation Frames
    originVData := make([][][][]Frame, UnitTypeAmount)

    for _, frameImg := range *packedFrameImgs {
        unitType := frameImg.MetaData.Type
        unitVar := frameImg.MetaData.Variation
        unitAnim := frameImg.MetaData.Animation
        unitFrame := frameImg.MetaData.Index

        // Check if Variation is missing, add up to it if necessary
        missingVars := int(unitVar + 1) - len(originVData[unitType])

        if missingVars > 0 {
            for i := 0; i < missingVars; i++ {
                originVData[unitType] = append(originVData[unitType], [][]Frame{})
            }
        }

        // Check if Animation is missing, add up to it if necessary
        missingAnims := int(unitAnim + 1) - len(originVData[unitType][unitVar])

        if missingAnims > 0 {
            for i := 0; i < missingAnims; i++ {
                originVData[unitType][unitVar] = append(originVData[unitType][unitVar], []Frame{})
            }
        }

        // Check if Animation Frame is missing, add up to it if necessary
        missingFrames := int(unitFrame + 1) - len(originVData[unitType][unitVar][unitAnim])

        if missingFrames > 0 {
            for i := 0; i < missingFrames; i++ {
                originVData[unitType][unitVar][unitAnim] = append(originVData[unitType][unitVar][unitAnim], Frame{})
            }
        }

        // Store data
        originVData[unitType][unitVar][unitAnim][unitFrame] = Frame{
            X: frameImg.X,
            Y: frameImg.Y,
            Width: frameImg.Width,
            Height: frameImg.Height,
        }
    }

    return &originVData
}

// Generate the destination visual data (visual data for final, in-game sprite sheet generated for each army) using
// packed Frame Images
func getUnitsDstVData(packedFrameImgs *[]FrameImage) *[][][]Frame {

    // Unit Type -> Animation -> Animation Frame
    destVData := make([][][]Frame, UnitTypeAmount)

    // Initialize all Animation slices
    for unitType := range destVData {
        destVData[unitType] = make([][]Frame, UnitAnimationFullAmount)
    }

    // Process every Frame Image, storing them into destination visual data
    for _, frameImg := range *packedFrameImgs {
        unitType := UnitType(frameImg.MetaData.Type)
        unitAnim := UnitAnimation(frameImg.MetaData.Animation)
        unitFrame := frameImg.MetaData.Index

        // Get amount of missing Frames up until the Frame we're processing
        missingFrames := (unitFrame + 1) - len(destVData[unitType][unitAnim])

        // Check if this Frame Image belongs to an Animation that is mirrored and was already previously stored.
        // If it was, it means we need to process a mirrored (extra) animation.
        // Update the Animation to be the one it's mirroring, and get the amount of missing frames that Animation.
        if (unitAnim == Idle || unitAnim == Right) && missingFrames < 1 && destVData[unitType][unitAnim][unitFrame].Width > 0 {
            unitAnim += UnitExtraAnimationConvert
            missingFrames = (unitFrame + 1) - len(destVData[unitType][unitAnim])
        }

        // Add missing Frame(s)
        if missingFrames > 0 {
            for i := 0; i < missingFrames; i++ {
                destVData[unitType][unitAnim] = append(destVData[unitType][unitAnim], Frame{})
            }
        }

        // Store data
        destVData[unitType][unitAnim][unitFrame] = Frame{
            X: frameImg.X,
            Y: frameImg.Y,
            Width: frameImg.Width,
            Height: frameImg.Height,
        }
    }

    return &destVData
}

// Take Frame Images and prepare them for destination visual data processing, removing Frame Images for extra variations
// and adding Frame Images for extra Animations
func getUnitsDstFrameImgs(frameImgs *[]FrameImage) *[]FrameImage {
    var resFrameImgs []FrameImage

    // Filter out Frame Images belonging to Variations other than the first
    for _, frameImg := range *frameImgs {
        unitAnim := UnitAnimation(frameImg.MetaData.Animation)

        // Ignore Unit Variations other than the first
        if UnitVariation(frameImg.MetaData.Variation) > FirstUnitVariation {
            continue
        }

        resFrameImgs = append(resFrameImgs, frameImg)

        // If this Frame Image belongs to an Animation that is mirrored, add it a second time
        if unitAnim == Idle || unitAnim == Right {
            resFrameImgs = append(resFrameImgs, frameImg)
        }
    }

    return &resFrameImgs
}
