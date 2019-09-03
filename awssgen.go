// Generates sprite sheets & data used by AWO (outputs visuals.json & spritesheet.png)
package main

import (
    "encoding/json"
    "fmt"
    "image"
    "image/png"
    _ "image/png"
    "io/ioutil"
    "log"
    "os"
    "sort"
)


func main() {
    var vData = VisualData{
        Units: generateUnitsData(),
        Tiles: generateTilesData(),
        SSMetaData: ssMetaData{},
    }

    outputSpriteSheet(joinSpriteSheets(&vData))
    outputJSON(&vData)
}

// Join all sprite sheets together, update their metadata in the Visual Data and return the final, raw sprite sheet
func joinSpriteSheets(vData *VisualData) *image.RGBA {
    var packedFrames *[]FrameImage

    // Pack all the sprite sheets together to make one
    packedFrames, vData.SSMetaData.Width, vData.SSMetaData.Height = pack(&[]FrameImage{
        vData.Units.frameImg,
        vData.Tiles.frameImg,
    })

    // Update sprite sheet meta data on each visual data object, after sorting the packed frames
    sort.Sort(SizeSorter(*packedFrames))

    vData.Units.X = (*packedFrames)[VisualDataUnits].X
    vData.Units.Y = (*packedFrames)[VisualDataUnits].Y
    vData.Tiles.X = (*packedFrames)[VisualDataTiles].X
    vData.Tiles.Y = (*packedFrames)[VisualDataTiles].Y

    // Return the final sprite sheet
    return drawPackedFrames(packedFrames, vData.SSMetaData.Width, vData.SSMetaData.Height)
}

// Output the visuals data JSON
func outputJSON(visualData *VisualData) {
    // data, err := json.Marshal(visualData)
    data, err := json.MarshalIndent(visualData, "", "\t")

    if err != nil {
        log.Fatal(err)
    }

    // Use either the AWO JSON environment variable path or this directory as a default
    var jsonOutputPath string
    var envExists bool

    if jsonOutputPath, envExists = os.LookupEnv(jsonOutputEnvVar); !envExists {
        // Environment variable doesn't exist, output in this directory directly
        jsonOutputPath = baseDirPath + "/" + jsonOutputDefaultName
    }

    err = ioutil.WriteFile(jsonOutputPath, data, 0644)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Output %s\n", jsonOutputPath)
}

// Output the game sprite sheet
func outputSpriteSheet(ss *image.RGBA) {
    // Use either the AWO sprite sheet environment variable path or this directory as a default
    var ssOutputPath string
    var envExists bool

    if ssOutputPath, envExists = os.LookupEnv(ssOutputEnvVar); !envExists {
        // Environment variable doesn't exist, output in this directory directly
        ssOutputPath = baseDirPath + "/" + ssOutputDefaultName
    }

    writeImage(ssOutputPath, ss)
    fmt.Printf("Output %s\n", ssOutputPath)
}

// Gets the image stored at the given path
func getImage(path string) image.Image {
    // Load Image file
    imgFile, err := os.Open(path)

    if err != nil {
        log.Fatal(err)
    }

    // Decode Image
    img, _, err := image.Decode(imgFile)

    if err != nil {
        log.Fatal(err)
    }

    return img
}

// Write a given image to the given path
func writeImage(path string, outputImg image.Image) {
    out, err := os.Create(path)

    if err != nil {
        log.Fatal(err)
    }

    if png.Encode(out, outputImg) != nil {
        log.Fatal(err)
    }
}
