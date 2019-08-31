// Generates units' sprite sheet & visual data
package main

import (
    "image"
    "image/draw"
    "io/ioutil"
    "log"
    "os"
)

// Callback function for loopUnitsImgData
type UnitImgDataLoopFunc func(unitKey int, varKey int, animKey int)

// Data on every single unit image
// Dimension 1: Unit Type
// Dimension 2: Variation
// Dimension 3: Animation
// Dimension 4: Animation Frames
var unitsImgData[][][][]FrameImage

// Origin frame data (raw sprite sheet)
// Dimensions: Same as unitsImgData
var unitsOriginVisualData[][][][]Frame

// Destination frame data (final, in-game used sprite sheet)
// Dimension 1: Unit Type
// Dimension 2: Animation
// Dimension 3: Animation Frames
var unitsDestVisualData[][][]Frame

// Units' sprite sheet image
var unitsSSImg *image.RGBA

func init() {
    // Initialize visual data slices
    unitsImgData = make([][][][]FrameImage, UnitTypeAmount)
    unitsOriginVisualData = make([][][][]Frame, UnitTypeAmount)
    unitsDestVisualData = make([][][]Frame, UnitTypeAmount)
}

// Generate units' sprite sheet & visuals data
func generateUnits() {
    gatherUnitsImgData()

    // Process origin visual data & sprite sheet
    rowsData, unitsSSHeight := gatherRowsData(false)

    unitsSSImg = image.NewRGBA(image.Rectangle{
        Min: image.Point{X: 0, Y: 0},
        Max: image.Point{X: unitsSSWidth, Y: unitsSSHeight},
    })

    processOutput(rowsData, false)

    // Process destination visual data
    destRowsData, unitsFullSSHeight := gatherRowsData(true)
    processOutput(destRowsData, true)

    // Attach results to global visual data
    visualData.Units = UnitsData{
        Origin: unitsOriginVisualData,
        Dest: unitsDestVisualData,
        X: 0,
        Y: 0,
        Width: unitsSSWidth,
        Height: unitsSSHeight,
        FullWidth: unitsSSWidth,
        FullHeight: unitsFullSSHeight,
    }
}

// Gather data on every row of images in the sprite sheet (rows Height/frames Amount).
// processDest specifies if we're gathering destination visual data. Otherwise, origin data will be gathered.
// Returns full Height of all rows put together.
func gatherRowsData(processDest bool) (*[]RowData, int) {
    var rowsData[]RowData
    var rowWidth, rowHeight, rowFramesAmount, rowY int // Current row values

    // Loop every animation in previously gathered unit image data
    loopCb := func(unitKey int, varKey int, animKey int) {

        // Loop every animation frame
        for frameIndex := range unitsImgData[unitKey][varKey][animKey] {
            frame := unitsImgData[unitKey][varKey][animKey][frameIndex]

            // Check if row complete, store & reset row values if it is
            if rowWidth+ frame.Width > unitsSSWidth {
                rowsData = append(rowsData, RowData{Height: rowHeight, Amount: rowFramesAmount, Y: rowY})
                rowY += rowHeight
                rowWidth, rowHeight, rowFramesAmount = 0, 0, 0
            }

            // Update current row values
            rowFramesAmount++
            rowWidth += frame.Width

            if frame.Height > rowHeight {
                rowHeight = frame.Height
            }
        }
    }

    loopStoredUnitAnimations(loopCb, processDest)

    // If gathering destination rows data, include additional animations for each unit
    if processDest {
        for unitKey := range unitsImgData {
            loopCb(unitKey, int(OS), int(Right)) // "Left" animation
            loopCb(unitKey, int(OS), int(Idle))  // "Done" animation
        }
    }

    rowsData = append(rowsData, RowData{Height: rowHeight, Amount: rowFramesAmount, Y: rowY})
    return  &rowsData, rowY + rowHeight
}

