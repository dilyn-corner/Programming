package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	switch len(l) {
	case 0:
		return FreqMap{}
	case 1:
		return Frequency(l[0])
	}

	conChannel := make(chan FreqMap)
	conFunc := func(l []string) {
		conChannel <- ConcurrentFrequency(l)
	}

	halfLen := len(l) / 2
	go conFunc(l[:halfLen])
	go conFunc(l[halfLen:])
	m := <-conChannel
	for r, n := range <-conChannel {
		m[r] += n
	}
	return m
}
