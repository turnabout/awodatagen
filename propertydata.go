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
    WeatherClear Weather = iota
    WeatherSnow
    WeatherRain
)

const WeatherFirst = WeatherClear
const WeatherLast  = WeatherRain
const WeatherCount = WeatherLast + 1

const FirstPropertyWeatherVariation = WeatherClear
const LastPropertyWeatherVariation = WeatherSnow
const PropertyWeatherVariationAmount = LastPropertyWeatherVariation + 1
