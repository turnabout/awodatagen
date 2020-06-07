package autovargen

import (
	"github.com/turnabout/awodatagen/pkg/tilegen/tiledata"
	"github.com/turnabout/awodatagen/pkg/utilities"
	"unicode"
	"unicode/utf8"
)

// Symbol types
const (
	SymbolTileType = iota
	SymbolCompound
	SymbolANDNOT
	SymbolOR
	SymbolEmpty
	SymbolUnknown
)

// Processes an adjacent tile string into an adjacent tile bit field
func ProcessAdjTileStr(rawString string) uint {

	// Loop every symbol and use them to determine the bit field
	var symbolType int
	var symbolString string
	var nextStartIndex int = 0

	var resultBitField uint = 0
	var applyANDNOT bool = false

	for {
		symbolType, symbolString = getNextSymbol(rawString, nextStartIndex)
		nextStartIndex += utf8.RuneCountInString(symbolString)

		if symbolType == SymbolEmpty {
			break
		}

		if symbolType == SymbolUnknown {
			utilities.LogFatalF(
				"Tile autovar: unknown symbol type '%d' in symbol string '%s' from full raw string string '%s'",
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

		// If symbol is OR, just get to the next value
		if symbolType == SymbolOR {
			continue
		}

		// Get value to apply
		var appliedVal uint

		switch symbolType {
		case SymbolTileType:
			appliedVal = 1 << uint(tiledata.TileReverseStrings[symbolString])
			break
		case SymbolCompound:
			appliedVal = uint(autoVarCompoundVals[symbolString])
			break
		default:
			utilities.LogFatalF(
				"tilesAutoVar: Unknown symbol '%d'\n",
				symbolType,
			)
		}

		// Apply the value to the resulting bit field
		if applyANDNOT {
			applyANDNOT = false
			resultBitField &= ^appliedVal
		} else {
			resultBitField |= appliedVal
		}
	}

	return resultBitField
}

// Get the next symbol from a raw adjacent tile string
// Returns: the symbol type and the symbol's full string.
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
	if !unicode.IsLetter(rune(stringToProcess[0])) {

		if stringToProcess[0] == '&' && stringToProcess[1] == '~' {
			return SymbolANDNOT, "&~"
		} else if stringToProcess[0] == '|' {
			return SymbolOR, "|"
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

	// Loop characters until a non-letter is found, or until the rest of the string is processed
	for _, char := range stringToProcess {
		if !unicode.IsLetter(char) {
			break
		}

		resultCharCount++
	}

	if resultCharCount == 0 {
		return SymbolEmpty, ""
	}

	return symbolType, stringToProcess[0:resultCharCount]
}
