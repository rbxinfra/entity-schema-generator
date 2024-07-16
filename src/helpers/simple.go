package helpers

// Loop is a helper function that returns a channel that will send integers from the 'from' value to the 'to' value.
func Loop(from, to int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := from; i <= to; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

// GetFirstElement returns the first element of a map as a slice.
func GetFirstElement(m map[string]string) []string {
	var keys []string
	for k, v := range m {
		keys = append(keys, k, v)
		break
	}

	return keys
}
