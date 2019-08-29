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

// Used to store a unit's frame's visual data within the game's sprite sheet
type UnitFrame struct {
    x, y, w, h int
}

// A unit's animation image data (image/width/height)
type UnitImg struct {
    img image.Image
    w int
    h int
}

// Data detailing a row of sprite images in a sprite sheet
type RowData struct {
    height int // Height in pixels
    amount int // Amount of images in the row
    y      int // Row's Y coordinate
}

// Data on every single unit image
// Dimension 1: Unit Type
// Dimension 2: Variation
// Dimension 3: Animation
// Dimension 4: Animation Frames
var unitsImgData[][][][]UnitImg

// Origin frame data (raw sprite sheet)
// Dimensions: Same as unitsImgData
var unitsOriginVisualData[][][][]UnitFrame

// Destination frame data (final, in-game used sprite sheet)
// Dimension 1: Unit Type
// Dimension 2: Animation
// Dimension 3: Animation Frames
var unitsDestVisualData[][][]UnitFrame

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
    unitsOriginVisualData = make([][][][]UnitFrame, len(unitTypes))
    unitsDestVisualData = make([][][]UnitFrame, len(unitTypes))
}

// Generate units' sprite sheet & visuals data
func generateUnits() {
    gatherUnitImgData()
    rowData, ssHeight := gatherRowsData()

    // Prepare sprite sheet image
    unitsSSImg = image.NewRGBA(image.Rectangle{
        Min: image.Point{X: 0, Y: 0},
        Max: image.Point{X: unitsSSWidth, Y: ssHeight},
    })

    processSpriteSheet(rowData)
}

// Gather data on every row of images in the sprite sheet (rows height/frames amount).
// Returns full height of all rows put together.
func gatherRowsData() (*[]RowData, int) {
    var rowsData[]RowData
    var rowWidth, rowHeight, rowFramesAmount, rowY int // Current row values

    // Loop every animation in previously gathered unit image data
    cb := func(unitKey int, varKey int, animKey int) {

        // Loop every animation frame
        for frameIndex := range unitsImgData[unitKey][varKey][animKey] {
            frame := unitsImgData[unitKey][varKey][animKey][frameIndex]

            // Check if row complete, store & reset row values if it is
            if rowWidth+ frame.w > unitsSSWidth {
                rowsData = append(rowsData, RowData{height: rowHeight, amount: rowFramesAmount, y: rowY})
                rowY += rowHeight
                rowWidth, rowHeight, rowFramesAmount = 0, 0, 0
            }

            // Update current row values
            rowFramesAmount++
            rowWidth += frame.w

            if frame.h > rowHeight {
                rowHeight = frame.h
            }
        }
    }

    loopStoredUnitAnimations(cb)
    rowsData = append(rowsData, RowData{height: rowHeight, amount: rowFramesAmount, y: rowY})

    return  &rowsData, rowY + rowHeight
}

// Process units' origin, gathering visual data & drawing the sprite sheet
func processSpriteSheet(rowsData *[]RowData) {
    var currentFrameIndex int // Index of next frame to add in current row
    var currentFrameX int     // Index of next frame to add's X
    var currentRowIndex int   // Index of the current row we're processing

    cb := func(unitKey int, varKey int, animKey int) {

        // Loop every animation frame
        for frameIndex := range unitsImgData[unitKey][varKey][animKey] {

            // Jump to next row if we've reached the end
            if currentFrameIndex >= (*rowsData)[currentRowIndex].amount {
                currentRowIndex++
                currentFrameX = 0
                currentFrameIndex = 0
            }

            frame := unitsImgData[unitKey][varKey][animKey][frameIndex]
            rect := frame.img.Bounds()

            // Get difference in height between this frame & this row
            rowHeightDiff := (*rowsData)[currentRowIndex].height - frame.h

            // Move image to the its x/y coordinates
            rect.Min.X += currentFrameX
            rect.Max.X += currentFrameX

            rect.Min.Y += (*rowsData)[currentRowIndex].y + rowHeightDiff
            rect.Max.Y += (*rowsData)[currentRowIndex].y + rowHeightDiff

            // Draw image on sprite sheet
            draw.Draw(unitsSSImg, rect, frame.img, image.Point{X: 0, Y: 0}, draw.Src)

            // Update current row values
            currentFrameIndex++
            currentFrameX += frame.w
        }
    }

    loopStoredUnitAnimations(cb)
}

// Loop every unit animation stored in 'unitsImgData'
func loopStoredUnitAnimations(callback UnitImgDataLoopFunc) {
    for unitKey := range unitsImgData {
        for varKey := range unitsImgData[unitKey] {
            for animKey := range unitsImgData[unitKey][varKey] {
                callback(unitKey, varKey, animKey)
            }
        }
    }
}

// Gathers data on every single image, filling out "unitsImgData"
func gatherUnitImgData() {
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
                img: imageObj,
                w: imageObj.Bounds().Max.X,
                h: imageObj.Bounds().Max.Y,
            },
        )
    }
}
