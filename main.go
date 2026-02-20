package main

import (
	"errors"
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
	// ваш код здесь
	if size <= 0 {
		fmt.Print(errors.New("incorrect size"))
		return nil
	}
	v := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		v[i] = rand.Int()
	}
	return v
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		fmt.Print(errors.New("empty slice"))
		return 0
	}
	max := data[0]
	for _, num := range data[1:] {
		if num > max {
			max = num
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	if data == nil {
		fmt.Print(errors.New("empty slice"))
		return 0
	}

	var wg sync.WaitGroup
	maxElem := make([]int, CHUNKS)

	lenDivChun := len(data) / CHUNKS

	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)

		partStart := lenDivChun * i
		partEnd := partStart + lenDivChun

		if i == CHUNKS-1 {
			partEnd = len(data)
		}

		go func(i int, part []int) {
			defer wg.Done()
			maxElem[i] = maximum(part)
		}(i, data[partStart:partEnd])

	}
	wg.Wait()
	return maximum(maxElem)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	// ваш код здесь
	rndElem := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	timeStart := time.Now()
	max := maximum(rndElem)
	elapsed := time.Since(timeStart).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	// ваш код здесь
	timeStart = time.Now()
	max = maxChunks(rndElem)
	elapsed = time.Since(timeStart).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
