package uigen

import (
    "github.com/turnabout/awodatagen"
    "log"
)

func getUiElementByString(str string) awodatagen.UIElement {
    var ok bool
    var uiElement awodatagen.UIElement

    if uiElement, ok = awodatagen.UIElementsReverseStrings[str]; !ok {
        log.Fatalf("UI Element string '%s' not part of the UiElement enum\n", str)
    }

    return uiElement
}
