package uigen

import (
    "github.com/turnabout/awodatagen"
    "github.com/turnabout/awodatagen/pkg/packer"
)

// Generates the UI-related game data
func GetUIData(packedFrameImgs *[]packer.FrameImage) *awodatagen.UIData {
    var UIData *awodatagen.UIData = getUIBaseData(packedFrameImgs)

    return UIData
}

func getUIBaseData(packedFrameImgs *[]packer.FrameImage) *awodatagen.UIData {

    // UI Element Type -> UI Element Frames
    UIData := make(awodatagen.UIData, awodatagen.UIElementCount)

    // Process frame images
    for _, frameImg := range *packedFrameImgs {

        // Ignore non-UI element frame images
        if frameImg.MetaData.FrameImageDataType != uint8(awodatagen.UIDataType) {
            continue
        }

        // Get metadata on the UI element this frame image represents
        uiElement := awodatagen.UIElement(frameImg.MetaData.Type)
        uiElFrame := frameImg.MetaData.Index

        // Add any frames missing up until the one we're adding
        if missingFrames := (uiElFrame + 1) - len(UIData[uiElement]); missingFrames > 0 {
            for i := 0; i < missingFrames; i++ {
                UIData[uiElement] = append(UIData[uiElement], awodatagen.Frame{})
            }
        }

        // Add the Frame data to the animation slice, and record it to the visual data
        frame := awodatagen.Frame{
            X: frameImg.X,
            Y: frameImg.Y,
            Width: frameImg.Width,
            Height: frameImg.Height,
        }

        UIData[uiElement][uiElFrame] = frame
    }

    return &UIData
}

