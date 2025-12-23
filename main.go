package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{} // Возвращаем пустой слайс, если size <= 0
	}

	elements := make([]int, size)
	for i := range elements {
		elements[i] = rand.Intn(100000001) // Генерируем случайные числа в диапазоне от 0 до 100_000_000
	}
	return elements
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		// Обрабатываем случай пустого слайса
		panic("Слайс пуст") // или можно вернуть специальное значение, например, ошибку
	}

	max := data[0]
	for _, value := range data {
		if value > max {
			max = value
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	chunkSize := len(data) / CHUNKS
	maxValues := make([]int, CHUNKS)

	var wg sync.WaitGroup
	wg.Add(CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize

		if i == CHUNKS-1 {
			end = len(data) // Последний срез должен включать все оставшиеся элементы
		}

		go func(start, end int) {
			defer wg.Done()
			maxValues[i] = maximum(data[start:end])
		}(start, end)
	}

	// Дождитесь завершения всех горутин
	wg.Wait()

	// Найдите максимальное значение среди максимумов
	max := maxValues[0]
	for _, value := range maxValues {
		if value > max {
			max = value
		}
	}

	return max
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	// ваш код здесь
	data := generateRandomElements(SIZE)
	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
