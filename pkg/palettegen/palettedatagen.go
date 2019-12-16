package palettegen

import (
    "github.com/turnabout/awodatagen"
    "github.com/turnabout/awodatagen/pkg/genio"
)

func AttachPaletteData(vData *awodatagen.GameData) {
    var basePalettes map[string]awodatagen.Palette
    var rawPalettes []awodatagen.Palette

    genio.AttachJSONData( awodatagen.GetInputPath(awodatagen.AdditionalDir, awodatagen.BasePalettesFileName), &basePalettes )
    genio.AttachJSONData( awodatagen.GetInputPath(awodatagen.AdditionalDir, awodatagen.PalettesFileName), &rawPalettes )

    // Generate final palette data
    // Unit palettes
    var baseUnitPalette awodatagen.Palette = basePalettes["units"]
    var baseUnitDonePalette awodatagen.Palette = basePalettes["unitsDone"]

    for i := awodatagen.ArmyTypeFirst; i <= awodatagen.ArmyTypeLast; i++ {
        var unitPalette awodatagen.Palette = rawPalettes[i * 2]
        var unitDonePalette awodatagen.Palette = rawPalettes[(i * 2) + 1]

        vData.Palettes = append(vData.Palettes, *makePalette(&baseUnitPalette, &unitPalette))
        vData.Palettes = append(vData.Palettes, *makePalette(&baseUnitDonePalette, &unitDonePalette))
    }

    // Tile palettes
    var baseTilePalette awodatagen.Palette = basePalettes["tiles"]
    var baseTileFogPalette awodatagen.Palette = basePalettes["tilesFog"]

    var tilePalettesStart int = int(awodatagen.ArmyTypeCount) * 2

    for i := awodatagen.WeatherFirst; i <= awodatagen.WeatherLast; i++ {
        var tilePalette awodatagen.Palette = rawPalettes[int(tilePalettesStart) + (int(i) * 2)]
        var tileFogPalette awodatagen.Palette = rawPalettes[int(tilePalettesStart) + (int(i) * 2) + 1]

        vData.Palettes = append(vData.Palettes, *makePalette(&baseTilePalette, &tilePalette))
        vData.Palettes = append(vData.Palettes, *makePalette(&baseTileFogPalette, &tileFogPalette))
    }

    // Property palettes
    var propertyPalettesStart int = tilePalettesStart + (int(awodatagen.WeatherCount) * 2)
    var basePropertyPalette awodatagen.Palette = basePalettes["properties"]

    // + 2 for fogged/neutral properties palette
    for i := awodatagen.ArmyTypeFirst; i <= awodatagen.ArmyTypeLast+ 2; i++ {
        var propPalette awodatagen.Palette = rawPalettes[propertyPalettesStart + int(i)]

        vData.Palettes = append(vData.Palettes, *makePalette(&basePropertyPalette, &propPalette))
    }
}

// Create palette struct using a base and a main Palette raw data
func makePalette(basePalette *awodatagen.Palette, mainPalette *awodatagen.Palette) *awodatagen.Palette {
    var resPalette awodatagen.Palette = make(awodatagen.Palette)

    // Apply base & main palettes on resulting palette
    for key, val := range *basePalette {
        resPalette[key] = val
    }

    for key, val := range *mainPalette {
        resPalette[key] = val
    }

    return &resPalette
}
