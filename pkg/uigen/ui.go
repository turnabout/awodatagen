package uigen

import (
    "log"
)

func getUiElementByString(str string) UIElement {
    var ok bool
    var uiElement UIElement

    if uiElement, ok = UIElementsReverseStrings[str]; !ok {
        log.Fatalf("UI Element string '%s' not part of the UiElement enum\n", str)
    }

    return uiElement
}
