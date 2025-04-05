package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"push_run_app/backend/models"
)

type MapService struct {
	apiKey string
}

func NewMapService(apiKey string) *MapService {
	return &MapService{
		apiKey: apiKey,
	}
}

// GetRouteInfo получает информацию о маршруте между точками
func (s *MapService) GetRouteInfo(coordinates []models.Coordinate) (*models.RouteInfo, error) {
	if len(coordinates) < 2 {
		return nil, fmt.Errorf("at least 2 coordinates are required")
	}

	// Формируем URL для запроса к API Яндекс Маршрутизации
	url := fmt.Sprintf("https://api.routing.yandex.net/v2/route?apikey=%s", s.apiKey)
	
	// Формируем точки маршрута
	var waypoints string
	for i, coord := range coordinates {
		if i > 0 {
			waypoints += "~"
		}
		waypoints += fmt.Sprintf("%f,%f", coord.Latitude, coord.Longitude)
	}
	url += "&waypoints=" + waypoints

	// Выполняем запрос
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get route info: %v", err)
	}
	defer resp.Body.Close()

	// Парсим ответ
	var result struct {
		Routes []struct {
			Distance float64 `json:"distance"`
			Duration float64 `json:"duration"`
			Geometry struct {
				Coordinates [][]float64 `json:"coordinates"`
			} `json:"geometry"`
		} `json:"routes"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to parse route response: %v", err)
	}

	if len(result.Routes) == 0 {
		return nil, fmt.Errorf("no routes found")
	}

	route := result.Routes[0]
	
	// Конвертируем координаты в формат приложения
	var routeCoordinates []models.Coordinate
	for _, coord := range route.Geometry.Coordinates {
		routeCoordinates = append(routeCoordinates, models.Coordinate{
			Latitude:  coord[1],
			Longitude: coord[0],
		})
	}

	return &models.RouteInfo{
		Distance:    route.Distance,
		Duration:    route.Duration,
		Coordinates: routeCoordinates,
	}, nil
}

// GetElevation получает информацию о высоте для координат
func (s *MapService) GetElevation(coordinates []models.Coordinate) ([]float64, error) {
	if len(coordinates) == 0 {
		return nil, fmt.Errorf("coordinates are required")
	}

	// Формируем URL для запроса к API Яндекс Высот
	url := fmt.Sprintf("https://api-maps.yandex.ru/2.1/?apikey=%s&load=package.full", s.apiKey)
	
	// Формируем точки для запроса высот
	var points string
	for i, coord := range coordinates {
		if i > 0 {
			points += "~"
		}
		points += fmt.Sprintf("%f,%f", coord.Latitude, coord.Longitude)
	}
	url += "&points=" + points

	// Выполняем запрос
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get elevation data: %v", err)
	}
	defer resp.Body.Close()

	// Парсим ответ
	var result struct {
		Elevations []float64 `json:"elevations"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to parse elevation response: %v", err)
	}

	return result.Elevations, nil
}

// GetWeather получает информацию о погоде для координат
func (s *MapService) GetWeather(coordinate models.Coordinate) (*models.Weather, error) {
	// Формируем URL для запроса к API Яндекс Погоды
	url := fmt.Sprintf("https://api.weather.yandex.ru/v2/forecast?lat=%f&lon=%f", 
		coordinate.Latitude, coordinate.Longitude)
	
	// Выполняем запрос
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create weather request: %v", err)
	}
	
	req.Header.Set("X-Yandex-API-Key", s.apiKey)
	
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get weather data: %v", err)
	}
	defer resp.Body.Close()

	// Парсим ответ
	var result struct {
		Fact struct {
			Temp      float64 `json:"temp"`
			FeelsLike float64 `json:"feels_like"`
			Condition string  `json:"condition"`
			WindSpeed float64 `json:"wind_speed"`
			Humidity  float64 `json:"humidity"`
		} `json:"fact"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to parse weather response: %v", err)
	}

	return &models.Weather{
		Temperature: result.Fact.Temp,
		FeelsLike:   result.Fact.FeelsLike,
		Condition:   result.Fact.Condition,
		WindSpeed:   result.Fact.WindSpeed,
		Humidity:    result.Fact.Humidity,
	}, nil
} 