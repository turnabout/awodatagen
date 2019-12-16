package awodatagen

// Data for all UI elements, attached to game data
type UIData [][]Frame

// UI element enum
type UIElement uint8

const(
    TileCursor UIElement = iota
    StarSm
    StarLg

    UIElementNone = 255
)

const UIElementFirst = TileCursor
const UIElementLast = StarLg
const UIElementCount = UIElementLast + 1

// Map for looking up a Ui Element using its corresponding full string
var UIElementsReverseStrings = map[string]UIElement {
    "TileCursor": TileCursor,
    "StarSm": StarSm,
    "StarLg": StarLg,
}