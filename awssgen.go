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
)


func main() {
    visualData := VisualData{SSMetaData: ssMetaData{}}

    // Process Units
    _, unitsData := generateUnitsData(0, 0)
    visualData.Units = unitsData

    // Process Tiles
    tilesSS, tilesData := generateTilesData(unitsData.Width, 0)
    visualData.Tiles = *tilesData

    // Put all sprite sheets together to make the entire raw sprite sheet TODO
    outputSpriteSheet(tilesSS)
    outputJSON(&visualData)
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
