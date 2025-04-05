package models

type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type RouteInfo struct {
	Distance    float64      `json:"distance"`     // в метрах
	Duration    float64      `json:"duration"`     // в секундах
	Coordinates []Coordinate `json:"coordinates"`  // точки маршрута
}

type Weather struct {
	Temperature float64 `json:"temperature"` // температура в градусах Цельсия
	FeelsLike   float64 `json:"feelsLike"`   // ощущаемая температура
	Condition   string  `json:"condition"`   // описание погодных условий
	WindSpeed   float64 `json:"windSpeed"`   // скорость ветра в м/с
	Humidity    float64 `json:"humidity"`    // влажность в процентах
}

type MapSettings struct {
	Center      Coordinate `json:"center"`
	Zoom        int        `json:"zoom"`
	MapType     string     `json:"mapType"`     // "map", "satellite", "hybrid"
	ShowTraffic bool       `json:"showTraffic"` // показывать ли пробки
}

type TrackPoint struct {
	Coordinate Coordinate `json:"coordinate"`
	Timestamp  int64      `json:"timestamp"`  // Unix timestamp
	Elevation  float64    `json:"elevation"`  // высота над уровнем моря
	Speed      float64    `json:"speed"`      // скорость в м/с
	Accuracy   float64    `json:"accuracy"`   // точность GPS в метрах
} 