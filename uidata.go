package awossgen

// Data for all UI elements, attached to game data
type UIData [][]Frame

// UI element enum
// TODO: rename -> UIElement
type UiElement uint8

const(
    TileCursor UiElement = iota
    StarSm
    StarLg

    UiElementNone = 255
)

const UiElementFirst = TileCursor
const UiElementLast = StarLg
const UiElementCount = UiElementLast + 1

// Map for looking up a Ui Element using its corresponding full string
var uiElementsReverseStrings = map[string]UiElement {
    "TileCursor": TileCursor,
    "StarSm": StarSm,
    "StarLg": StarLg,
}
