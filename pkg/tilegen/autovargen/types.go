package autovargen

// Raw auto var data directly loaded from input json
type rawAutoVarsData map[string][]rawAutoVarData

type rawAutoVarData struct {
    TileVar string                         // The tile variation
    AdjacentTiles [4]string `json:"Tiles"` // Strings describing the adjacent tiles that correspond to the tile
                                           // variation. In order, the adjacent tiles go: up, right, down, left
}

// Auto var data's adjacent tile indexes
// TODO: lowercase
const AUTOVAR_ADJACENT_TILE_UP    = 0
const AUTOVAR_ADJACENT_TILE_RIGHT = 1
const AUTOVAR_ADJACENT_TILE_DOWN  = 2
const AUTOVAR_ADJACENT_TILE_LEFT  = 3
const ADJACENT_TILE_COUNT         = 4
