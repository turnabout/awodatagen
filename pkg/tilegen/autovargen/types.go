package autovargen

// Raw auto var data directly loaded from input json
type rawAutoVarsData map[string][]rawAutoVarData

type rawAutoVarData struct {
    TileVar string                         // The tile variation
    AdjacentTiles [4]string `json:"Tiles"` // Strings describing the adjacent tiles that correspond to the tile
                                           // variation. In order, the adjacent tiles go: up, right, down, left
}

// Auto var data's adjacent tile indexes
type adjacentTileIndex int

const(
    adjacentTileUp adjacentTileIndex = iota
    adjacentTileRight
    adjacentTileDown
    adjacentTileLeft
)

const adjacentTileCount = 4
