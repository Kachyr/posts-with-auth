package helpers

import "math"

func CalculateTotalPages(totalElements, size int) int {
	if size <= 0 {
		// Handle invalid page size
		return 0
	}

	return int(math.Ceil(float64(totalElements) / float64(size)))
}
