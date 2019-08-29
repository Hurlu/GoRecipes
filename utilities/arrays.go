package utilities

// Return -1 if not in the array: else, the position in the array.
func Contains(haystack []string, needle string) int {
	for pos, hay := range haystack {
		if hay == needle {
			return pos
		}
	}
	return -1
}

func Remove(haystack []string, needle string) []string {
	for pos, hay := range haystack {
		if hay == needle {
			return append(haystack[:pos], haystack[pos+1:]...)
		}
	}
	return haystack
}