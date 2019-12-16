package palettegen

import "github.com/turnabout/awossgen"

// Make up palette using a base and a main Palette raw data
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

func AttachPaletteData(vData *awossgen.GameData) {
    var basePalettes map[string]awossgen.Palette
    var rawPalettes []awossgen.Palette

    // TODO
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
