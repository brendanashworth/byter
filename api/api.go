// Package byter is an API for counting occurrences of same-bytes, and stripping stray
// occurrences until a threshold.
package byter

// Counts byte occurrences within a given array of bytes.
// Returns an int representing the total bytes in the array and a map of
// each byte to the amount of occurrences within that array.
func CountOccurrences(data []byte) (int, map[byte]int) {
	// get total amount of bytes
	totalBytes := len(data)

	// data is an array of bytes
	// create map [byte]int
	percentages := make(map[byte]int)

	// iterate over all the bytes
	for _, piece := range data {
		// if it does not have the byte already
		if _, ok := percentages[piece]; !ok {
			percentages[piece] = 1
			continue
		}

		percentages[piece]++
	}

	return totalBytes, percentages
}

// Iterates over given map of bytes (with size in float64), deleting
// each key and value pair within the map when its amount of occurrences
// does not cut the threshold. The threshold is between 0 and 1 (a percentage).
func StripOccurrences(data map[byte]int, size float64, threshold float64) map[byte]int {
	// iterate over bytes
	for key, occurrences := range data {
		if (float64(occurrences) / size) < threshold {
			delete(data, key)
		}
	}

	return data
}

// Retrieves the top byte occurrences within a dataset. If 2+ bytes
// occur in the dataset with the same number of occurrences, they will
// all be within the array.
func GetTopBytes(data map[byte]int) (result []byte) {
	// iterate over bytes
	top := 0

	for key, occurrences := range data {
		if occurrences > top {
			// clear array
			result = result[:0]

			// add new byte
			result[0] = key
		} else if occurrences == top {
			// append byte
			result[len(result)] = key
		}
	}
	
	return
}
