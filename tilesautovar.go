package main

import (
    "log"
    "unicode"
    "unicode/utf8"
)

// Take multiple tile types and compose a bit field for usage in auto var data
func composeTileTypeBitField(values []TileType) uint {
    var result uint = 0

    for _, val := range values {
        result |= (1 << val)
    }

    return result
}

// Values corresponding to auto var compound symbols
var compoundAutoVarValues = map[string]uint{
    "any": 0xFFF,
    "shadowing": composeTileTypeBitField([]TileType{Forest, Mountain, Silo}),
    "oob": composeTileTypeBitField([]TileType{OOB}),
    "land": composeTileTypeBitField([]TileType{Plain, Forest, Mountain, Road, Bridge, Pipe, PipeFragile, Silo}),
}

// Attach auto var data to accumulated tiles data
func attachTilesAutoVarData(tilesDir string, vData *TilesData) {
    var rawData RawAutoVarsData

    // Load raw auto var data file into structure
    attachJSONData(tilesDir + tilesAutoVarFileName, &rawData)

    // Loop every tile type
    for tileTypeStr, tileTypeAutoVars := range rawData {
        var tileType TileType = tileReverseStrings[tileTypeStr]

        // TODO: remove temporary debug condition
        if tileType != Forest && tileType != Plain && tileType != Bridge {
            continue
        }

        // fmt.Printf("%s\n", tileType)

        // Add initial slice for the tile type
        vData.Src[tileType].AutoVars = []AutoVarData{}

        // Loop auto var values, appending every one of them to this tile type's AutoVars field
        for _, autoVarData := range tileTypeAutoVars {
            vData.Src[tileType].AutoVars = append(
                vData.Src[tileType].AutoVars,
                processRawAutoVar(autoVarData),
            )
        }

        // Order the auto var data for this tile type, placing data with the least amount of active bits in its adjacent
        // tiles first.
        // TODO
    }
}

// Process the adjacent tiles in a raw autovar data struct and produce a final exported struct, containing a short
// string version of the tile variation and 4 bit field numbers representing the acceptable adjacent tiles for this var.
func processRawAutoVar(rawAutoVarData RawAutoVarData) AutoVarData {

    var tileVar TileVariation = tileVarsReverseStrings[rawAutoVarData.TileVar]

    result := AutoVarData{
        TileVar: tileVar.String(),
        AdjacentTiles: [4]int{0, 0, 0, 0},
    }

    // fmt.Printf("%s\n", tileVar.String())

    // Process every adjacent tile string into a bit field number representing acceptable tile types
    for i := 0; i < ADJACENT_TILE_AMOUNT; i++ {
        result.AdjacentTiles[i] = translateAdjTileStr(rawAutoVarData.AdjacentTiles[i])
        // fmt.Printf("%s(%s)\n", rawAutoVarData.AdjacentTiles[i], strconv.FormatUint(uint64(result.AdjacentTiles[i]), 2))
    }

    // fmt.Printf("\n\n")
    // fmt.Printf("%#v\n\n", result)

    return result
}

// Symbol types
const (
    SymbolTileType = iota
    SymbolCompound
    SymbolANDNOT
    SymbolEmpty
    SymbolUnknown
)

// Translate an adjacent tile string into a bir field number.
func translateAdjTileStr(rawString string) int {

    // Loop every symbol and use them to determine the bit field
    var symbolType int
    var symbolString string
    var nextStartIndex int = 0

    var resultBitField int = 0
    var applyANDNOT bool = false

    for {
        symbolType, symbolString = getNextSymbol(rawString, nextStartIndex)
        nextStartIndex += utf8.RuneCountInString(symbolString)

        if symbolType == SymbolEmpty {
            break
        }

        if symbolType == SymbolUnknown {
            log.Fatalf(
                "Tiles autovar: unknown symbol type '%d' in symbol string '%s' from full raw string string '%s'",
                symbolType,
                symbolString,
                rawString,
            )
        }

        // If symbol signifies ANDNOT, store that the next operation should use it
        if symbolType == SymbolANDNOT {
            applyANDNOT = true
            continue
        }

        // Get value to apply
        var appliedVal int

        switch symbolType {
        case SymbolTileType:
            appliedVal = 1 << uint(tileReverseStrings[symbolString])
            break
        case SymbolCompound:
            appliedVal = int(compoundAutoVarValues[symbolString])
            break
        default:
            log.Fatalf("Unknown symbol type")
            break
        }

        // Apply the value to the resulting bit field
        if applyANDNOT {
            applyANDNOT = false
            resultBitField &= ^appliedVal
        } else {
            resultBitField |= appliedVal
        }

        // fmt.Printf("%s(%d)__", symbolString, symbolType)
    }

    // fmt.Printf("\n")

    return resultBitField
}

// Get the next symbol from a raw adjacent tile string
// Returns: the symbol type and the symbol's full string.
// unicode.IsSymbol()
func getNextSymbol(rawString string, startIndex int) (int, string) {
    var stringToProcess string
    var resultCharCount int = 0
    var symbolType int

    // Nothing left to process, empty symbol
    if startIndex >= utf8.RuneCountInString(rawString) {
        return SymbolEmpty, ""
    }

    stringToProcess = string(rawString[startIndex:])

    // Determine symbol type
    // If first character is a symbol, we can easily determine the symbol type
    if stringToProcess[0] == '&' {
        if stringToProcess[1] == '~' {
            return SymbolANDNOT, "&~"
        }

        return SymbolUnknown, ""
    }

    // If first character is not a symbol, the symbol type should either be a tile type or a compound type
    if unicode.IsUpper(rune(stringToProcess[0])) {
        symbolType = SymbolTileType
    } else if unicode.IsLower(rune(stringToProcess[0])) {
        symbolType = SymbolCompound
    } else {
        return SymbolUnknown, ""
    }

    // Loop characters until a symbol is found, or until the rest of the string is processed
    for _, char := range stringToProcess {
        if unicode.IsSymbol(char) {
            resultCharCount--
            break
        }

        resultCharCount++
    }

    return symbolType, stringToProcess[0:resultCharCount]
}
