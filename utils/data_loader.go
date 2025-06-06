package utils

import (
	"fmt"
	"time"
)

// LoadDataSync simulates loading heavy data synchronously.
func LoadDataSync() []string {
	start := time.Now()
	data := make([]string, 0)

	// Simulate heavy data loading
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second) // Simulate delay
		data = append(data, fmt.Sprintf("Data %d", i))
	}

	fmt.Printf("Synchronous loading took %v\n", time.Since(start))
	return data
}

// LoadDataAsync simulates loading heavy data asynchronously.
func LoadDataAsync(ch chan<- []string) {
	start := time.Now()
	data := make([]string, 0)

	// Simulate heavy data loading
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second) // Simulate delay
		data = append(data, fmt.Sprintf("Data %d", i))
	}

	fmt.Printf("Asynchronous loading took %v\n", time.Since(start))
	ch <- data
}