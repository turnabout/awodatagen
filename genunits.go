// Generates units' sprite sheet & visual data
package main

import (
    "fmt"
    "image"
    "io/ioutil"
    "log"
    "os"
)

// Unit types "enum"
type unitType uint8

const (
    Infantry unitType = iota
    Mech
    Recon
    Tank
    MdTank
    NeoTank
    APC
    Artillery
    Rockets
    Missiles
    AntiAir
    Battleship
    Cruiser
    Lander
    Sub
    Fighter
    Bomber
    BattleCopter
    TransportCopter
)

// Unit variations "enum"
type unitVariation uint8

const (
    OS unitVariation = iota
    BM
    GE
    YC
    BH
)

// Unit animations "enum"
type unitAnimation uint8

const (
    Idle unitAnimation = iota
    Right
    Up
    Down
)

// Used to store a unit's frame's visual data within the game's sprite sheet
type unitFrame struct {
    x, y, w, h int
}

// Used to store a unit's animation image data (image/width/height)
type unitImg struct {
    img image.Image
    w int
    h int
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

    fmt.Printf("%d %d\n", unitsImgData[int(Infantry)][int(OS)][int(Idle)][0].w, unitsImgData[int(Infantry)][int(OS)][int(Idle)][0].h)
    fmt.Printf("%d %d\n", unitsImgData[int(Infantry)][int(OS)][int(Idle)][1].w, unitsImgData[int(Infantry)][int(OS)][int(Idle)][1].h)
    fmt.Printf("%d %d\n", unitsImgData[int(Infantry)][int(OS)][int(Idle)][2].w, unitsImgData[int(Infantry)][int(OS)][int(Idle)][2].h)
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
