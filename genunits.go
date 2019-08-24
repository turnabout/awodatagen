// Generates units' sprite sheet & visual data
package main

import (
    "image"
    "io/ioutil"
    "log"
    "os"
)

// Callback function for loopUnitsImgData
type UnitImgDataLoopFunc func(unitKey int, varKey int, animKey int)

// Used to store a unit's frame's visual data within the game's sprite sheet
type unitFrame struct {
    x, y, w, h int
}

// A unit's animation image data (image/width/height)
type unitImg struct {
    img image.Image
    w int
    h int
}

// Data detailing a row of sprite images in a sprite sheet
type rowData struct {
    height int // Height in pixels
    amount int // Amount of images in the row
}

// Data on every single unit image
// Index 1: Unit Type
// Index 2: Variation
// Index 3: Animation
var unitsImgData[][][][]unitImg

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
    // Initialize unitsImgData slice
    unitsImgData = make([][][][]unitImg, len(unitTypes))
}

// Generate units' spritesheet & visuals data
func generateUnits() {
    gatherUnitImgData()
    gatherRowsData()

    // fmt.Printf("%d %d\n", unitsImgData[int(Infantry)][int(OS)][int(Idle)][0].w, unitsImgData[int(Infantry)][int(OS)][int(Idle)][0].h)
    // fmt.Printf("%d %d\n", unitsImgData[int(Infantry)][int(OS)][int(Idle)][1].w, unitsImgData[int(Infantry)][int(OS)][int(Idle)][1].h)
    // fmt.Printf("%d %d\n", unitsImgData[int(Infantry)][int(OS)][int(Idle)][2].w, unitsImgData[int(Infantry)][int(OS)][int(Idle)][2].h)
}

// Gather data on every row of images in the sprite sheet
func gatherRowsData() {
    var rows[]rowData
    var rowWidth, rowHeight, rowFramesAmount int

    // Loop every animation in previously gathered unit image data
    cb := func(unitKey int, varKey int, animKey int) {

        // Loop every animation frame
        for frameIndex := range unitsImgData[unitKey][varKey][animKey] {
            frame := unitsImgData[unitKey][varKey][animKey][frameIndex]

            // Check if row complete, store & reset row values if it is
            if rowWidth+ frame.w > unitsSSWidth {
                rows = append(rows, rowData{height: rowHeight, amount: rowFramesAmount})
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
            unitsImgData[unitKey] = append(unitsImgData[unitKey], [][]unitImg{})

            // Loop every variation animation
            for animKey := range sortedAnimKeys {
                // Add array for this animation to unitsImgData
                unitsImgData[unitKey][varKey] = append(unitsImgData[unitKey][varKey], []unitImg{})

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
            unitImg {
                img: imageObj,
                w: imageObj.Bounds().Max.X,
                h: imageObj.Bounds().Max.Y,
            },
        )
    }
}
