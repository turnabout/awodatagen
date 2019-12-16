package awossgen

// Data for all property tiles, attached to game data
type PropertiesData [][][]Frame

// Property type enum
type PropertyType uint8

const(
    HQ PropertyType = iota
    City
    Base
    Airport
    Port
)

func (p PropertyType) String() string {
    return [...]string{
        "HQ",
        "City",
        "Base",
        "Airport",
        "Port",
    }[p]
}

const FirstPropertyType = HQ
const LastPropertyType = Port
const PropertyTypeAmount = LastPropertyType + 1

// Weather enum
type Weather uint8

const(
    Clear Weather = iota
    Snow
    Rain
)

const WeatherFirst = Clear
const WeatherLast  = Rain
const WeatherCount = WeatherLast + 1

const PropWeatherVarFirst = Clear
const PropWeatherVarLast  = Snow
const PropWeatherVarCount = PropWeatherVarLast + 1

func (w Weather) String() string {
    return [...]string{
        "Clear",
        "Snow",
        "Rain",
    }[w]
}
