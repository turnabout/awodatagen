package placementrulesgen

type rawTilePlacementRules map[string][]rawTilePlacementRule
type rawTilePlacementRule []rawTilePlacementRuleComponent

// Component of a rule prohibiting tile placement in design room mode (raw, must be transformed)
type rawTilePlacementRuleComponent struct {
    Position string `json:"position"` // The position of the tile this placement rule applies to
    Tiles    string `json:"tiles"`    // Bit field describing the tiles the placement rule applies to
}
