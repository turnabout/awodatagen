package uigen

import (
    "github.com/turnabout/awossgen"
    "log"
)

func getUiElementByString(str string) awossgen.UIElement {
    var ok bool
    var uiElement awossgen.UIElement

    if uiElement, ok = awossgen.UIElementsReverseStrings[str]; !ok {
        log.Fatalf("UI Element string '%s' not part of the UiElement enum\n", str)
    }

    return uiElement
}
