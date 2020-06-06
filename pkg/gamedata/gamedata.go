package gamedata

import (
	"github.com/turnabout/awodatagen/pkg/cogen"
	"github.com/turnabout/awodatagen/pkg/palettegen"
	"github.com/turnabout/awodatagen/pkg/propertygen"
	"github.com/turnabout/awodatagen/pkg/tilegen/tiledata"
	"github.com/turnabout/awodatagen/pkg/uigen"
	"github.com/turnabout/awodatagen/pkg/unitgen"
)

// Game data JSON structure
type GameData struct {
	Units                 unitgen.UnitData         `json:"units"`
	Tiles                 tiledata.TileData        `json:"tiles"`
	Properties            propertygen.PropertyData `json:"properties"`
	UI                    uigen.UIData             `json:"ui"`
	Palettes              palettegen.PaletteData   `json:"palettes"`
	Stages                StageData                `json:"stages"`
	COs                   cogen.COData             `json:"COs"`
	Clocks                []AnimationClock         `json:"clocks"`
	SpriteSheetDimensions SSDimensions             `json:"ssDimensions"`
}

// Animation Clock structure
type AnimationClock struct {
	Frames []int `json:"frames"` // Frame counts that make the animation clock tick
	Values []int `json:"values"` // Values emitted by the animation clock when it ticks
}

// Sprite sheet dimensions
type SSDimensions struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}
