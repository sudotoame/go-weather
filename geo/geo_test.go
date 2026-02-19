package geo_test

import (
	"testing"

	"demo/weather/geo"
)

func TestGetMyLocation(t *testing.T) {
	// Arange - Подготовка, expected результат, данные для функции
	city := "London"
	expected := geo.GeoData{
		City: "London",
	}
	// Act - выполнения функции
	data, err := geo.GetMyLocation(city)
	// Assert - Проверка результата с expected
	if err != nil {
		t.Error("Ошибка получении города")
	}
	if data.City != expected.City {
		t.Errorf("Ожидалось: %v, Получено: %v", expected, data.City)
	}
}
