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

// Property weather variation enum
type PropertyWeatherVariation uint8

const(
    Clear PropertyWeatherVariation = iota
    Snow
)

func (p PropertyWeatherVariation) String() string {
    return [...]string{
        "Clear",
        "Snow",
    }[p]
}

const FirstPropertyWeatherVariation = Clear
const LastPropertyWeatherVariation = Snow
const PropertyWeatherVariationAmount = Snow + 1

// Weather enum
type Weather uint8

const(
    WeatherClear Weather = iota
    WeatherSnow
    WeatherRain
)

const FirstWeather = WeatherClear
const LastWeather = WeatherRain
const WeatherCount = LastWeather + 1
