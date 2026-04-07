package main

import (
	"testing"
)

// Пример тестовой функции
func TestGenerateRandomElements(t *testing.T) {
	testCases := []struct {
		size int
		want int
	}{
		{0, 0},
		{1, 1},
		{5, 5},
	}

	for _, tc := range testCases {
		result := generateRandomElements(tc.size)
		if len(result) != tc.want {
			t.Errorf("При size=%d, ожидается слайс с %d элементами, но имеет %d", tc.size, tc.want, len(result))
		}
	}
}
func TestMaximum(t *testing.T) {
	testCases := []struct {
		data   []int
		expect int
	}{
		{[]int{}, 0}, // Обратите внимание, что здесь мы ожидаем панику или ошибку
		{[]int{42}, 42},
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{5, 1, 2}, 5},
	}

	for _, tc := range testCases {
		defer func() {
			if r := recover(); r != nil && tc.data == nil {
				// Ожидаем панику для пустого слайса
				return
			}
			if tc.expect != maximum(tc.data) {
				t.Errorf("Для данных %v, ожидалось %d, получено %d", tc.data, tc.expect, maximum(tc.data))
			}
		}()
	}
}
