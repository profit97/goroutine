package main

import (
	"testing"
)

// Пример тестовой функции
func TestGenerateRandomElements(t *testing.T) {
	// Тест для пустого слайса
	result := generateRandomElements(0)
	if len(result) != 0 {
		t.Errorf("Ожидиается пустой слайс, но имеет %d элементов", len(result))
	}

	// Тест для слайса с одним элементом
	result = generateRandomElements(1)
	if len(result) != 1 {
		t.Errorf("Ожидается слайс с 1 элементом, но имеет %d", len(result))
	}

	// Тест для слайса с несколькими элементами
	result = generateRandomElements(5)
	if len(result) != 5 {
		t.Errorf("ожидается слайс с 5 элементами, но имеет %d", len(result))
	}
}
func TestMaximum(t *testing.T) {
	// Тест для пустого слайса
	if maximum([]int{}) != 0 {
		t.Error("Ожидалось panic или ошибка для пустого слайса")
	}

	// Тест для слайса с одним элементом
	if result := maximum([]int{42}); result != 42 {
		t.Errorf("Ожидалось 42, получено %d", result)
	}

	// Тест для слайса с несколькими элементами
	if result := maximum([]int{1, 2, 3, 4, 5}); result != 5 {
		t.Errorf("Ожидалось 5, получено %d", result)
	}

	// Тест для слайса, где максимальное число в начале
	if result := maximum([]int{5, 1, 2}); result != 5 {
		t.Errorf("Ожидалось 5, получено %d", result)
	}
}
