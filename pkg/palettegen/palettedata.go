package palettegen

// Data for all palettes, attached to game data
type PaletteData []Palette

// Generic palette
type Palette map[string]RGB

// Array representing an RGB pixel value
type RGB [3]int
