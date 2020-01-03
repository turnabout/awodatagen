package placementrulesgen

type rawTilePlacementRules map[string][][]rawTilePlacementRule

// Rule prohibiting tile placement in design room mode (raw - must be transformed to a `awodatagen.TilePlacementRule`)
type rawTilePlacementRule struct {
    Position string `json:"position"` // The position of the tile this placement rule applies to
    Tiles    string `json:"tiles"`    // Bit field describing the tiles the placement rule applies to
}
