package tilegen

import "github.com/turnabout/awossgen"

// Map for looking up a Tile Type using its corresponding full string
var tileReverseStrings = map[string]awossgen.TileType {

    // Basic neutral tiles, represented visually
    "Plain": awossgen.Plain,
    "Forest": awossgen.Forest,
    "Mountain": awossgen.Mountain,
    "Road": awossgen.Road,
    "Bridge": awossgen.Bridge,
    "River": awossgen.River,
    "Sea": awossgen.Sea,
    "Reef": awossgen.Reef,
    "Shore": awossgen.Shore,
    "Pipe": awossgen.Pipe,
    "PipeFragile": awossgen.PipeFragile,
    "Silo": awossgen.Silo,

    // Additional neutral tiles, represented visually (standard size)
    "BaseSmoke": awossgen.BaseSmoke,
    "Empty": awossgen.Empty,

    // Additional neutral tiles, represented visually (non-standard size)
    "LandPiece": awossgen.LandPiece,
}

// Map for looking up a Tile Variation using its corresponding full string
var tileVarsReverseStrings = map[string]awossgen.TileVariation {
    "Default": awossgen.Default,
    "Horizontal": awossgen.Horizontal,
    "Vertical": awossgen.Vertical,
    "VerticalEnd": awossgen.VerticalEnd,
    "Top": awossgen.Top,
    "Bottom": awossgen.Bottom,
    "DirLeft": awossgen.DirLeft,
    "DirRight": awossgen.DirRight,
    "TopLeft": awossgen.TopLeft,
    "TopRight": awossgen.TopRight,
    "BottomLeft": awossgen.BottomLeft,
    "BottomRight": awossgen.BottomRight,
    "Middle": awossgen.Middle,
    "ShadowedDefault": awossgen.ShadowedDefault,
    "ShadowedTopLeft": awossgen.ShadowedTopLeft,
    "ShadowedBottomLeft": awossgen.ShadowedBottomLeft,
    "ShadowedLeft": awossgen.ShadowedLeft,
    "ShadowedHorizontal": awossgen.ShadowedHorizontal,
    "ShadowedVertical": awossgen.ShadowedVertical,
    "ShadowedVerticalEnd": awossgen.ShadowedVerticalEnd,
    "ShadowedTLeft": awossgen.ShadowedTLeft,
    "TTop": awossgen.TTop,
    "TBottom": awossgen.TBottom,
    "TLeft": awossgen.TLeft,
    "TRight": awossgen.TRight,
    "Small": awossgen.Small,
    "WaterfallUp": awossgen.WaterfallUp,
    "WaterfallDown": awossgen.WaterfallDown,
    "WaterfallLeft": awossgen.WaterfallLeft,
    "WaterfallRight": awossgen.WaterfallRight,
    "Hole": awossgen.Hole,
    "HoleHorizontal": awossgen.HoleHorizontal,
    "HoleVertical": awossgen.HoleVertical,
    "HoleLeft": awossgen.HoleLeft,
    "HoleRight": awossgen.HoleRight,
    "HoleTop": awossgen.HoleTop,
    "HoleBottom": awossgen.HoleBottom,
    "TopConnectedLeft": awossgen.TopConnectedLeft,
    "TopConnectedRight": awossgen.TopConnectedRight,
    "TopConnectedFull": awossgen.TopConnectedFull,
    "BottomConnectedLeft": awossgen.BottomConnectedLeft,
    "BottomConnectedRight": awossgen.BottomConnectedRight,
    "BottomConnectedFull": awossgen.BottomConnectedFull,
    "LeftConnectedTop": awossgen.LeftConnectedTop,
    "LeftConnectedBottom": awossgen.LeftConnectedBottom,
    "LeftConnectedFull": awossgen.LeftConnectedFull,
    "RightConnectedTop": awossgen.RightConnectedTop,
    "RightConnectedBottom": awossgen.RightConnectedBottom,
    "RightConnectedFull": awossgen.RightConnectedFull,
    "TopLeftConnectedVertical": awossgen.TopLeftConnectedVertical,
    "TopLeftConnectedHorizontal": awossgen.TopLeftConnectedHorizontal,
    "TopLeftConnectedFull": awossgen.TopLeftConnectedFull,
    "TopRightConnectedVertical": awossgen.TopRightConnectedVertical,
    "TopRightConnectedHorizontal": awossgen.TopRightConnectedHorizontal,
    "TopRightConnectedFull": awossgen.TopRightConnectedFull,
    "BottomLeftConnectedVertical": awossgen.BottomLeftConnectedVertical,
    "BottomLeftConnectedHorizontal": awossgen.BottomLeftConnectedHorizontal,
    "BottomLeftConnectedFull": awossgen.BottomLeftConnectedFull,
    "BottomRightConnectedVertical": awossgen.BottomRightConnectedVertical,
    "BottomRightConnectedHorizontal": awossgen.BottomRightConnectedHorizontal,
    "BottomRightConnectedFull": awossgen.BottomRightConnectedFull,
    "HorizontalClosed": awossgen.HorizontalClosed,
    "HorizontalOpen": awossgen.HorizontalOpen,
    "VerticalClosed": awossgen.VerticalClosed,
    "VerticalOpen": awossgen.VerticalOpen,
    "Used": awossgen.Used,
}
