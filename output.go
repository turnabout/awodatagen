// Functions used to process outputs
package main

import (
    "encoding/json"
    "fmt"
    "image"
    "image/png"
    "io/ioutil"
    "log"
    "os"
    "regexp"
)

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

// Attach the JSON data at the given file path and stores the result in the value pointed to by v
func attachJSONData(jsonPath string , v interface{}) {
    data, err := ioutil.ReadFile(jsonPath)

    if err != nil {
        log.Fatal(err)
    }

    // Make Regexp used to remove comments
    re := regexp.MustCompile(`//.*`)

    // Unmarshal and store the result
    if err := json.Unmarshal(re.ReplaceAll(data, []byte("")), v); err != nil {
        log.Fatal(err)
    }
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

// Gather additional visual data and attach to the main visual data object
func attachAdditionalVData(vData *VisualData) {
    addDir := baseDirPath + inputsDirName + additionalDirName

    attachJSONData(addDir + stagesFileName, &vData.Stages)
    attachJSONData(addDir +animClocksFileName, &vData.AnimationSubClocks)
    attachPaletteData(vData, addDir)
}

// Make up palette using a base and a main Palette raw data
func makePalette(basePalette *Palette, mainPalette *Palette) *Palette {
    var resPalette Palette = make(Palette)

    // Apply base & main palettes on resulting palette
    for key, val := range *basePalette {
        resPalette[key] = val
    }

    for key, val := range *mainPalette {
        resPalette[key] = val
    }

    return &resPalette
}

func attachPaletteData(vData *VisualData, addDir string) {
    var basePalettes map[string]Palette
    var rawPalettes []Palette

    attachJSONData(addDir + basePalettesFileName, &basePalettes)
    attachJSONData(addDir + palettesFileName, &rawPalettes)

    // Generate final palette data
    // Unit palettes
    var baseUnitPalette Palette = basePalettes["units"]
    var baseUnitDonePalette Palette = basePalettes["unitsDone"]

    for i := FirstUnitVariation; i <= LastUnitVariation; i++ {
        var unitPalette Palette = rawPalettes[i * 2]
        var unitDonePalette Palette = rawPalettes[(i * 2) + 1]

        vData.Palettes = append(vData.Palettes, *makePalette(&baseUnitPalette, &unitPalette))
        vData.Palettes = append(vData.Palettes, *makePalette(&baseUnitDonePalette, &unitDonePalette))
    }

    // Tile palettes
    var baseTilePalette Palette = basePalettes["tiles"]
    var baseTileFogPalette Palette = basePalettes["tilesFog"]

    var tilePalettesStart int = int(UnitVariationAmount) * 2

    for i := FirstWeather; i <= LastWeather; i++ {
        var tilePalette Palette = rawPalettes[int(tilePalettesStart) + (int(i) * 2)]
        var tileFogPalette Palette = rawPalettes[int(tilePalettesStart) + (int(i) * 2) + 1]

        vData.Palettes = append(vData.Palettes, *makePalette(&baseTilePalette, &tilePalette))
        vData.Palettes = append(vData.Palettes, *makePalette(&baseTileFogPalette, &tileFogPalette))
    }

    // Property palettes
    var propertyPalettesStart int = tilePalettesStart + (int(WeatherCount) * 2)
    var basePropertyPalette Palette = basePalettes["properties"]

    // + 2 for fogged/neutral properties palette
    for i := FirstUnitVariation; i <= LastUnitVariation + 2; i++ {
        var propPalette Palette = rawPalettes[propertyPalettesStart + int(i)]

        vData.Palettes = append(vData.Palettes, *makePalette(&basePropertyPalette, &propPalette))
    }
}

// Adjust the X/Y coordinates of units' src frames, adding units' sprite sheet X/Y position within the full sprite sheet
func adjustUnitsSrc(vData *VisualData) {
    for typeKey := range vData.Units.Src {
        for varKey := range vData.Units.Src[typeKey] {
            for animKey := range vData.Units.Src[typeKey][varKey] {
                for frameIndex := range vData.Units.Src[typeKey][varKey][animKey] {
                    vData.Units.Src[typeKey][varKey][animKey][frameIndex].X += vData.Units.srcX
                    vData.Units.Src[typeKey][varKey][animKey][frameIndex].Y += vData.Units.srcY
                }
            }
        }
    }
}
