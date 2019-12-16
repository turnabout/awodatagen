package palettegen

import (
    "github.com/turnabout/awossgen"
    "github.com/turnabout/awossgen/pkg/genio"
)

func AttachPaletteData(vData *awossgen.GameData) {
    var basePalettes map[string]awossgen.Palette
    var rawPalettes []awossgen.Palette

    genio.AttachJSONData( awossgen.GetInputPath(awossgen.AdditionalDir, awossgen.BasePalettesFileName), &basePalettes )
    genio.AttachJSONData( awossgen.GetInputPath(awossgen.AdditionalDir, awossgen.PalettesFileName), &rawPalettes )

    // Generate final palette data
    // Unit palettes
    var baseUnitPalette awossgen.Palette = basePalettes["units"]
    var baseUnitDonePalette awossgen.Palette = basePalettes["unitsDone"]

    for i := awossgen.ArmyTypeFirst; i <= awossgen.ArmyTypeLast; i++ {
        var unitPalette awossgen.Palette = rawPalettes[i * 2]
        var unitDonePalette awossgen.Palette = rawPalettes[(i * 2) + 1]

        vData.Palettes = append(vData.Palettes, *makePalette(&baseUnitPalette, &unitPalette))
        vData.Palettes = append(vData.Palettes, *makePalette(&baseUnitDonePalette, &unitDonePalette))
    }

    // Tile palettes
    var baseTilePalette awossgen.Palette = basePalettes["tiles"]
    var baseTileFogPalette awossgen.Palette = basePalettes["tilesFog"]

    var tilePalettesStart int = int(awossgen.ArmyTypeCount) * 2

    for i := awossgen.WeatherFirst; i <= awossgen.WeatherLast; i++ {
        var tilePalette awossgen.Palette = rawPalettes[int(tilePalettesStart) + (int(i) * 2)]
        var tileFogPalette awossgen.Palette = rawPalettes[int(tilePalettesStart) + (int(i) * 2) + 1]

        vData.Palettes = append(vData.Palettes, *makePalette(&baseTilePalette, &tilePalette))
        vData.Palettes = append(vData.Palettes, *makePalette(&baseTileFogPalette, &tileFogPalette))
    }

    // Property palettes
    var propertyPalettesStart int = tilePalettesStart + (int(awossgen.WeatherCount) * 2)
    var basePropertyPalette awossgen.Palette = basePalettes["properties"]

    // + 2 for fogged/neutral properties palette
    for i := awossgen.ArmyTypeFirst; i <= awossgen.ArmyTypeLast+ 2; i++ {
        var propPalette awossgen.Palette = rawPalettes[propertyPalettesStart + int(i)]

        vData.Palettes = append(vData.Palettes, *makePalette(&basePropertyPalette, &propPalette))
    }
}

// Create palette struct using a base and a main Palette raw data
func makePalette(basePalette *awossgen.Palette, mainPalette *awossgen.Palette) *awossgen.Palette {
    var resPalette awossgen.Palette = make(awossgen.Palette)

    // Apply base & main palettes on resulting palette
    for key, val := range *basePalette {
        resPalette[key] = val
    }

    for key, val := range *mainPalette {
        resPalette[key] = val
    }

    return &resPalette
}
