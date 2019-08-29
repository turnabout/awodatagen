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

// A unit's animation image data (image/width/Height)
type UnitImg struct {
    Image  image.Image
    Width  int
    Height int
}

// Data detailing a row of sprite images in a sprite sheet
type RowData struct {
    Height int // Height in pixels
    Amount int // Amount of images in the row
    Y      int // Row's Y coordinate
}

// Data on every single unit image
// Dimension 1: Unit Type
// Dimension 2: Variation
// Dimension 3: Animation
// Dimension 4: Animation Frames
var unitsImgData[][][][]UnitImg

// Origin frame data (raw sprite sheet)
// Dimensions: Same as unitsImgData
var unitsOriginVisualData[][][][]Frame

// Destination frame data (final, in-game used sprite sheet)
// Dimension 1: Unit Type
// Dimension 2: Animation
// Dimension 3: Animation Frames
var unitsDestVisualData[][][]Frame

// Sprite sheet image
var unitsSSImg *image.RGBA

// Every unit type names as indexes, and their corresponding numbered values
var unitTypes = map[int]string {
    int(Infantry):        "Infantry",
    int(Mech):            "Mech",
    int(Recon):           "Recon",
    int(Tank):            "Tank",
    int(MdTank):          "MdTank",
    int(NeoTank):         "NeoTank",
    int(APC):             "APC",
    int(Artillery):       "Artillery",
    int(Rockets):         "Rockets",
    int(Missiles):        "Missiles",
    int(AntiAir):         "AntiAir",
    int(Battleship):      "Battleship",
    int(Cruiser):         "Cruiser",
    int(Lander):          "Lander",
    int(Sub):             "Sub",
    int(Fighter):         "Fighter",
    int(Bomber):          "Bomber",
    int(BattleCopter):    "BattleCopter",
    int(TransportCopter): "TransportCopter",
}

// Every unit variation names as indexes, and their corresponding numbered values
var unitVariations = map[int]string {
    int(OS): "OS",
    int(BM): "BM",
    int(GE): "GE",
    int(YC): "YC",
    int(BH): "BH",
}

// Every unit animation names as indexes, and their corresponding numbered values
var unitAnimations = map[int]string {
    int(Idle):  "Idle",
    int(Right): "Right",
    int(Up):    "Up",
    int(Down):  "Down",
}

func init() {
    // Initialize visual data slices
    unitsImgData = make([][][][]UnitImg, len(unitTypes))
    unitsOriginVisualData = make([][][][]Frame, len(unitTypes))
    unitsDestVisualData = make([][][]Frame, len(unitTypes))
}

// Generate units' sprite sheet & visuals data
func generateUnits() {
    gatherUnitsImgData()

    // Process origin visual data & sprite sheet
    rowsData, ssHeight := gatherRowsData(false)

    unitsSSImg = image.NewRGBA(image.Rectangle{
        Min: image.Point{X: 0, Y: 0},
        Max: image.Point{X: unitsSSWidth, Y: ssHeight},
    })

    processOutput(rowsData, false)

    // Process destination visual data
    destRowsData, _ := gatherRowsData(true)
    processOutput(destRowsData, true)
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

    // Get sorted keys for every unit map
    sortedUnitKeys := getMapSortedKeys(unitTypes)
    sortedVarKeys := getMapSortedKeys(unitVariations)
    sortedAnimKeys := getMapSortedKeys(unitAnimations)

    // Loop every unit in order of values from unitTypes
    for unitKey := range sortedUnitKeys {
        var unitDirPath string = unitsDirPath + unitTypes[unitKey] + "/"

        // Loop every unit variation in order of values from unitVariations
        for varKey := range sortedVarKeys {
            var varDirPath string = unitDirPath + unitVariations[varKey] + "/"

            // Ignore this variation if it does not exist on this unit
            if _, err := os.Stat(varDirPath); os.IsNotExist(err) {
                break
            }

            // Add array for this variation to unitsImgData
            unitsImgData[unitKey] = append(unitsImgData[unitKey], [][]UnitImg{})

            // Loop every variation animation
            for animKey := range sortedAnimKeys {
                // Add array for this animation to unitsImgData
                unitsImgData[unitKey][varKey] = append(unitsImgData[unitKey][varKey], []UnitImg{})

                // Gather data from this animation's images
                gatherAnimationData(unitKey, varKey, animKey, varDirPath + unitAnimations[animKey] + "/")
            }
        }
    }
}

// Process a unit animation's directory, gathering all of its images' data for unitsImgData
func gatherAnimationData(unitKey int, varKey int, animKey int, animDir string) {
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
            UnitImg{
                Image:  imageObj,
                Width:  imageObj.Bounds().Max.X,
                Height: imageObj.Bounds().Max.Y,
            },
        )
    }
}
