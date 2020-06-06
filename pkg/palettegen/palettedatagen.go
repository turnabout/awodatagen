package palettegen

// TODO: Return palette data instead of attaching
/*
import (
    "github.com/turnabout/awodatagen"
    "github.com/turnabout/awodatagen/pkg/gamedata"
    "github.com/turnabout/awodatagen/pkg/genio"
	"github.com/turnabout/awodatagen/pkg/propertygen"
    "github.com/turnabout/awodatagen/pkg/unitgen"
)

func AttachPaletteData(vData *gamedata.GameData) {
    var basePalettes map[string]Palette
    var rawPalettes []Palette

    genio.AttachJSONData( awodatagen.GetInputPath(awodatagen.OtherDir, awodatagen.BasePalettesFileName), &basePalettes )
    genio.AttachJSONData( awodatagen.GetInputPath(awodatagen.OtherDir, awodatagen.PalettesFileName), &rawPalettes )

    // Generate final palette data
    // Unit palettes
    var baseUnitPalette Palette = basePalettes["units"]
    var baseUnitDonePalette Palette = basePalettes["unitsDone"]

    for i := unitgen.ArmyTypeFirst; i <= unitgen.ArmyTypeLast; i++ {
        var unitPalette Palette = rawPalettes[i * 2]
        var unitDonePalette Palette = rawPalettes[(i * 2) + 1]

        vData.Palettes = append(vData.Palettes, *makePalette(&baseUnitPalette, &unitPalette))
        vData.Palettes = append(vData.Palettes, *makePalette(&baseUnitDonePalette, &unitDonePalette))
    }

    // Tile palettes
    var baseTilePalette Palette = basePalettes["tiles"]
    var baseTileFogPalette Palette = basePalettes["tilesFog"]

    var tilePalettesStart int = int(unitgen.ArmyTypeCount) * 2

    for i := propertygen.WeatherFirst; i <= propertygen.WeatherLast; i++ {
        var tilePalette Palette = rawPalettes[int(tilePalettesStart) + (int(i) * 2)]
        var tileFogPalette Palette = rawPalettes[int(tilePalettesStart) + (int(i) * 2) + 1]

        vData.Palettes = append(vData.Palettes, *makePalette(&baseTilePalette, &tilePalette))
        vData.Palettes = append(vData.Palettes, *makePalette(&baseTileFogPalette, &tileFogPalette))
    }

    // Property palettes
    var propertyPalettesStart int = tilePalettesStart + (int(propertygen.WeatherCount) * 2)
    var basePropertyPalette Palette = basePalettes["properties"]

    // + 2 for fogged/neutral properties palette
    for i := unitgen.ArmyTypeFirst; i <= unitgen.ArmyTypeLast+ 2; i++ {
        var propPalette Palette = rawPalettes[propertyPalettesStart + int(i)]

        vData.Palettes = append(vData.Palettes, *makePalette(&basePropertyPalette, &propPalette))
    }
}

// Create palette struct using a base and a main Palette raw data
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
 */