// Process units' output
// processDest == false: Populate JSON origin visual data & draw sprite sheet
// processDest == true: Populate JSON destination visual data
func processOutput(rowsData *[]RowData, processDest bool) {
    var currentFrameIndex int // Index of next frame to add in current row
    var currentFrameX int     // Index of next frame to add's X
    var currentRowIndex int   // Index of the current row we're processing

    loopCb := func(unitKey int, varKey int, animKey int) {
        // Add this unit's variation array if it doesn't already exist (origin only)
        if !processDest && len(unitsOriginVisualData[unitKey]) < varKey + 1 {
            unitsOriginVisualData[unitKey] = append(unitsOriginVisualData[unitKey], [][]Frame{})
        }

        // Add this unit's animation array
        if processDest {
            unitsDestVisualData[unitKey] = append(unitsDestVisualData[unitKey], []Frame{})
        } else {
            unitsOriginVisualData[unitKey][varKey] = append(unitsOriginVisualData[unitKey][varKey], []Frame{})
        }

        // Loop every animation frame
        for frameIndex := range unitsImgData[unitKey][varKey][animKey] {

            // Jump to next row if we've reached the end
            if currentFrameIndex >= (*rowsData)[currentRowIndex].Amount {
                currentRowIndex++
                currentFrameX = 0
                currentFrameIndex = 0
            }

            frame := unitsImgData[unitKey][varKey][animKey][frameIndex]
            rect := frame.Image.Bounds()

            // Get difference in Height between this frame & this row
            rowHeightDiff := (*rowsData)[currentRowIndex].Height - frame.Height

            // Move image to the its X/Y coordinates
            rect.Min.X += currentFrameX
            rect.Max.X += currentFrameX

            rect.Min.Y += (*rowsData)[currentRowIndex].Y + rowHeightDiff
            rect.Max.Y += (*rowsData)[currentRowIndex].Y + rowHeightDiff

            // Draw image on sprite sheet (origin only)
            if !processDest {
                draw.Draw(unitsSSImg, rect, frame.Image, image.Point{X: 0, Y: 0}, draw.Src)
            }

            resultFrame := Frame{X: currentFrameX, Y: rect.Min.Y, Width: frame.Width, Height: frame.Height}

            // Store resulting visual data
            if processDest {
                unitsDestVisualData[unitKey][animKey] = append(unitsDestVisualData[unitKey][animKey], resultFrame)
            } else {
                unitsOriginVisualData[unitKey][varKey][animKey] = append(
                    unitsOriginVisualData[unitKey][varKey][animKey],
                    resultFrame,
                )
            }

            // Update current row values
            currentFrameIndex++
            currentFrameX += frame.Width
        }
    }

    loopStoredUnitAnimations(loopCb, processDest)

    if !processDest {
        return
    }

    // Destination-only loop callback for processing additional animations
    destLoopCb := func(unitKey int, ogAnimKey int, destAnimKey int) {

        // Add this unit's animation array
        unitsDestVisualData[unitKey] = append(unitsDestVisualData[unitKey], []Frame{})

        for frameIndex := range unitsImgData[unitKey][int(OS)][ogAnimKey] {
            if currentFrameIndex >= (*rowsData)[currentRowIndex].Amount {
                currentRowIndex++
                currentFrameX = 0
                currentFrameIndex = 0
            }

            frame := unitsImgData[unitKey][int(OS)][ogAnimKey][frameIndex]

            // Get difference in Height between this frame & this row
            rowHeightDiff := (*rowsData)[currentRowIndex].Height - frame.Height

            // Store resulting visual data
            unitsDestVisualData[unitKey][destAnimKey] = append(
                unitsDestVisualData[unitKey][destAnimKey],
                Frame{
                    X: currentFrameX,
                    Y: (*rowsData)[currentRowIndex].Y + rowHeightDiff,
                    Width: frame.Width,
                    Height: frame.Height,
                },
            )
        }
    }

    // Process additional destination data (additional animations for each unit)
    for unitKey := range unitsImgData {
        destLoopCb(unitKey, int(Right), int(Left)) // "Left" animation
        destLoopCb(unitKey, int(Idle), int(Done))  // "Done" animation
    }
}

// Loop every unit animation stored in 'unitsImgData'
// singleVar == true: Only loop one variation per unit
func loopStoredUnitAnimations(callback UnitImgDataLoopFunc, singleVar bool) {
    for unitKey := range unitsImgData {
        for varKey := range unitsImgData[unitKey] {
            if singleVar && varKey > int(OS) {
                continue
            }

            for animKey := range unitsImgData[unitKey][varKey] {
                callback(unitKey, varKey, animKey)
            }
        }
    }
}

// Gathers data on every single image, filling out "unitsImgData"
func gatherUnitsImgData() {
    // Get path of base directory containing unit images
    var unitsDirPath string = baseDirPath + imageInputsDirName + unitsDirName + "/"

    // Loop Units
    for unitType := FirstUnitType; unitType <= LastUnitType; unitType++ {
        unitDirPath := unitsDirPath + unitType.String() + "/"

        // Loop Variations of this Unit
        for unitVar := FirstUnitVariation; unitVar <= LastUnitVariation; unitVar++ {
            varDirPath := unitDirPath + unitVar.String() + "/"

            // Ignore this variation if it does not exist on this unit
            if _, err := os.Stat(varDirPath); os.IsNotExist(err) {
                break
            }

            // Add array for this variation to unitsImgData
            unitsImgData[unitType] = append(unitsImgData[unitType], [][]FrameImage{})

            // Loop Animations of this Variation
            for anim := FirstUnitAnimation; anim <= LastUnitAnimation; anim++ {
                unitsImgData[unitType][unitVar] = append(unitsImgData[unitType][unitVar], []FrameImage{})
                gatherAnimationData(unitType, unitVar, anim, varDirPath + anim.String() + "/")
            }
        }
    }
}

// Process a unit animation's directory, gathering all of its images' data for unitsImgData
func gatherAnimationData(unitKey UnitType, varKey UnitVariation, animKey UnitAnimation, animDir string) {
    // Loop imageNames
    files, err := ioutil.ReadDir(animDir)

    if err != nil {
        log.Fatal(err)
    }

    // Gather every image's data
    for _, file := range files {
        imageObj := getImage(animDir + file.Name())

        unitsImgData[unitKey][varKey][animKey] = append(
            unitsImgData[unitKey][varKey][animKey],
            FrameImage{
                Image:  imageObj,
                Width:  imageObj.Bounds().Max.X,
                Height: imageObj.Bounds().Max.Y,
            },
        )
    }
}
