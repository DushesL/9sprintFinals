package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
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
		return []int{}
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
	sort.Ints(data)
	return data[len(data)-1]
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

	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)

		partStart := len(data) / CHUNKS * i
		partEnd := partStart + len(data)/CHUNKS

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
	elapsed := time.Now().Sub(timeStart).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	// ваш код здесь
	timeStart = time.Now()
	max = maxChunks(rndElem)
	elapsed = time.Now().Sub(timeStart).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
