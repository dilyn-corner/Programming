/*
Package letter counts the frequency of letters in a given text.
It utilizes concurrency to speed this process up substantially.
*/
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
	// In case l is just too dang small
	switch len(l) {
	case 0:
		return FreqMap{}
	case 1:
		return Frequency(l[0])
	}

	// Create a channel with the type FreqMap ^
	conChannel := make(chan FreqMap)

	// Create a function that submits the argument l to our channel
	conFunc := func(l []string) {
		conChannel <- ConcurrentFrequency(l)
	}

	// Split the string in half so we can act on each half concurrently
	halfLen := len(l) / 2

	// Call our concurrency function in a goroutine, one for each half
	go conFunc(l[:halfLen])
	go conFunc(l[halfLen:])

	// Create a variable which holds whatever is inside our channel
	// In this case, it's going to hold the frequency map FreqMap
	msg := <-conChannel

	// Range over the contents of our channel, incrementing appropriately
	for r, n := range <-conChannel {
		msg[r] += n
	}
	// Ultimately, this function is merely a clone of Frequency(), just
	// implemented using concurrency instead of just lollygagging.
	return msg
}
