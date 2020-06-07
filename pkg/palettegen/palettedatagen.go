package palettegen

import (
	"github.com/turnabout/awodatagen/internal/config"
	"github.com/turnabout/awodatagen/internal/genio"
	"github.com/turnabout/awodatagen/pkg/propertygen"
	"github.com/turnabout/awodatagen/pkg/unitgen"
)

func GetPaletteData() *PaletteData {
	var basePalettes map[string]Palette
	var rawPalettes []Palette
	var data PaletteData

	genio.AttachJSONData(genio.GetInputPath(config.OtherDir, config.BasePalettesFileName), &basePalettes)
	genio.AttachJSONData(genio.GetInputPath(config.OtherDir, config.PalettesFileName), &rawPalettes)

	// Generate final palette data
	// Unit palettes
	var baseUnitPalette Palette = basePalettes["units"]
	var baseUnitDonePalette Palette = basePalettes["unitsDone"]

	for i := unitgen.ArmyTypeFirst; i <= unitgen.ArmyTypeLast; i++ {
		var unitPalette Palette = rawPalettes[i*2]
		var unitDonePalette Palette = rawPalettes[(i*2)+1]

		data = append(data, *makePalette(&baseUnitPalette, &unitPalette))
		data = append(data, *makePalette(&baseUnitDonePalette, &unitDonePalette))
	}

	// Tile palettes
	var baseTilePalette Palette = basePalettes["tiles"]
	var baseTileFogPalette Palette = basePalettes["tilesFog"]

	var tilePalettesStart int = int(unitgen.ArmyTypeCount) * 2

	for i := propertygen.WeatherFirst; i <= propertygen.WeatherLast; i++ {
		var tilePalette Palette = rawPalettes[int(tilePalettesStart)+(int(i)*2)]
		var tileFogPalette Palette = rawPalettes[int(tilePalettesStart)+(int(i)*2)+1]

		data = append(data, *makePalette(&baseTilePalette, &tilePalette))
		data = append(data, *makePalette(&baseTileFogPalette, &tileFogPalette))
	}

	// Property palettes
	var propertyPalettesStart int = tilePalettesStart + (int(propertygen.WeatherCount) * 2)
	var basePropertyPalette Palette = basePalettes["properties"]

	// + 2 for fogged/neutral properties palette
	for i := unitgen.ArmyTypeFirst; i <= unitgen.ArmyTypeLast+2; i++ {
		var propPalette Palette = rawPalettes[propertyPalettesStart+int(i)]

		data = append(data, *makePalette(&basePropertyPalette, &propPalette))
	}

	return &data
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
