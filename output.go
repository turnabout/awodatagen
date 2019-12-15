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
    "path"
    "regexp"
    "runtime/debug"
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
        fmt.Printf("%s\n", jsonPath)
        debug.PrintStack()
        log.Fatal(err)
    }
}

// Output the visuals data JSON
func outputJSON(visualData *GameData) {
    // data, err := json.Marshal(visualData)
    data, err := json.MarshalIndent(visualData, "", "\t")

    if err != nil {
        log.Fatal(err)
    }

    // Use either the AWO JSON environment variable path or this directory as a default
    var jsonOutputPath string
    var envExists bool

    // If environment variable doesn't exist, output in this directory directly
    if jsonOutputPath, envExists = os.LookupEnv(jsonOutputEnvVar); !envExists {
        jsonOutputPath = path.Join(".", jsonOutputDefaultName)
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

    // If environment variable doesn't exist, output in this directory directly
    if ssOutputPath, envExists = os.LookupEnv(ssOutputEnvVar); !envExists {
        ssOutputPath = path.Join(".", ssOutputDefaultName)
    }

    writeImage(ssOutputPath, ss)
    fmt.Printf("Output %s\n", ssOutputPath)
}

// Gather additional visual data and attach to the main visual data object
func attachAdditionalVData(gameData *GameData) {

    attachJSONData( getFullProjectPath(additionalDir, stagesFileName), &gameData.Stages )
    attachJSONData( getFullProjectPath(additionalDir, animClocksFileName), &gameData.AnimationClocks)

    attachPaletteData(gameData)
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

func attachPaletteData(vData *GameData) {
    var basePalettes map[string]Palette
    var rawPalettes []Palette

    attachJSONData( getFullProjectPath(additionalDir, basePalettesFileName), &basePalettes )
    attachJSONData( getFullProjectPath(additionalDir, palettesFileName), &rawPalettes )

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
